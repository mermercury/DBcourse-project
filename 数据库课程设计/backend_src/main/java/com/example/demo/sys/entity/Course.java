package com.example.demo.sys.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;

/**
 * <p>
 * 
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
public class Course implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 课程id
     */
    @TableId(value = "course_id", type = IdType.AUTO)
    private Integer courseId;

    /**
     * 课程名
     */
    private String courseName;

    /**
     * 年级
     */
    private Integer grade;

    /**
     * 授课教师id
     */
    private Integer teacherId;

    /**
     * 系id
     */
    private Integer departmentId;

    /**
     * 学分
     */
    private Integer credit;

    /**
     * 上课时间
     */
    private String courseTime;

    /**
     * 上课地点
     */
    private String location;

    /**
     * 已选人数
     */
    private Integer selected;

    /**
     * 容量
     */
    private Integer size;

    /**
     * 考试日期
     */
    private String examDate;

    /**
     * 考试地点
     */
    private String examLoc;

    public Integer getCourseId() {
        return courseId;
    }

    public void setCourseId(Integer courseId) {
        this.courseId = courseId;
    }
    public String getCourseName() {
        return courseName;
    }

    public void setCourseName(String courseName) {
        this.courseName = courseName;
    }
    public Integer getGrade() {
        return grade;
    }

    public void setGrade(Integer grade) {
        this.grade = grade;
    }
    public Integer getTeacherId() {
        return teacherId;
    }

    public void setTeacherId(Integer teacherId) {
        this.teacherId = teacherId;
    }
    public Integer getDepartmentId() {
        return departmentId;
    }

    public void setDepartmentId(Integer departmentId) {
        this.departmentId = departmentId;
    }
    public Integer getCredit() {
        return credit;
    }

    public void setCredit(Integer credit) {
        this.credit = credit;
    }
    public String getCourseTime() {
        return courseTime;
    }

    public void setCourseTime(String courseTime) {
        this.courseTime = courseTime;
    }
    public String getLocation() {
        return location;
    }

    public void setLocation(String location) {
        this.location = location;
    }
    public Integer getSelected() {
        return selected;
    }

    public void setSelected(Integer selected) {
        this.selected = selected;
    }
    public Integer getSize() {
        return size;
    }

    public void setSize(Integer size) {
        this.size = size;
    }
    public String getExamDate() {
        return examDate;
    }

    public void setExamDate(String examDate) {
        this.examDate = examDate;
    }
    public String getExamLoc() {
        return examLoc;
    }

    public void setExamLoc(String examLoc) {
        this.examLoc = examLoc;
    }

    @Override
    public String toString() {
        return "Course{" +
            "courseId=" + courseId +
            ", courseName=" + courseName +
            ", grade=" + grade +
            ", teacherId=" + teacherId +
            ", departmentId=" + departmentId +
            ", credit=" + credit +
            ", courseTime=" + courseTime +
            ", location=" + location +
            ", selected=" + selected +
            ", size=" + size +
            ", examDate=" + examDate +
            ", examLoc=" + examLoc +
        "}";
    }
}
