package com.example.demo.sys.dao.mapper;

import com.example.demo.sys.entity.Course;
import com.example.demo.sys.entity.Student;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * <p>
 *  Mapper 接口
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@Mapper
public interface StudentMapper extends BaseMapper<Student> {
    // 查询所选课程
    @Select("SELECT department_name FROM department WHERE department_id = #{departmentId}")
    List<Course> selectAllCourses(String stuId);

    @Select("SELECT * FROM student WHERE student_name = #{name}")
    Student selectByName(String name);

    @Delete("DELETE FROM student WHERE major_id = #{id}")
    void deleteStudentsByMajorId(int id);
}
