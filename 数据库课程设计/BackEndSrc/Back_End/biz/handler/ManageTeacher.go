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

func GetTeacherPageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetTeachersReq
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

	teacher := mysql.Operator.Teacher
	teacherOp := teacher.Where()

	if req.DepartmentName != "" {
		teacherOp = teacherOp.Where(teacher.DepartmentName.Eq(req.DepartmentName))
	}

	if req.TeacherName != "" {
		teacherOp = teacherOp.Where(teacher.TeacherName.Eq(req.TeacherName))
	}

	total, err := teacherOp.Count()

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

func GetTeacherPage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetTeachersReq

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

	teacher := mysql.Operator.Teacher
	teacherOp := teacher.Where()

	if req.DepartmentName != "" {
		teacherOp = teacherOp.Where(teacher.DepartmentName.Eq(req.DepartmentName))
	}

	if req.TeacherName != "" {
		teacherOp = teacherOp.Where(teacher.TeacherName.Eq(req.TeacherName))
	}

	res, err := teacherOp.Offset(index * req.PageSize).Limit(req.PageSize).Find()

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

func GetTeacherById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	teacher := mysql.Operator.Teacher
	res, err := teacher.Where(teacher.TeacherID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此老师"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = *res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateTeacher(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.TeacherInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	Teacher := mysql.Operator.Teacher

	res, err := Teacher.Where(Teacher.TeacherID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"teacher_name":    req.TeacherName,
		"phone":           req.Phone,
		"department_name": req.DepartmentName,
		"password":        req.Password,
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "修改失败:没有这个老师"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateTeacher(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.TeacherInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Teacher{
		Phone:          req.Phone,
		TeacherName:    req.TeacherName,
		DepartmentName: req.DepartmentName,
		Password:       req.Password,
	}
	
	department := mysql.Operator.Department
	

	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {
		res, err := department.
			Where(department.DepartmentName.
				Eq(toAdd.DepartmentName)).
			Update(department.TeacherCount, department.TeacherCount.Add(1))
		if err != nil {
			return nil
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such department")
		}

		err = mysql.Operator.Teacher.Create(&toAdd)

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
	ctx.JSON(http.StatusOK, resp)
	return
}

func DeleteTeacher(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	teacher := mysql.Operator.Teacher
	department := mysql.Operator.Department
	teacherCheck, err := teacher.Where(teacher.TeacherID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(teacherCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此老师"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toDelete := teacherCheck[0]

	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {
		res, err := department.
			Where(department.DepartmentName.
				Eq(toDelete.DepartmentName)).
			Update(department.TeacherCount, department.TeacherCount.Sub(1))
		if err != nil {
			return nil
		}

		res, err = teacher.Where(teacher.TeacherID.Eq(toDelete.TeacherID)).Delete()

		if res.RowsAffected <= 0 {
			return errors.New("no such teacher")
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

func GetTeacherList(ctx *gin.Context) {
	var nameList []model.NameListItem
	var resp model.JSONResp
	teachers, err := mysql.Operator.Teacher.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range teachers {
		nameList = append(nameList, model.NameListItem{
			ID:   int(v.TeacherID),
			Name: v.TeacherName,
		})
	}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = nameList
	ctx.JSON(http.StatusOK, resp)
	return
}
