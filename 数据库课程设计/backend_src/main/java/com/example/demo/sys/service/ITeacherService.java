package com.example.demo.sys.service;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.CourseSelect;
import com.example.demo.sys.entity.CustomStudentCourse;
import com.example.demo.sys.entity.Student;
import com.example.demo.sys.entity.Teacher;
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
public interface ITeacherService extends IService<Teacher> {

    Result getTeacherCourseList(String teacherName);

    Teacher getTeacherByName(String username);

    Result getTeacherTimetable(Integer teacherId);

    Result getCoursePageCount(String courseName, String studentName, int pageSize, int teacherId);

    Result getCoursesByPage(String courseName, String studentName, int offset, int limit, int teacherId);

    Map<String, Object> getGrade(int courseSelectId);

    Result updateGrade(int studentCourseId, int score);
}
