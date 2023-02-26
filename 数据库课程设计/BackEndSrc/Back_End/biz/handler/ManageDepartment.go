package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDepartmentPageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetDepartmentsReq
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

	department := mysql.Operator.Department
	departmentOp := department.Where()

	if req.DepartmentName != "" {
		departmentOp = departmentOp.Where(department.DepartmentName.Eq(req.DepartmentName))
	}

	total, err := departmentOp.Count()

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

func GetDepartmentPage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetDepartmentsReq

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

	department := mysql.Operator.Department
	departmentOp := department.Where()

	if req.DepartmentName != "" {
		departmentOp = departmentOp.Where(department.DepartmentName.Eq(req.DepartmentName))
	}

	res, err := departmentOp.Offset(index * req.PageSize).Limit(req.PageSize).Find()

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

func GetDepartmentById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	department := mysql.Operator.Department
	res, err := department.Where(department.DepartmentID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此学院"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = *res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateDepartment(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.DepartmentInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	department := mysql.Operator.Department

	res, err := department.Where(department.DepartmentID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"department_name": req.DepartmentName,
		"major_count":     req.MajorCount,
		"teacher_count":   req.TeacherCount,
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "修改失败:没有这个学院"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateDepartment(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.DepartmentInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Department{
		DepartmentName: req.DepartmentName,
	}
	err = mysql.Operator.Department.Create(&toAdd)

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

func DeleteDepartment(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	department := mysql.Operator.Department
	departmentCheck, err := department.Where(department.DepartmentID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(departmentCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此学院"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := department.Where(department.DepartmentID.Eq(uint(id))).Delete()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "删除失败: 没有这个学院"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetDepartmentList(ctx *gin.Context) {
	var nameList []model.NameListItem
	var resp model.JSONResp
	departments, err := mysql.Operator.Department.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range departments {
		nameList = append(nameList, model.NameListItem{
			ID:   int(v.DepartmentID),
			Name: v.DepartmentName,
		})
	}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = nameList
	ctx.JSON(http.StatusOK, resp)
	return
}
