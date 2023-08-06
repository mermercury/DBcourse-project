package com.example.demo.sys.dao.mapper;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Department;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

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
public interface DepartmentMapper extends BaseMapper<Department> {
    @Select("SELECT department_name FROM department WHERE department_id = #{departmentId}")
    String getDepartmentNameById(Integer departmentId);

    @Select("SELECT COUNT(*) FROM department")
    int getCount();

    @Select("SELECT COUNT(*) FROM department " +
            "WHERE department_name LIKE CONCAT('%', #{departmentName}, '%')")
    int getCountByName(String departmentName);

    @Select("SELECT department_id AS id, department_name AS name, major_count AS majorCount, teacher_count AS teacherCount " +
            "FROM department " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getDepartmentByPage(int offset, int limit);

    @Select("SELECT department_id AS id, department_name AS name, major_count AS majorCount, teacher_count AS teacherCount " +
            "FROM department " +
            "WHERE department_name LIKE CONCAT('%', #{departmentName}, '%') " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getDepartmentByPageByName(String departmentName, int offset, int limit);

    @Select("SELECT department_id AS id, department_name AS name, major_count AS majorCount, teacher_count AS teacherCount " +
            "FROM department " +
            "WHERE department_id = #{id}")
    Map<String, Object> getDepartmentById(int id);

    @Update("UPDATE department " +
            "SET department_name = #{departmentName}, " +
            "major_count = #{majorCount}, " +
            "teacher_count = #{teacherCount} " +
            "WHERE department_id = #{departmentId}")
    void updateDepartment(int departmentId, String departmentName, int majorCount, int teacherCount);

    @Select("SELECT department_id AS id, department_name AS name " +
            "FROM department")
    List<Map<String, Object>> getDepartmentNameList();

    @Select("SELECT department_id FROM department WHERE department_name = #{name}")
    int getDepartmentIdByName(String name);

    @Update("UPDATE department SET major_count = major_count - 1 WHERE department_id = #{departmentId}")
    void decreaseMajorCount(int departmentId);

    @Update("UPDATE department SET major_count = major_count + 1 WHERE department_id = #{departmentId}")
    void increaseMajorCount(int departmentId);
}
