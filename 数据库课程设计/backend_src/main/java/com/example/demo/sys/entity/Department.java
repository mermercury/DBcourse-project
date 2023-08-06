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
public class Department implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 系ID
     */
    @TableId(value = "department_id", type = IdType.AUTO)
    private Integer departmentId;

    /**
     * 系名
     */
    private String departmentName;

    /**
     * 专业数
     */
    private Integer majorCount;

    /**
     * 教师数
     */
    private Integer teacherCount;

    public Integer getDepartmentId() {
        return departmentId;
    }

    public void setDepartmentId(Integer departmentId) {
        this.departmentId = departmentId;
    }
    public String getDepartmentName() {
        return departmentName;
    }

    public void setDepartmentName(String departmentName) {
        this.departmentName = departmentName;
    }
    public Integer getMajorCount() {
        return majorCount;
    }

    public void setMajorCount(Integer majorCount) {
        this.majorCount = majorCount;
    }
    public Integer getTeacherCount() {
        return teacherCount;
    }

    public void setTeacherCount(Integer teacherCount) {
        this.teacherCount = teacherCount;
    }

    @Override
    public String toString() {
        return "Department{" +
            "departmentId=" + departmentId +
            ", departmentName=" + departmentName +
            ", majorCount=" + majorCount +
            ", teacherCount=" + teacherCount +
        "}";
    }
}
