package com.example.demo.sys.dao.mapper;

import com.example.demo.sys.entity.CourseSelect;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.demo.sys.entity.CustomCourseExamInfo;
import com.example.demo.sys.entity.CustomCourseSelect;
import com.example.demo.sys.entity.CustomStudentTimetable;
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
public interface CourseSelectMapper extends BaseMapper<CourseSelect> {
    @Select("SELECT course_select.course_select_id AS studentCourseId, course.course_name AS courseName, teacher.teacher_name AS teacherName, course.credit AS credit, course_score AS score " +
            "FROM course_select " +
//            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE course_select.student_id = #{studentId}")
    List<CustomCourseSelect> getSelectedCoursesByStudentId(@Param("studentId") Integer studentId);

    @Insert("INSERT INTO course_select (student_id, course_id) VALUES (#{studentId}, #{courseId})")
    void insert(@Param("studentId") Integer studentId, @Param("courseId") Integer courseId);

    @Select("SELECT course_select.course_select_id AS studentCourseId, course.course_name AS courseName, teacher.teacher_name AS teacherName, course.exam_date AS examDate, course.exam_loc AS examLocation " +
            "FROM course_select " +
//            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE course_select.student_id = #{studentId}")
    List<CustomCourseExamInfo> getExamList(Integer studentId);

    @Select("SELECT course.course_name AS courseName, teacher.teacher_name AS teacherName, location , course_time AS courseTime " +
            "FROM course_select " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "INNER JOIN teacher ON course.teacher_id = teacher.teacher_id " +
            "WHERE course_select.student_id = #{studentId}")
    List<CustomStudentTimetable> getStudentTimetable(Integer studentId);

    @Select("SELECT COUNT(*) " +
            "FROM course_select " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "WHERE course.teacher_id = #{teacherId}")
    int getCourseCount(Integer teacherId);

    @Select("SELECT COUNT(*) " +
            "FROM course_select " +
            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "WHERE course.teacher_id = #{teacherId} AND student.student_name LIKE CONCAT('%', #{studentName}, '%')")
    int getCourseCountByStudent(String studentName, int teacherId);

    @Select("SELECT COUNT(*) " +
            "FROM course_select " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "WHERE course.teacher_id = #{teacherId} AND course.course_name LIKE CONCAT('%', #{courseName}, '%')")
    int getCourseCountByCourse(String courseName, int teacherId);

    @Select("SELECT COUNT(*) " +
            "FROM course_select " +
            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "WHERE course.teacher_id = #{teacherId} AND course.course_name LIKE CONCAT('%', #{courseName}, '%') AND student.student_name LIKE CONCAT('%', #{studentName}, '%')")
    int getCourseCountByBoth(String studentName, String courseName, int teacherId);

    @Select("SELECT course_select_id AS studentCourseId, course_name AS courseName, student_name AS studentName, course_score AS score " +
            "FROM course_select " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "WHERE course.teacher_id = #{teacherId} " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getCoursesByPage(int offset, int limit, int teacherId);

    @Select("SELECT course_select_id AS studentCourseId, course_name AS courseName, student_name AS studentName, course_score AS score " +
            "FROM course_select " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "WHERE teacher_id = #{teacherId} AND student.student_name LIKE CONCAT('%', #{studentName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getCoursesByPageByStudent(String studentName, int offset, int limit, int teacherId);

    @Select("SELECT course_select_id AS studentCourseId, course_name AS courseName, student_name AS studentName, course_score AS score " +
            "FROM course_select " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "WHERE course.teacher_id = #{teacherId} AND course.course_name LIKE CONCAT('%', #{courseName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getCoursesByPageByCourse(String courseName, int offset, int limit, int teacherId);

    @Select("SELECT course_select_id AS studentCourseId, course_name AS courseName, student_name AS studentName, course_score AS score " +
            "FROM course_select " +
            "INNER JOIN student ON course_select.student_id = student.student_id " +
            "INNER JOIN course ON course_select.course_id = course.course_id " +
            "WHERE course.teacher_id = #{teacherId} AND course.course_name LIKE CONCAT('%', #{courseName}, '%') AND student.student_name LIKE CONCAT('%', #{studentName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getCoursesByPageByBoth(String courseName, String studentName, int offset, int limit, int teacherId);

    @Select("SELECT course_select_id AS studentCourseId, course_score AS score FROM course_select WHERE course_select_id = #{courseSelectId}")
    Map<String, Object> getGrade(int courseSelectId);

    @Update("UPDATE course_select SET course_score = #{score} WHERE course_select_id = #{studentCourseId}")
    void updateGrade(int studentCourseId, int score);
}
