package com.example.demo.sys.dao.mapper;

import com.example.demo.sys.entity.Class;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;

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
public interface ClassMapper extends BaseMapper<Class> {
    @Select("SELECT class_name FROM class WHERE class_id = #{classId}")
    String getClassNameById(Integer classId);

    @Delete("DELETE FROM class WHERE major_id = #{id}")
    void deleteClassesByMajorId(int id);

    @Select("SELECT class_id AS id, class_name AS name, grade, department_name AS departmentName, major_name AS majorName " +
            "FROM class " +
            "INNER JOIN department ON department.department_id = class.department_id " +
            "INNER JOIN major ON major.major_id = class.major_id " +
            "WHERE class_id = #{classId}")
    Map<String, Object> getAClassById(int classId);
}
