package com.example.demo.sys.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import java.io.Serializable;

/**
 * <p>
 * 
 * </p>
 *
 * @author baomidou
 * @since 2023-07-23
 */
//@TableName("course_select")
public class CourseSelect implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 选课记录id
     */
    @TableId(value = "course_select_id", type = IdType.AUTO)
    private Integer courseSelectId;

    /**
     * 学生id
     */
    private Integer studentId;

    /**
     * 课程id
     */
    private Integer courseId;

    /**
     * 成绩
     */
    private Integer courseScore;

    /**
     * 评分
     */
    private Integer evaluateScore;

    /**
     * 课程评价
     */
    private String evaluation;

    public Integer getCourseSelectId() {
        return courseSelectId;
    }

    public void setCourseSelectId(Integer courseSelectId) {
        this.courseSelectId = courseSelectId;
    }

    public Integer getStudentId() {
        return studentId;
    }

    public void setStudentId(Integer studentId) {
        this.studentId = studentId;
    }
    public Integer getCourseId() {
        return courseId;
    }

    public void setCourseId(Integer courseId) {
        this.courseId = courseId;
    }
    public Integer getCourseScore() {
        return courseScore;
    }

    public void setCourseScore(Integer courseScore) {
        this.courseScore = courseScore;
    }
    public Integer getEvaluateScore() {
        return evaluateScore;
    }

    public void setEvaluateScore(Integer evaluateScore) {
        this.evaluateScore = evaluateScore;
    }
    public String getEvaluation() {
        return evaluation;
    }

    public void setEvaluation(String evaluation) {
        this.evaluation = evaluation;
    }

    @Override
    public String toString() {
        return "CourseSelect{" +
            "courseSelectId=" + courseSelectId +
            ", studentId=" + studentId +
            ", courseId=" + courseId +
            ", courseScore=" + courseScore +
            ", evaluateScore=" + evaluateScore +
            ", evaluation=" + evaluation +
        "}";
    }
}
