package com.example.demo.sys.dao.dao;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.example.demo.common.constant.UserType;
import com.example.demo.sys.dao.mapper.StudentMapper;
import com.example.demo.sys.entity.Course;
import com.example.demo.sys.entity.Student;
import com.example.demo.sys.entity.UserEntity;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class StudentDAO {
    public static final int PAGE_SIZE = 6;
    @Autowired
    private StudentMapper mapper;

    // TODO: Autowired就行 但下面这样写就不行...
//    public StudentDAO(StudentMapper mapper) {
//        this.mapper = mapper;
//        System.out.println("aaa+ " + mapper);
//    }

    public UserEntity checkLogin(String name) {
        Student s = getByName(name);
        if (s == null) return null;
        return new UserEntity(name, s.getPassword(), UserType.STUDENT);
    }

    // 查询所选课程
//    @Select("SELECT department_name FROM department WHERE department_id = #{departmentId}")
//    public List<Course> selectAllCourses(String stuId) {
//        mapper.se
//    }

    public Student getByName(String name) {
        LambdaQueryWrapper<Student> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Student::getStudentName, name);
        return this.mapper.selectOne(wrapper);
    }


    public int insert(Student entity) {
        return mapper.insert(entity);
    }

    public int delete(Integer id) {
        return mapper.deleteById(id);
    }

    public Student get(Integer id) {
        return mapper.selectById(id);
    }

    public int update(Student entity) {
        return mapper.updateById(entity);
    }

    public StudentDAO() {
    }

    //    public int count(String majorName, String className, String name) {
//        return mapper.count(majorName, className, name);
//    }
//
//    public List<StudentItemVO> getPage(Integer index, String majorName, String className, String name) {
//        Page<StudentItemVO> page = new Page<>(index, PAGE_SIZE);
//
//        return mapper.getPage(page, majorName, className, name).getRecords();
//    }
//
//    public Integer countByClassId(Integer id) {
//        LambdaQueryWrapper<Student> wrapper = new LambdaQueryWrapper<>();
//        wrapper.eq(Student::getClassId, id);
//
//        return mapper.selectCount(wrapper);
//    }
//
//    public List<Student> listName() {
//        LambdaQueryWrapper<Student> wrapper = new LambdaQueryWrapper<>();
//        wrapper.select(Student::getId, Student::getName);
//
//        return mapper.selectList(wrapper);
//    }
//
//    public Integer getDepartmentIdById(Integer studentId) {
//        return mapper.getDepartmentIdById(studentId);
//    }
//
//    public Integer getGradeById(Integer studentId) {
//        return mapper.getGradeById(studentId);
//    }
//
//    public StudentInfoVO getStudentInfoById(Integer studentId) {
//        return mapper.getStudentInfoById(studentId);
//    }
}
