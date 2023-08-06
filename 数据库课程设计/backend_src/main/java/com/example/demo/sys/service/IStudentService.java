package com.example.demo.sys.service;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.*;
import com.baomidou.mybatisplus.extension.service.IService;

import java.util.List;
import java.util.Map;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
public interface IStudentService extends IService<Student> {
    Result<Map<String, Object>> getStudentInfo (String username);

    Student getStudentByName(String name);

    List<CustomCourseSelect> getCourseSelect(Integer studentId);

    Result getCoursePageCount(String courseName, String teacherName, int pageSize);

    Result<List<CustomStudentCourse>> getCoursesByPage(String courseName, String teacherName, int offset, int limit);

    Result selectCourse(Integer stuId, Integer courseId);

    Result deleteCourse(Integer stuId, Integer studentCourseId);

    Result<List<CustomCourseExamInfo>> getExamList(Integer studentId);

    Result updateStudentInfo(Student student);

    Result getStudentTimetable(Integer studentId);

}
