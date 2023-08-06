package com.example.demo.sys.entity;

import lombok.Data;

@Data
public class CustomStudentTimetable {
    // 学生已选课程的课程表
    private String courseName;
    private String teacherName;
    private String location;
    private String courseTime;
}
