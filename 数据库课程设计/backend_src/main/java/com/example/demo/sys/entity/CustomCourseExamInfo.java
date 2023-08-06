package com.example.demo.sys.entity;

import lombok.Data;

@Data
public class CustomCourseExamInfo {
    // 学生已选课程的考试信息
    private Integer studentCourseId;
    private String courseName;
    private String teacherName;
    private String examDate;
    private String examLocation;
}