package com.example.demo.sys.service.impl;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.dao.mapper.*;
import com.example.demo.sys.entity.CourseSelect;
import com.example.demo.sys.entity.CustomStudentCourse;
import com.example.demo.sys.entity.CustomStudentTimetable;
import com.example.demo.sys.entity.Teacher;
import com.example.demo.sys.service.ITeacherService;
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
public class TeacherServiceImpl extends ServiceImpl<TeacherMapper, Teacher> implements ITeacherService {
    @Autowired
    private TeacherMapper teacherMapper;
    @Autowired
    private CourseMapper courseMapper;
    @Autowired
    private CourseSelectMapper selectMapper;

    @Override
    public Result getTeacherCourseList(String teacherName) {
        List<Map<String, Object>> data = courseMapper.getTeacherCourseList(teacherName);
        for (Map<String, Object> map : data) {
            String courseTime = (String) map.get("courseTime");
            String[] parts = courseTime.split("-");
            int weekDay = Integer.parseInt(parts[0]);
            int startPeriod = Integer.parseInt(parts[1]);
            int sections = Integer.parseInt(parts[2]);
            // 定义一个数组用于存储星期几的名称
            String[] weekDays = {"null", "周一", "周二", "周三", "周四", "周五", "周六", "周日"};
            // 将数字转换为对应的星期几
            String weekDayName = weekDays[weekDay];
            // 构建要返回的格式字符串
            String result = weekDayName + "第" + startPeriod + "~" + (startPeriod + sections) + "节";
            map.put("courseTime", result);
        }
        return Result.success(data);
    }

    @Override
    public Teacher getTeacherByName(String username) {
        return teacherMapper.selectByName(username);
    }

    @Override
    public Result getTeacherTimetable(Integer teacherId) {
        List<Map<String, Object>> list = courseMapper.getTeacherTimetable(teacherId);
        // 初始化二维列表

        Object[][] timetable = new Object[7][10];

        // 遍历list中的每个CustomStudentTimetable对象
        for (Map<String, Object> course : list) {
            // 解析courseTime字段，格式为"星期几-第几节开始-上几节课"
            String[] timeParts = ((String) (course.get("courseTime"))).split("-");
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

    @Override
    public Result getCoursePageCount(String courseName, String studentName, int pageSize, int teacherId) {
        int count = 0;
        if ((courseName == null || courseName.equals("")) && (studentName == null || studentName.equals(""))) {
            // 查询这位老师教的所有课程
            count = selectMapper.getCourseCount(teacherId);
        } else if (courseName == null || courseName.equals("")) {
            count = selectMapper.getCourseCountByStudent(studentName, teacherId);
        } else if ((studentName == null || studentName.equals(""))) {
            count = selectMapper.getCourseCountByCourse(courseName, teacherId);
        } else {
            count = selectMapper.getCourseCountByBoth(studentName, courseName, teacherId);
        }
        Map<String, Object> data = new HashMap<>();
        count  = (int) Math.ceil(count * 1.0 / pageSize);
        data.put("count", count);
        return Result.success(data);
    }

    @Override
    public Result getCoursesByPage(String courseName, String studentName, int offset, int limit, int teacherId) {
        List<Map<String, Object>> data = null;
        if ((courseName == null || courseName.equals("")) && (studentName == null || studentName.equals(""))) {
            // 查询所有课程
            data = selectMapper.getCoursesByPage(offset, limit, teacherId);
        } else if (courseName == null || courseName.equals("")) {
            data = selectMapper.getCoursesByPageByStudent(studentName, offset, limit, teacherId);
        } else if ((studentName == null || studentName.equals(""))) {
            data = selectMapper.getCoursesByPageByCourse(courseName, offset, limit, teacherId);
        } else {
            data = selectMapper.getCoursesByPageByBoth(courseName, studentName, offset, limit, teacherId);
        }
        return Result.success(data);
    }

    @Override
    public Map<String, Object> getGrade(int courseSelectId) {
        return selectMapper.getGrade(courseSelectId);
    }

    @Override
    public Result updateGrade(int studentCourseId, int score) {
        if (score == -1) {
            return Result.fail(-1, "当前未登成绩!");
        } else if (score < -1 || score > 100) {
            return Result.fail(-2, "当前成绩不合法！");
        } else {
            selectMapper.updateGrade(studentCourseId, score);
            return Result.success("更新成绩成功！");
        }
    }
}
