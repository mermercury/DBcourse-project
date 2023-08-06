package com.example.demo.sys.controller;

import com.example.demo.common.vo.Result;
import com.example.demo.sys.entity.UserEntity;
import com.example.demo.sys.service.impl.UserService;
import com.example.demo.utils.JwtUtil;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

@RequestMapping("/api/v1")
@RestController
public class UserController {
//    @Autowired
    private final UserService service;

    public UserController(UserService service) {
        this.service = service;
    }

    @PostMapping("/login")
    public Result<Map<String, Object>> login(@Validated @RequestBody UserEntity user) {
        String username = user.getUsername();
        String password = user.getPassword();
        Integer usertype = user.getUsertype();
        String token = JwtUtil.generateToken(username, usertype);
        System.out.println("为用户：" + username + " 生成的token：" + token);
        return service.login(username, password, usertype, token);
    }

    @GetMapping("/status")
    public Result getLoginStatus(@RequestHeader("Authorization") String authorizationHeader) {
        // 解析token
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Integer usertype = (Integer) claims.get("usertype");
        return service.getLoginStatus(username, usertype);
    }


    @PostMapping("/logout")
    public Result logout(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        if (token == null || token.equals("")) {
            return Result.fail(-1, "用户未登录");
        }
        return Result.success("登出成功！");
    }

    @GetMapping("/welcome")
    public Result getUserName(@RequestHeader("Authorization") String authorizationHeader) {
        String token = JwtUtil.extractTokenFromAuthorizationHeader(authorizationHeader);
        Map<String, Object> claims = JwtUtil.getClaimsByToken(token).getBody();
        String username = (String) claims.get("username");
        Integer usertype = (Integer) claims.get("usertype");
        Map data = new HashMap<>();
        data.put("userName", username);
        data.put("userType", usertype);
        return Result.success(data);
    }
}
