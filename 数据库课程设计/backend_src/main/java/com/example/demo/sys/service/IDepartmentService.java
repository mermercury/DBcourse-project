package com.example.demo.sys.service;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Department;
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
public interface IDepartmentService extends IService<Department> {
    Result getDepartmentPageCount(String departmentName, int pageSize);

    Result getDepartmentByPage(String departmentName, int offset, int limit);

    Result getDepartmentById(int id);

    Result updateDepartment(int departmentId, String departmentName, int majorCount, int teacherCount);

    Result createDepartment(String departmentName, int majorCount, int teacherCount);

    Result getDepartmentNameList();
}
