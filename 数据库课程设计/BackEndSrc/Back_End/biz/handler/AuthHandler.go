package handler

import (
	"Back_End/biz/service"
	"Back_End/conf"
	"Back_End/model"
	"Back_End/utils/redis"
	"Back_End/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleLogin(ctx *gin.Context) {
	fmt.Println(ctx.Request.Method)
	var resp model.JSONResp
	var loginReq model.LoginInfo
	ck, err := ctx.Cookie("auth_session")
	var info map[string]string
	var infoNotFound error
	if err == nil {
		info, infoNotFound = redis.GetHash(ck)
	} else {
		fmt.Println("not found cookie")
		uid := uuid.New()
		uidStr := uid.String()
		// ctx.SetCookie("auth_session", uidStr, 3600, "/", "", false, false)
		// ctx.SetSameSite(http.SameSiteLaxMode)
		http.SetCookie(ctx.Writer, &http.Cookie{
			Name:     "auth_session", //你的cookie的名字
			Value:    uidStr,         //cookie值
			Path:     "/",
			Domain:   "139.9.143.161",
			MaxAge:   3600,
			Secure:   false,
			HttpOnly: false,
			SameSite: 2, //下面是详细解释
		})
		ck = uidStr
		info = map[string]string{
			"valid": "false",
		}
	}
	if err != nil || infoNotFound != nil {
		redis.SetHash(ck, info, time.Hour*3)
	}

	if status, ok := info["valid"]; ok && status == redis.TRUE {
		response.SetJSONResp(&resp, 0, "登录成功", nil)
		ctx.JSON(http.StatusOK, resp)
		return
	}

	err = ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		response.SetJSONResp(&resp, -3, "请求信息非法", nil)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if err != nil {
		response.SetJSONResp(&resp, -3, "请求信息非法", nil)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if err != nil || loginReq.Type > 3 || loginReq.Type < 1 {
		response.SetJSONResp(&resp, -3, "请求信息非法", nil)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	status, userinfo := service.ServeLogin(uint(loginReq.UserID), loginReq.Password, model.UserType(loginReq.Type))
	// status := 0
	switch status {
	case 0:
		err := redis.UpdateHash(ck, conf.RedisValidField, redis.TRUE)
		err = redis.UpdateHash(ck, conf.RedisIDField, strconv.Itoa(int(userinfo.GetID())))
		err = redis.UpdateHash(ck, conf.RedisNameField, userinfo.GetName())
		err = redis.UpdateHash(ck, conf.RedisTypeField, strconv.Itoa(int(userinfo.GetType())))
		err = redis.UpdateHash(ck, conf.RedisPermField, strconv.Itoa(userinfo.GetPerm()))
		if err != nil {
			response.SetJSONResp(&resp, -4, "状态更新失败", nil)
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}
		response.SetJSONResp(&resp, 0, "登录成功", nil)
		ctx.JSON(http.StatusOK, resp)
		return
	case -1:
		response.SetJSONResp(&resp, -1, "密码错误", nil)
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	case -2:
		response.SetJSONResp(&resp, -2, "无此用户或用户未激活", nil)
		ctx.JSON(http.StatusNotFound, resp)
		return
	default:
		response.SetJSONResp(&resp, -3, "内部错误", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
}

func HandleLogout(ctx *gin.Context) {
	var resp model.JSONResp
	ck, err := ctx.Cookie("auth_session")
	var info map[string]string
	var infoNotFound error
	if err == nil {
		info, infoNotFound = redis.GetHash(ck)
	}
	if err != nil || infoNotFound != nil {
		response.SetJSONResp(&resp, -1, "用户未登录", nil)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if status, ok := info["valid"]; ok && status == redis.FALSE {
		response.SetJSONResp(&resp, -1, "用户未登录", nil)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx.SetCookie("auth_session", "", -1, "/", ctx.GetHeader("origin"), true, true)
	err = redis.UnsetKey(ck)
	if err != nil {
		response.SetJSONResp(&resp, -2, "状态更新失败", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	response.SetJSONResp(&resp, 0, "登出成功", nil)
	ctx.JSON(http.StatusOK, resp)
}

func HandleGetStatus(ctx *gin.Context) {
	var resp model.JSONResp
	ck, err := ctx.Cookie("auth_session")
	var info map[string]string
	if err != nil {
		data := model.LoginStatusResp{
			LoggedIn:   false,
			UserId:     "-1",
			Username:   "null",
			Usertype:   -1,
			Permission: -1,
		}
		resp.Code = -1
		resp.Message = "未登录"
		resp.Data = data
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	info, infoNotFound := redis.GetHash(ck)

	if infoNotFound != nil || info["valid"] == redis.FALSE {
		data := model.LoginStatusResp{
			LoggedIn:   false,
			UserId:     "-1",
			Username:   "null",
			Usertype:   -1,
			Permission: -1,
		}
		resp.Code = -1
		resp.Message = "未登录"
		resp.Data = data
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	resp.Code = 0
	resp.Message = "已登录"
	userType, _ := strconv.Atoi(info[conf.RedisTypeField])
	userPerm, _ := strconv.Atoi(info[conf.RedisPermField])
	data := model.LoginStatusResp{
		LoggedIn:   info["valid"] == redis.TRUE,
		UserId:     info[conf.RedisIDField],
		Username:   info[conf.RedisNameField],
		Usertype:   userType,
		Permission: userPerm,
	}

	// data := model.LoginStatusResp{
	// 	LoggedIn: true,
	// 	UserId:   "1",
	// 	Username: "无鉴权学生test",
	// 	Usertype: 1,
	// 	Permission: 0,
	// }
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
}
