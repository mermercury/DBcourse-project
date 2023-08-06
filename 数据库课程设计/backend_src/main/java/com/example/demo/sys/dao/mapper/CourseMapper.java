package com.example.demo.sys.dao.mapper;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Course;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.demo.sys.entity.CustomStudentCourse;
import org.apache.ibatis.annotations.*;

import java.util.List;
import java.util.Map;

/**
 * <p>
 *  Mapper 接口
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@Mapper
public interface CourseMapper extends BaseMapper<Course> {
    @Select("SELECT COUNT(*) FROM course")
    int getCourseCount();

    @Select("SELECT COUNT(*) " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE teacher.teacher_name LIKE CONCAT('%', #{teacherName}, '%')")
    int getCourseCountByTeacher(@Param("teacherName") String teacherName);

    @Select("SELECT COUNT(*) " +
            "FROM course " +
            "WHERE course_name LIKE CONCAT('%', #{courseName}, '%')")
    int getCourseCountByCourse(@Param("courseName") String courseName);

    @Select("SELECT COUNT(*) " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE teacher.teacher_name LIKE CONCAT('%', #{teacherName}, '%') AND course_name = #{courseName}")
    int getCourseCountByBoth(@Param("teacherName") String teacherName, @Param("courseName") String courseName);

    @Select("SELECT course_id AS courseId, course_name AS courseName, teacher.teacher_name AS teacherName, course.credit AS credit, course_time AS time, selected AS selectedCount, size AS maxSize " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE course_name LIKE CONCAT('%', #{courseName}, '%') " +
            "AND teacher_name LIKE CONCAT('%', #{teacherName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<CustomStudentCourse> getCoursesByPageByBoth(
            @Param("courseName") String courseName,
            @Param("teacherName") String teacherName,
            @Param("offset") int offset,
            @Param("limit") int limit
    );

    @Select("SELECT course_id AS courseId, course_name AS courseName, teacher.teacher_name AS teacherName, course.credit AS credit, course_time AS time, selected AS selectedCount, size AS maxSize " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "LIMIT #{offset}, #{limit}")
    List<CustomStudentCourse> getCoursesByPage(
            @Param("offset") int offset,
            @Param("limit") int limit
    );

    @Select("SELECT course_id AS courseId, course_name AS courseName, teacher.teacher_name AS teacherName, course.credit AS credit, course_time AS time, selected AS selectedCount, size AS maxSize " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE teacher_name LIKE CONCAT('%', #{teacherName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<CustomStudentCourse> getCoursesByPageByTeacher(
            @Param("teacherName") String teacherName,
            @Param("offset") int offset,
            @Param("limit") int limit
    );

    @Select("SELECT course_id AS courseId, course_name AS courseName, teacher.teacher_name AS teacherName, course.credit AS credit, course_time AS time, selected AS selectedCount, size AS maxSize " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE course_name LIKE CONCAT('%', #{courseName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<CustomStudentCourse> getCoursesByPageByCourse(
            @Param("courseName") String courseName,
            @Param("offset") int offset,
            @Param("limit") int limit
    );

    @Update("UPDATE course SET selected = selected + 1 WHERE course_id = #{courseId}")
    int incrementSelected(@Param("courseId") Integer courseId);

    @Update("UPDATE course SET selected = selected - 1 WHERE course_id = #{courseId}")
    int minusSelected(@Param("courseId") Integer courseId);

    @Select("SELECT course_id AS id, course_name AS courseName, grade, course.credit AS credit, course_time AS courseTime, location, selected, exam_date AS examDate " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE teacher_name = #{teacherName}")
    List<Map<String, Object>> getTeacherCourseList(String teacherName);


    @Select("SELECT course_name AS courseName, teacher.teacher_name AS teacherName, location, course_time AS courseTime " +
            "FROM course " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE course.teacher_id = #{teacherId}")
    List<Map<String, Object>> getTeacherTimetable(Integer teacherId);
}
