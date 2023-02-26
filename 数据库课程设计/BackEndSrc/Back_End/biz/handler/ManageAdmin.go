package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAdminById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	admin := mysql.Operator.Admin
	res, err := admin.Where(admin.AdminID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此管理员"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = *res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateAdmin(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.AdminInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	admin := mysql.Operator.Admin

	res, err := admin.Where(admin.AdminID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"admin_name": req.AdminName,
		"privilege":  req.Privilege,
		"password":   req.Password,
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "修改失败:没有这个管理员"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateAdmin(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.AdminInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Admin{
		AdminName: req.AdminName,
		Privilege: req.Privilege,
		Password:  req.Password,
	}
	err = mysql.Operator.Admin.Create(&toAdd)

	if err != nil {
		resp.Code = -1
		resp.Message = "添加失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func DeleteAdmin(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	admin := mysql.Operator.Admin
	adminCheck, err := admin.Where(admin.AdminID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(adminCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此管理员"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := admin.Where(admin.AdminID.Eq(uint(id))).Delete()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "删除失败: 没有这个管理员"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetAdminList(ctx *gin.Context) {
	var resp model.JSONResp
	adminss, err := mysql.Operator.Admin.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = adminss
	ctx.JSON(http.StatusOK, resp)
	return
}
