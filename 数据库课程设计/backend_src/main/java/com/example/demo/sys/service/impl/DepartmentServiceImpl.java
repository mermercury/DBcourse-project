package com.example.demo.sys.service.impl;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Department;
import com.example.demo.sys.dao.mapper.DepartmentMapper;
import com.example.demo.sys.service.IDepartmentService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
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
public class DepartmentServiceImpl extends ServiceImpl<DepartmentMapper, Department> implements IDepartmentService {
    @Autowired
    private DepartmentMapper departmentMapper;

    @Override
    public Result getDepartmentPageCount(String departmentName, int pageSize) {
        int count = 0;
        if (departmentName == null || departmentName.equals("")) {
            // 查询所有系
            count = departmentMapper.getCount();
        } else {
            count = departmentMapper.getCountByName(departmentName);
        }
        Map<String, Object> data = new HashMap<>();
        count  = (int) Math.ceil(count * 1.0 / pageSize);
        data.put("count", count);
        return Result.success(data);
    }

    @Override
    public Result getDepartmentByPage(String departmentName, int offset, int limit) {
        List<Map<String, Object>> data;
        if (departmentName == null || departmentName.equals("")) {
            // 查询所有系
            data = departmentMapper.getDepartmentByPage(offset, limit);
        } else {
            data = departmentMapper.getDepartmentByPageByName(departmentName, offset, limit);
        }
        return Result.success(data);
    }

    @Override
    public Result getDepartmentById(int id) {
        Map<String, Object> data = departmentMapper.getDepartmentById(id);
        return Result.success(data);
    }

    @Override
    public Result updateDepartment(int departmentId, String departmentName, int majorCount, int teacherCount) {
        departmentMapper.updateDepartment(departmentId, departmentName, majorCount, teacherCount);
        return Result.success();
    }

    @Override
    public Result createDepartment(String departmentName, int majorCount, int teacherCount) {
        Department department = new Department();
        department.setDepartmentName(departmentName);
        department.setMajorCount(majorCount);
        department.setTeacherCount(teacherCount);
        departmentMapper.insert(department);
        return Result.success();
    }

    @Override
    public Result getDepartmentNameList() {
        List<Map<String, Object>> data = departmentMapper.getDepartmentNameList();
        return Result.success(data);
    }


}
