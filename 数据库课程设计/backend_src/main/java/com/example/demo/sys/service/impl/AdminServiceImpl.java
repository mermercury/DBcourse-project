package com.example.demo.sys.service.impl;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.dao.mapper.DepartmentMapper;
import com.example.demo.sys.entity.Admin;
import com.example.demo.sys.dao.mapper.AdminMapper;
import com.example.demo.sys.service.IAdminService;
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
public class AdminServiceImpl extends ServiceImpl<AdminMapper, Admin> implements IAdminService {

}
