package com.example.demo.sys.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.core.conditions.update.UpdateWrapper;
import com.example.demo.common.vo.Result;
import com.example.demo.sys.dao.dao.StudentDAO;
import com.example.demo.sys.dao.mapper.*;
import com.example.demo.sys.entity.*;
import com.example.demo.sys.service.IStudentService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@Service
public class StudentServiceImpl extends ServiceImpl<StudentMapper, Student> implements IStudentService {
    @Autowired
    private StudentMapper studentMapper;
    @Autowired
    private DepartmentMapper departmentMapper;
    @Autowired
    private ClassMapper classMapper;
    @Autowired
    private MajorMapper majorMapper;
    @Autowired
    private CourseSelectMapper selectMapper;
    @Autowired
    private CourseMapper courseMapper;

    public Student getStudentByName(String name) {
        Student s = studentMapper.selectByName(name);
        return s;
    }

    @Override
    public List<CustomCourseSelect> getCourseSelect(Integer studentId) {
        return selectMapper.getSelectedCoursesByStudentId(studentId);
    }

    @Override
    public Result getCoursePageCount(String courseName, String teacherName, int pageSize) {
        int count = 0;
        if ((courseName == null || courseName.equals("")) && (teacherName == null || teacherName.equals(""))) {
            // 查询所有课程
            count = courseMapper.getCourseCount();
        } else if (courseName == null || courseName.equals("")) {
            count = courseMapper.getCourseCountByTeacher(teacherName);
        } else if ((teacherName == null || teacherName.equals(""))) {
            count = courseMapper.getCourseCountByCourse(courseName);
        } else {
            count = courseMapper.getCourseCountByBoth(teacherName, courseName);
        }
        Map<String, Object> data = new HashMap<>();
        count  = (int) Math.ceil(count * 1.0 / pageSize);
        data.put("count", count);
        return Result.success(data);
    }

    @Override
    public Result<List<CustomStudentCourse>> getCoursesByPage(String courseName, String teacherName, int offset, int limit) {
        List<CustomStudentCourse> data = null;
        if ((courseName == null || courseName.equals("")) && (teacherName == null || teacherName.equals(""))) {
            // 查询所有课程
            data = courseMapper.getCoursesByPage(offset, limit);
        } else if (courseName == null || courseName.equals("")) {
            data = courseMapper.getCoursesByPageByTeacher(teacherName, offset, limit);
        } else if ((teacherName == null || teacherName.equals(""))) {
            data = courseMapper.getCoursesByPageByCourse(courseName, offset, limit);
        } else {
            data = courseMapper.getCoursesByPageByBoth(courseName, teacherName, offset, limit);
        }
        return Result.success(data);
    }

    // 学生试图选这门课
    @Override
    public Result selectCourse(Integer studentId, Integer courseId) {
        LambdaQueryWrapper<CourseSelect> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(CourseSelect::getStudentId, studentId)
                .eq(CourseSelect::getCourseId, courseId);
        Object obj = selectMapper.selectOne(wrapper);
        if (obj != null) {
            return Result.fail(-1, "本课程已被选修过！");
        } else if (checkConflict(studentId, courseId)) {
            return Result.fail(-2, "本课程与已选课程冲突！");
        } else {
            selectMapper.insert(studentId, courseId);
            courseMapper.incrementSelected(courseId);
            return Result.success("选课成功");
        }
    }

    @Override
    public Result deleteCourse(Integer stuId, Integer studentCourseId) {
        // 删除一条选课记录
        // TODO: HOW TO PROMISE THIS IS A TRANSACTION!!!
        LambdaQueryWrapper<CourseSelect> queryWrapper = new LambdaQueryWrapper<>();
        queryWrapper.eq(CourseSelect::getCourseSelectId, studentCourseId);
        CourseSelect courseSelect = selectMapper.selectOne(queryWrapper);

        int courseId = -1;
        if (courseSelect != null) {
            courseId = courseSelect.getCourseId();
        }
        queryWrapper.eq(CourseSelect::getCourseId, courseId)
                .eq(CourseSelect::getStudentId, stuId);
        selectMapper.delete(queryWrapper);
        // 把课表中selected字段+1
        courseMapper.minusSelected(courseId);
        return Result.success();
    }

    @Override
    public Result<List<CustomCourseExamInfo>> getExamList(Integer studentId) {
        List<CustomCourseExamInfo> list = selectMapper.getExamList(studentId);
        return Result.success(list);
    }

    @Override
    public Result updateStudentInfo(Student student) {
        UpdateWrapper<Student> updateWrapper = new UpdateWrapper<>();
        // 设置更新条件，这里假设你根据学号id来更新数据
        updateWrapper.eq("student_id", student.getStudentId());
        System.out.println(student);

        // 执行更新操作
        studentMapper.update(student, updateWrapper);
        return null;
    }

    @Override
    public Result getStudentTimetable(Integer studentId) {
        List<CustomStudentTimetable> list = selectMapper.getStudentTimetable(studentId);
        // 初始化二维列表
        CustomStudentTimetable[][] timetable = new CustomStudentTimetable[7][10];

        // 遍历list中的每个CustomStudentTimetable对象
        for (CustomStudentTimetable course : list) {
            // 解析courseTime字段，格式为"星期几-第几节开始-上几节课"
            String[] timeParts = course.getCourseTime().split("-");
            int dayOfWeek = Integer.parseInt(timeParts[0]); // 星期几
            int startSection = Integer.parseInt(timeParts[1]); // 第几节开始
            int numSections = Integer.parseInt(timeParts[2]); // 上几节课

            // 将课程信息填充到二维列表中
            for (int i = startSection - 1; i < startSection - 1 + numSections; i++) {
                timetable[dayOfWeek - 1][i] = course;
            }
        }
        return Result.success(timetable);
    }

    public Result<Map<String, Object>> getStudentInfo (String username) {
        Student s = studentMapper.selectByName(username);
        Map<String, Object> data = new HashMap<>();
        data.put("studentId", s.getStudentId());
        data.put("studentName", s.getStudentName());
        data.put("departmentName", departmentMapper.getDepartmentNameById(s.getDepartmentId()));
        data.put("majorName", majorMapper.getMajorNameById(s.getMajorId()));
        data.put("className", classMapper.getClassNameById(s.getClassId()));
        data.put("email", s.getEmail());
        data.put("birthday", s.getBirthday());
        data.put("sex", s.getSex());
        data.put("password", s.getPassword());
        return Result.success(data);
    }

    private boolean checkConflict(Integer studentId, Integer courseId) {
        List<CustomStudentTimetable> list = selectMapper.getStudentTimetable(studentId);
        String time = courseMapper.selectById(courseId).getCourseTime();
        String[] times = time.split("-");
        int day = Integer.parseInt(times[0]); // 星期几
        int start = Integer.parseInt(times[1]); // 第几节开始
        int nums = Integer.parseInt(times[2]); // 上几节课

        // 遍历list中的每个CustomStudentTimetable对象
        for (CustomStudentTimetable course : list) {
            // 解析courseTime字段，格式为"星期几-第几节开始-上几节课"
            String[] timeParts = course.getCourseTime().split("-");
            int dayOfWeek = Integer.parseInt(timeParts[0]); // 星期几
            int startSection = Integer.parseInt(timeParts[1]); // 第几节开始
            int numSections = Integer.parseInt(timeParts[2]); // 上几节课

            if (dayOfWeek == day && Math.max(start, startSection) < Math.min(start + nums, startSection + numSections)) {
                return true; // 有冲突
            }
        }

        return false;
    }
}
