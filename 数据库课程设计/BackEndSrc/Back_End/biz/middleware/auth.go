package middleware

import (
	"Back_End/model"
	"Back_End/utils/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthAsStudent() gin.HandlerFunc {
	return func(context *gin.Context) {
		ck, err := context.Cookie("auth_session")

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "未认证的用户"
			context.JSON(http.StatusUnauthorized, resp)
			context.Abort()
			return
		}

		userInfo, err := redis.GetHash(ck)

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		userTypeStr, ok := userInfo["type"]

		if !ok {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.SetCookie("auth_session", "", -1, "/", "*", true, true)
			redis.UnsetKey(ck)
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		userType, err := strconv.Atoi(userTypeStr)

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.SetCookie("auth_session", "", -1, "/", "*", true, true)
			redis.UnsetKey(ck)
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		if userType != model.TypeStudent {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "错误的用户类型"
			context.JSON(http.StatusUnauthorized, resp)
			context.Abort()
			return
		}

		context.Next()

	}
}

func AuthAsTeacher() gin.HandlerFunc {
	return func(context *gin.Context) {
		ck, err := context.Cookie("auth_session")

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "未认证的用户"
			context.JSON(http.StatusUnauthorized, resp)
			context.Abort()
			return
		}

		userInfo, err := redis.GetHash(ck)

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		userTypeStr, ok := userInfo["type"]

		if !ok {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.SetCookie("auth_session", "", -1, "/", "*", true, true)
			redis.UnsetKey(ck)
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		userType, err := strconv.Atoi(userTypeStr)

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.SetCookie("auth_session", "", -1, "/", "*", true, true)
			redis.UnsetKey(ck)
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		if userType != model.TypeTeacher {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "错误的用户类型"
			context.JSON(http.StatusUnauthorized, resp)
			context.Abort()
			return
		}

		context.Next()

	}
}

func AuthAsAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		ck, err := context.Cookie("auth_session")

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "未认证的用户"
			context.JSON(http.StatusUnauthorized, resp)
			context.Abort()
			return
		}

		userInfo, err := redis.GetHash(ck)

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		userTypeStr, ok := userInfo["type"]

		if !ok {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.SetCookie("auth_session", "", -1, "/", "*", true, true)
			redis.UnsetKey(ck)
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		userType, err := strconv.Atoi(userTypeStr)

		if err != nil {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "鉴权服务失败"
			context.SetCookie("auth_session", "", -1, "/", "*", true, true)
			redis.UnsetKey(ck)
			context.JSON(http.StatusInternalServerError, resp)
			context.Abort()
			return
		}

		if userType != model.TypeAdmin {
			var resp model.JSONResp
			resp.Code = -1
			resp.Message = "错误的用户类型"
			context.JSON(http.StatusUnauthorized, resp)
			context.Abort()
			return
		}

		context.Next()

	}
}
