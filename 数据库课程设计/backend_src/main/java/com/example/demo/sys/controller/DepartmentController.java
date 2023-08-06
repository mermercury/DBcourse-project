package com.example.demo.sys.controller;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.Department;
import com.example.demo.sys.service.IDepartmentService;
import com.example.demo.sys.service.impl.DepartmentServiceImpl;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RestController;

import java.lang.reflect.Array;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/**
 * <p>
 *  前端控制器
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@RestController
@RequestMapping("/api/v1/department")
@RequiredArgsConstructor
public class DepartmentController {

    private final IDepartmentService departmentService;

    @GetMapping("/all")
    public Result<List<Department>> getAllDepartment() {
        List<Department> list = departmentService.list();
        return Result.success("查询所有系成功", list);
    }

}
