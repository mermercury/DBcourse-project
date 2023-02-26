SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
CREATE DATABASE IF NOT EXISTS CourseDB;
USE CourseDB;


CREATE TABLE IF NOT EXISTS department
(
    department_id   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '系ID',
    department_name VARCHAR(256) UNIQUE COMMENT '系名',
    major_count     INT COMMENT '专业数',
    teacher_count   INT COMMENT '教师数',
    CHECK ( major_count >= 0 ),
    CHECK ( teacher_count >= 0 )
);

CREATE TABLE IF NOT EXISTS major
(
    major_id        INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '专业ID',
    major_name      VARCHAR(256) UNIQUE COMMENT '专业名',
    department_name VARCHAR(256) COMMENT '所属系名',
    FOREIGN KEY (department_name) REFERENCES department (department_name) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS class
(
    class_id        INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '班级ID',
    class_name      VARCHAR(256) NOT NULL UNIQUE COMMENT '班级名',
    grade           INT UNSIGNED NOT NULL COMMENT '年级',
    department_name VARCHAR(255) COMMENT '系名',
    major_name      VARCHAR(255),
    FOREIGN KEY (department_name) REFERENCES department (department_name) ON UPDATE CASCADE ON DELETE SET NULL ,
    FOREIGN KEY (major_name) REFERENCES major (major_name) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS student
(
    student_id      INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '学号',
    student_name    VARCHAR(256) COMMENT '学生姓名',
    department_name VARCHAR(256) COMMENT '系名',
    major_name      VARCHAR(256) COMMENT '专业名',
    class_name      VARCHAR(256) COMMENT '班级名',
    email           VARCHAR(64) NOT NULL COMMENT '邮箱',
    birthday        VARCHAR(128) COMMENT '生日',
    sex             INT UNSIGNED COMMENT '性别',
    password        VARCHAR(32) NOT NULL COMMENT '密码',
    FOREIGN KEY (major_name) REFERENCES major (major_name) ON UPDATE CASCADE ON DELETE SET NULL ,
    FOREIGN KEY (department_name) REFERENCES department (department_name) ON UPDATE CASCADE ON DELETE SET NULL,
    FOREIGN KEY (class_name) REFERENCES class (class_name) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS teacher
(
    teacher_id      INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '老师工号',
    teacher_name    VARCHAR(256) NOT NULL COMMENT '老师名字',
    department_name VARCHAR(256) COMMENT '所属系号',
    phone           VARCHAR(11)  COMMENT '电话号码',
    password        VARCHAR(32)  NOT NULL COMMENT '密码',
    FOREIGN KEY (department_name) REFERENCES department (department_name) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS admin
(
    admin_id   INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '管理员id',
    admin_name VARCHAR(256) COMMENT '管理员用户名',
    privilege  INT         NOT NULL COMMENT '管理员权限',
    password   VARCHAR(32) NOT NULL COMMENT '密码'
);

CREATE TABLE IF NOT EXISTS course
(
    course_id       INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '课程id',
    course_name     VARCHAR(256) NOT NULL COMMENT '课程名',
    grade           INT          NOT NULL COMMENT '年级',
    teacher_id      INT UNSIGNED COMMENT '授课教师id',
    department_name VARCHAR(256) COMMENT '系名',
    credit          INT UNSIGNED NOT NULL COMMENT '学分',
    course_time     VARCHAR(32) COMMENT '上课时间',
    location        VARCHAR(256) COMMENT '上课地点',
    selected        INT UNSIGNED NOT NULL COMMENT '已选人数' DEFAULT 0,
    size            INT UNSIGNED NOT NULL COMMENT '容量',
    exam_date       VARCHAR(256) NOT NULL COMMENT '考试日期',
    exam_loc        VARCHAR(256) NOT NULL COMMENT '考试地点',
    FOREIGN KEY (teacher_id) REFERENCES teacher (teacher_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (department_name) REFERENCES department (department_name) ON UPDATE CASCADE ON DELETE SET NULL,
    CHECK ( selected <= size AND selected >= 0)
);

CREATE TABLE IF NOT EXISTS course_select
(
    course_select_id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT '选课记录id',
    student_id       INT UNSIGNED NOT NULL COMMENT '学生id',
    course_id        INT UNSIGNED NOT NULL COMMENT '课程id',
    course_score     INT COMMENT '成绩' DEFAULT -1,
    evaluate_score   INT COMMENT '评分' DEFAULT -1,
    evaluation       TEXT COMMENT '课程评价',
    UNIQUE (course_id, student_id),
    FOREIGN KEY (student_id) REFERENCES student (student_id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES course (course_id) ON UPDATE CASCADE ON DELETE CASCADE
);

