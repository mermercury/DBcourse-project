package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMajorPageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetMajorsReq
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



	major := mysql.Operator.Major
	majorOp := major.Where()

	if req.MajorName != "" {
		majorOp = majorOp.Where(major.MajorName.Eq(req.MajorName))
	}

	if req.DepartmentName != "" {
		majorOp = majorOp.Where(major.DepartmentName.Eq(req.DepartmentName))
	}

	total, err := majorOp.Count()

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

func GetMajorPage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetMajorsReq

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

	major := mysql.Operator.Major
	majorOp := major.Where()

	if req.MajorName != "" {
		majorOp = majorOp.Where(major.MajorName.Eq(req.MajorName))
	}

	if req.DepartmentName != "" {
		majorOp = majorOp.Where(major.DepartmentName.Eq(req.DepartmentName))
	}

	res, err := majorOp.Offset(index * req.PageSize).Limit(req.PageSize).Find()

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

func GetMajorById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	major := mysql.Operator.Major
	res, err := major.Where(major.MajorID.Eq(uint(id))).Find()

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

func UpdateMajor(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.MajorInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	major := mysql.Operator.Major

	res, err := major.Where(major.MajorID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"major_name":      req.MajorName,
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

func CreateMajor(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.MajorInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Major{
		MajorName:      req.MajorName,
		DepartmentName: req.DepartmentName,
	}

	department := mysql.Operator.Department
	
	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {
		res, err := department.Where(department.DepartmentName.Eq(toAdd.DepartmentName)).Update(department.MajorCount, department.MajorCount.Add(1))
		
		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such department")
		}

		err = mysql.Operator.Major.Create(&toAdd)

		return err
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "添加失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK,resp)
	return
}

func DeleteMajor(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	major := mysql.Operator.Major
	department := mysql.Operator.Department
	majorCheck, err := major.Where(major.MajorID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(majorCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此专业"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toDelete := majorCheck[0]

	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {
		res, err := department.Where(department.DepartmentName.Eq(toDelete.DepartmentName)).Update(department.MajorCount, department.MajorCount.Sub(1))
		
		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such major")
		}

		res, err = major.Where(major.MajorID.Eq(toDelete.MajorID)).Delete()

		if res.RowsAffected <= 0 {
			return errors.New("no such major")
		}

		return err
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "删除失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetMajorList(ctx *gin.Context) {
	var nameList []model.NameListItem
	var resp model.JSONResp
	majors, err := mysql.Operator.Major.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range majors {
		nameList = append(nameList, model.NameListItem{
			ID:   int(v.MajorID),
			Name: v.MajorName,
		})
	}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = nameList
	ctx.JSON(http.StatusOK, resp)
	return
}
