package com.example.demo.sys.entity;

import lombok.Data;

@Data
public class CustomStudentCourse {
    // 学生想选修这门课程
    private Integer courseId;
    private String courseName;
    private String teacherName;
    private Integer credit;
    private String time; // 上课时间
    private int selectedCount;
    private int maxSize;
}
