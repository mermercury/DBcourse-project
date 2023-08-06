package com.example.demo.sys.controller;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.*;
import com.example.demo.sys.service.IStudentService;
import com.example.demo.utils.JwtUtil;
import org.springframework.validation.annotation.Validated;
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
@RequestMapping("/api/v1/student")
public class StudentController {
    private IStudentService studentService;

    public StudentController(IStudentService studentService) {
        this.studentService = studentService;
    }

    @GetMapping("/info")
        public Result<Map<String, Object>> getStudentInfo(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        return this.studentService.getStudentInfo(username);
    }

    @PutMapping("/info")
    public Result updateStudentInfo(@RequestBody Student newstudent) {
        return this.studentService.updateStudentInfo(newstudent);
    }

    // 查询所选课程
    @GetMapping("/course")
    public Result<List<CustomCourseSelect>> getStudentCourse(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Student s = studentService.getStudentByName(username);
        List<CustomCourseSelect> list = studentService.getCourseSelect(s.getStudentId());
        if (list != null) {
            return Result.success(list);
        } else {
            return Result.fail("当前尚未选课！");
        }
    }

    @GetMapping("/course/page/count")
    public Result getCoursePageCount(
            @RequestParam("courseName") String courseName,
            @RequestParam("teacherName") String teacherName,
            @RequestParam("pageSize") int pageSize
    ) {
        return studentService.getCoursePageCount(courseName, teacherName, pageSize);
    }

    @GetMapping("/course/page/{index}")
    public Result<List<CustomStudentCourse>> getCoursesByPage(
            @PathVariable int index,
            @RequestParam("courseName") String courseName,
            @RequestParam("teacherName") String teacherName,
            @RequestParam("pageSize") int pageSize) {

        // 使用index参数进行分页计算
        int offset = (index - 1) * pageSize;
        int limit = pageSize;

        return studentService.getCoursesByPage(courseName, teacherName, offset, limit);
    }

    @PostMapping("/course/select")
    public Result selectCourse(@RequestBody Map<String, Object> requestBody, @RequestHeader("Authorization") String authorizationHeader) {
        Integer courseId = (Integer) requestBody.get("courseId");
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Student s = studentService.getStudentByName(username);
        Integer stuId = s.getStudentId();
        return studentService.selectCourse(stuId, courseId);
    }

    // 学生退课
    @DeleteMapping("/course/{studentCourseId}")
    public Result deleteCourse(@PathVariable int studentCourseId, @RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Student s = studentService.getStudentByName(username);
        Integer stuId = s.getStudentId();
        return studentService.deleteCourse(stuId, studentCourseId);
    }

    @GetMapping("/exam")
    public Result<List<CustomCourseExamInfo>> getExamList(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Student s = studentService.getStudentByName(username);
        return studentService.getExamList(s.getStudentId());
    }

    @GetMapping("/timetable")
    public Result getStudentTimetable(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Student s = studentService.getStudentByName(username);
        return studentService.getStudentTimetable(s.getStudentId());
    }
}
