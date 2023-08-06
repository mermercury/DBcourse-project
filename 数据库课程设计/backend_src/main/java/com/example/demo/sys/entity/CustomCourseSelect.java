package com.example.demo.sys.entity;

import lombok.Data;

@Data
public class CustomCourseSelect {
    // 学生已选课程
    private Integer studentCourseId;
    private String courseName;
    private String teacherName;
    private Integer credit;
    private Integer score;
}
