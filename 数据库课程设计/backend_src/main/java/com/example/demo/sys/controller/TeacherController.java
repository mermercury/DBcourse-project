package com.example.demo.sys.controller;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.*;
import com.example.demo.sys.service.ITeacherService;
import com.example.demo.utils.JwtUtil;
import org.springframework.web.bind.annotation.*;
import org.springframework.stereotype.Controller;

import java.util.ArrayList;
import java.util.HashMap;
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
@RequestMapping("/api/v1/teacher")
public class TeacherController {
    private ITeacherService teacherService;

    public TeacherController(ITeacherService teacherService) {
        this.teacherService = teacherService;
    }

    // 教师授课查询
    @GetMapping("/course/list")
    public Result getTeacherCourseList(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        return teacherService.getTeacherCourseList(username);
    }

    @GetMapping("/timetable")
    public Result getTeacherTimetable(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Teacher teacher = teacherService.getTeacherByName(username);
        return teacherService.getTeacherTimetable(teacher.getTeacherId());
    }

    @GetMapping("/grade/page/count")
    public Result getCoursePageCount(
            @RequestParam("courseName") String courseName,
            @RequestParam("studentName") String studentName,
            @RequestParam("pageSize") int pageSize,
            @RequestHeader("Authorization") String authorizationHeader
    ) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Teacher teacher = teacherService.getTeacherByName(username);
        return teacherService.getCoursePageCount(courseName, studentName, pageSize, teacher.getTeacherId());
    }

    @GetMapping("/grade/page/{index}")
    public Result getCoursesByPage(
            @PathVariable int index,
            @RequestParam("courseName") String courseName,
            @RequestParam("studentName") String studentName,
            @RequestParam("pageSize") int pageSize,
            @RequestHeader("Authorization") String authorizationHeader) {

        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Teacher teacher = teacherService.getTeacherByName(username);

        // 使用index参数进行分页计算
        int offset = (index - 1) * pageSize;
        int limit = pageSize;
        return teacherService.getCoursesByPage(courseName, studentName, offset, limit, teacher.getTeacherId());
    }

    // 选课表的course_select_id
    @GetMapping("/grade/{id}")
    public Result getGrade(@PathVariable int id) {
        Map<String, Object> select = teacherService.getGrade(id);
        return Result.success(select);
    }

    // 修改学生某门课的成绩
    @PutMapping("/grade")
    public Result updateGrade(@RequestBody Map<String, Object> entityForm) {
        int studentCourseId = (int) entityForm.get("studentCourseId");
        int score = (int) entityForm.get("score");

        return teacherService.updateGrade(studentCourseId, score);
    }

}
