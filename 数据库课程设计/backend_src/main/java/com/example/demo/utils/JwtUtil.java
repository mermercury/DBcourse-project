package com.example.demo.utils;

import io.jsonwebtoken.*;
import org.springframework.stereotype.Component;

import javax.crypto.SecretKey;
import javax.crypto.spec.SecretKeySpec;
import java.util.*;

@Component
public class JwtUtil {
    // 7天过期
    private static long expire = 604800; // unit: second秒
    // 32位?密钥
    private static String secret = "iWFJSOhdjoabjwolzjrgbkaojndadddddijsbwlajdnwogjjjjjjskaliqoanjdfiWFJSOhdjoabjwolzjrgbkaojndadddddijsbwlajdnwogjjjjjjskaliqoanjdf";

    // 生成token
    public static String generateToken(String username, Integer usertype) {
        Date now = new Date();
        Date expiration = new Date(now.getTime() + 1000 * expire);

        Map<String, Object> claims = new HashMap<>();
        claims.put("username", username);
        claims.put("usertype", usertype);
        return Jwts.builder()
                .setHeaderParam("type", "JWT")
                .setClaims(claims)
                .setIssuedAt(now)           // 生效时间
                .setExpiration(expiration)  // 过期时间
                .signWith(SignatureAlgorithm.HS512, secret)
                .compact();
    }

    // 解析token
    public static Jws<Claims> getClaimsByToken(String token) {
        return Jwts.parser()
                .setSigningKey(secret)
                .parseClaimsJws(token);
    }

    public static String extractTokenFromAuthorizationHeader(String authorizationHeader) {
        // 假设传递的token是以"Bearer "开头的
        // 你可能需要根据具体的token传递方式来提取实际的token值
        if (authorizationHeader != null && authorizationHeader.startsWith("Bearer ")) {
            return authorizationHeader.substring(7); // 去掉"Bearer "前缀
        }
        return null; // 或者抛出异常，视情况而定
    }
}

