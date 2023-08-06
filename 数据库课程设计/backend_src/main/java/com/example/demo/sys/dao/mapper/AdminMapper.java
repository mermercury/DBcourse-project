package com.example.demo.sys.dao.mapper;

import com.example.demo.sys.entity.Admin;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.demo.sys.entity.Teacher;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.springframework.stereotype.Component;

/**
 * <p>
 *  Mapper 接口
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@Mapper
public interface AdminMapper extends BaseMapper<Admin> {
    @Select("SELECT * FROM admin WHERE admin_name = #{name}")
    Admin selectByName(String name);
}
