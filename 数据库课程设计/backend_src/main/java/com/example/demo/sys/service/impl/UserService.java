package com.example.demo.sys.service.impl;

import com.example.demo.common.constant.UserType;
import com.example.demo.common.vo.Result;
import com.example.demo.sys.dao.dao.StudentDAO;
import com.example.demo.sys.dao.mapper.AdminMapper;
import com.example.demo.sys.dao.mapper.StudentMapper;
import com.example.demo.sys.dao.mapper.TeacherMapper;
import com.example.demo.sys.entity.Admin;
import com.example.demo.sys.entity.Student;
import com.example.demo.sys.entity.Teacher;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;

@Service
public class UserService implements UserDetailsService {
    @Autowired
    private StudentMapper studentMapper;
    @Autowired
    private TeacherMapper teacherMapper;
    @Autowired
    private AdminMapper adminMapper;

    public Result<Map<String, Object>> login(String username, String password, Integer userType, String token) {
        if (userType == UserType.STUDENT) {
            Student s = studentMapper.selectByName(username);
            if (s == null) {
                return Result.fail(-2, "⽆此⽤⼾或⽤⼾未激活");
            }
            if (!s.getPassword().equals(password)) {
                return Result.fail(-1, "密码错误");
            }
            // 生成token 将用户信息存入redis TODO:终极解决方案是JWT
            Map<String, Object> data = new HashMap<>();
            data.put("token", token);
            data.put("loggedIn", true);
            data.put("usertype", userType);
            data.put("username", username);
            return Result.success("登陆成功", data);
        } else if (userType == UserType.TEACHER) {
            Teacher s = teacherMapper.selectByName(username);
            if (s == null) {
                return Result.fail(-2, "⽆此⽤⼾或⽤⼾未激活");
            }
            if (!s.getPassword().equals(password)) {
                return Result.fail(-1, "密码错误");
            }
            // 生成token 将用户信息存入redis TODO:终极解决方案是JWT
            Map<String, Object> data = new HashMap<>();
            data.put("token", token);
            data.put("loggedIn", true);
            data.put("usertype", userType);
            data.put("username", username);
            return Result.success("登陆成功", data);
        } else if (userType == UserType.ADMIN) {
            Admin admin = adminMapper.selectByName(username);
            if (admin == null) {
                return Result.fail(-2, "⽆此⽤⼾或⽤⼾未激活");
            }
            if (!admin.getPassword().equals(password)) {
                return Result.fail(-1, "密码错误");
            }
            // 生成token 将用户信息存入redis TODO:终极解决方案是JWT
            Map<String, Object> data = new HashMap<>();
            data.put("token", token);
            data.put("loggedIn", true);
            data.put("usertype", userType);
            data.put("username", username);
            return Result.success("登陆成功", data);
        }
        return null;
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        return null;
    }


    public Result getLoginStatus(String username, Integer usertype) {
        if (username != null && (usertype == UserType.STUDENT || usertype == UserType.TEACHER || usertype == UserType.ADMIN)) {
            Map<String, Object> data = new HashMap<>();
            data.put("loggedIn", true);
            data.put("usertype", usertype);
            data.put("username", username);
            return Result.success("当前用户已登录",data);
        } else {
            Map<String, Object> data = new HashMap<>();
            data.put("loggedIn", false);
            return Result.fail(data);
        }
    }
}
