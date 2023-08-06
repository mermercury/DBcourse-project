package com.example.demo.sys.dao.mapper;

import com.example.demo.sys.entity.Major;
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
public interface MajorMapper extends BaseMapper<Major> {
    @Select("SELECT major_name FROM major WHERE major_id = #{majorId}")
    String getMajorNameById(Integer majorId);

    @Select("SELECT COUNT(*) FROM major")
    int getMajorCount();

    @Select("SELECT COUNT(*) " +
            "FROM major " +
            "WHERE major_name LIKE CONCAT('%', #{majorName}, '%')")
    int getMajorCountByMajorName(String majorName);

    @Select("SELECT COUNT(*) " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "WHERE department_name LIKE CONCAT('%', #{departmentName}, '%')")
    int getMajorCountByDepartmentName(String departmentName);

    @Select("SELECT COUNT(*) " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "WHERE department_name LIKE CONCAT('%', #{departmentName}, '%') AND major_name LIKE CONCAT('%', #{majorName}, '%')")
    int getMajorCountByBoth(String majorName, String departmentName);

    @Select("SELECT major_id AS id, major_name AS name, department_name AS departmentName " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "ORDER BY major.major_id ASC " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getMajorByPage(int offset, int limit);

    @Select("SELECT major_id AS id, major_name AS name, department_name AS departmentName " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "WHERE major_name LIKE CONCAT('%', #{majorName}, '%') " +
            "ORDER BY major.major_id ASC " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getMajorByPageByMajorName(String majorName, int offset, int limit);

    @Select("SELECT major_id AS id, major_name AS name, department_name AS departmentName " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "WHERE department_name LIKE CONCAT('%', #{departmentName}, '%') " +
            "ORDER BY major.major_id ASC " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getMajorByPageByDepartmentName(String departmentName, int offset, int limit);

    @Select("SELECT major_id AS id, major_name AS name, department_name AS departmentName " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "WHERE department_name LIKE CONCAT('%', #{departmentName}, '%') AND major_name LIKE CONCAT('%', #{majorName}, '%') " +
            "ORDER BY major.major_id ASC " +
            "LIMIT #{offset}, #{limit}")
    List<Map<String, Object>> getMajorByPageByBoth(String departmentName, String majorName, int offset, int limit);

    @Select("SELECT major_id AS id, major_name AS name, department_name AS departmentName " +
            "FROM major " +
            "INNER JOIN department ON department.department_id = major.department_id " +
            "WHERE major_id = #{id}")
    Map<String, Object> getMajorById(int id);

    @Update("UPDATE major " +
            "SET major_name = #{majorNewName} " +
            "WHERE major_id = #{majorId}")
    void updateMajorName(int majorId, String majorNewName);
}
