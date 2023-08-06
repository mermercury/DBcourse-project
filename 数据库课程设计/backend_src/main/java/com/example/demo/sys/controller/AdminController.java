package com.example.demo.sys.controller;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.CustomStudentCourse;
import com.example.demo.sys.service.IAdminService;
import com.example.demo.sys.service.IClassService;
import com.example.demo.sys.service.IDepartmentService;
import com.example.demo.sys.service.IMajorService;
import org.springframework.web.bind.annotation.*;
import org.springframework.stereotype.Controller;

import java.util.List;
import java.util.Map;

/**
 * <p>
 *  前端控制器
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
@RestController
@RequestMapping("/api/v1/admin")
public class AdminController {
    // TODO: 后续需要对token进行统一判定其为admin【swagger？】

    private IDepartmentService departmentService;
    private IMajorService majorService;
    private IClassService classService;

    public AdminController(IDepartmentService departmentService, IMajorService majorService, IClassService classService) {
        this.departmentService = departmentService;
        this.majorService = majorService;
        this.classService = classService;
    }

    // TODO: 系管理 DEPARTMENT
    @GetMapping("/department/page/count")
    public Result getDepartmentPageCount(
            @RequestParam("name") String departmentName,
            @RequestParam("pageSize") int pageSize
    ) {
        return departmentService.getDepartmentPageCount(departmentName, pageSize);
    }

    @GetMapping("/department/page/{index}")
    public Result getDepartmentByPage(
            @PathVariable int index,
            @RequestParam("name") String departmentName,
            @RequestParam("pageSize") int pageSize) {

        // 使用index参数进行分页计算
        int offset = (index - 1) * pageSize;
        int limit = pageSize;

        return departmentService.getDepartmentByPage(departmentName, offset, limit);
    }

    @GetMapping("/department/{id}")
    public Result getDepartmentById(@PathVariable int id) {
        return departmentService.getDepartmentById(id);
    }

    @PutMapping("/department")
    public Result updateDepartment(@RequestBody Map<String, Object> requestBody) {
        int departmentId = (int) requestBody.get("id");
        String departmentName = (String) requestBody.get("name");
        int majorCount = (int) requestBody.get("majorCount");
        int teacherCount = (int) requestBody.get("teacherCount");
        if (departmentId != -1) {
            return departmentService.updateDepartment(departmentId, departmentName, majorCount, teacherCount);
        }
        else {
            return departmentService.createDepartment(departmentName, majorCount, teacherCount);
        }
    }

    @GetMapping("/department/names")
    public Result getDepartmentNameList() {
        return departmentService.getDepartmentNameList();
    }


    // TODO: 专业管理 MAJOR
    @GetMapping("/major/page/count")
    public Result getMajorPageCount(
            @RequestParam("departmentName") String departmentName,
            @RequestParam("name") String majorName, // 专业名
            @RequestParam("pageSize") int pageSize
    ) {
        return majorService.getMajorPageCount(departmentName, majorName, pageSize);
    }

    @GetMapping("/major/page/{index}")
    public Result getMajorByPage(
            @PathVariable int index,
            @RequestParam("departmentName") String departmentName,
            @RequestParam("name") String majorName, // 专业名
            @RequestParam("pageSize") int pageSize) {

        // 使用index参数进行分页计算
        int offset = (index - 1) * pageSize;
        int limit = pageSize;

        return majorService.getMajorByPage(departmentName, majorName, offset, limit);
    }

    @GetMapping("/major/{id}")
    public Result getAMajor(@PathVariable int id) {
        return majorService.getAMajor(id);
    }

    @PutMapping("/major")
    public Result updateMajorName(@RequestBody Map<String, Object> requestBody) {
        int majorId = (int) requestBody.get("id");
        if (majorId != -1) {
            String majorNewName = (String) requestBody.get("name");
            return majorService.updateMajorName(majorId, majorNewName);
        } else {
            String majorName = (String) requestBody.get("name");
            String departmentName = (String) requestBody.get("departmentName");
            return majorService.createMajor(majorName, departmentName);
        }
    }

    // TODO: 删除一个专业【场景：一个专业这届学生毕业之后合并进别的专业了】 =>
    // 1.删除该专业所有学生 2.对应department的major_count减1 3.删除该专业所有班级
    @DeleteMapping("/major/{id}")
    public Result deleteAMajor(@PathVariable int id) {
        return majorService.deleteAMajor(id);
    }

    // TODO: 班级管理
    @GetMapping("/class/{id}")
    public Result getClassInfo(@PathVariable int id) {
        return classService.getClassInfo(id);
    }

    @GetMapping("/class/page/count")
    public Result getClassPageCount(
            @RequestParam("majorName") String majorName,
            @RequestParam("departmentName") String departmentName,
            @RequestParam("name") String studentName,
            @RequestParam("pageSize") int pageSize
    ) {
        return classService.getClassPageCount(majorName, departmentName, studentName, pageSize);
    }

    @GetMapping("/class/page/{index}")
    public Result getClassByPage(
            @PathVariable int index,
            @RequestParam("majorName") String majorName,
            @RequestParam("departmentName") String departmentName,
            @RequestParam("name") String studentName,
            @RequestParam("pageSize") int pageSize) {

        // 使用index参数进行分页计算
        int offset = (index - 1) * pageSize;
        int limit = pageSize;

        return classService.getClassByPage(majorName, departmentName, studentName, limit);
    }

}
