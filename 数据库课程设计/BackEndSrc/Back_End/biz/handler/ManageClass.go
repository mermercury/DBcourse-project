package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetClassPageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetClassReq
	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.PageSize <= 0 {
		resp.Code = -1
		resp.Message = "参数错误: pageSize 不应该小于1"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	class := mysql.Operator.Class
	classOp := class.Where()

	if req.DepartmentName != "" {
		classOp = classOp.Where(class.DepartmentName.Eq(req.DepartmentName))
	}

	if req.MajorName != "" {
		classOp = classOp.Where(class.MajorName.Eq(req.MajorName))
	}

	if req.ClassName != "" {
		classOp = classOp.Where(class.ClassName.Eq(req.ClassName))
	}

	total, err := classOp.Count()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	ans := int(total) / req.PageSize

	if int(total)%req.PageSize > 0 {
		ans++
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = model.CountResp{
		Count: ans,
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetClassPage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetClassReq

	indexStr := ctx.Param(conf.ParamIndex)
	index, err := strconv.Atoi(indexStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.PageSize <= 0 {
		resp.Code = -1
		resp.Message = "参数错误: pageSize 不应该小于1"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if index > 0 {
		index -= 1
	}

	class := mysql.Operator.Class
	classOp := class.Where()

	if req.DepartmentName != "" {
		classOp = classOp.Where(class.DepartmentName.Eq(req.DepartmentName))
	}

	if req.MajorName != "" {
		classOp = classOp.Where(class.MajorName.Eq(req.MajorName))
	}

	if req.ClassName != "" {
		classOp = classOp.Where(class.ClassName.Eq(req.ClassName))
	}

	res, err := classOp.Offset(index * req.PageSize).Limit(req.PageSize).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = res
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetClassById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	class := mysql.Operator.Class
	res, err := class.Where(class.ClassID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此专业"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = *res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateClass(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.ClassInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	class := mysql.Operator.Class

	res, err := class.Where(class.ClassID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"class_name":      req.ClassName,
		"major_name":      req.MajorName,
		"grade":           req.Grade,
		"department_name": req.DepartmentName,
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "修改失败:没有这个专业"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateClass(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.ClassInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Class{
		ClassName:      req.ClassName,
		DepartmentName: req.DepartmentName,
		MajorName:      req.MajorName,
		Grade:          uint(req.Grade),
	}
	err = mysql.Operator.Class.Create(&toAdd)

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

func DeleteClass(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	class := mysql.Operator.Class
	classCheck, err := class.Where(class.ClassID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(classCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此专业"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := class.Where(class.ClassID.Eq(uint(id))).Delete()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "删除失败: 没有这个专业"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetClassList(ctx *gin.Context) {
	var nameList []model.NameListItem
	var resp model.JSONResp
	classes, err := mysql.Operator.Class.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range classes {
		nameList = append(nameList, model.NameListItem{
			ID:   int(v.ClassID),
			Name: v.ClassName,
		})
	}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = nameList
	ctx.JSON(http.StatusOK, resp)
	return
}
