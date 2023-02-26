package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudentPageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetStudentsReq
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

	student := mysql.Operator.Student
	studentOp := student.Where()

	if req.MajorName != "" {
		studentOp = studentOp.Where(student.MajorName.Eq(req.MajorName))
	}

	if req.ClassName != "" {
		studentOp = studentOp.Where(student.ClassName.Eq(req.ClassName))
	}

	if req.StudentName != "" {
		studentOp = studentOp.Where(student.StudentName.Eq(req.StudentName))
	}

	total, err := studentOp.Count()

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

func GetStudentPage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetStudentsReq

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

	student := mysql.Operator.Student
	studentOp := student.Where()

	if req.MajorName != "" {
		studentOp = studentOp.Where(student.MajorName.Eq(req.MajorName))
	}

	if req.ClassName != "" {
		studentOp = studentOp.Where(student.ClassName.Eq(req.ClassName))
	}

	if req.StudentName != "" {
		studentOp = studentOp.Where(student.StudentName.Eq(req.StudentName))
	}

	res, err := studentOp.Offset(index * req.PageSize).Limit(req.PageSize).Find()

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

func GetStudentById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	student := mysql.Operator.Student
	res, err := student.Where(student.StudentID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此学生"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = *res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateStudent(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.StudentInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	

	student := mysql.Operator.Student

	res, err := student.Where(student.StudentID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"student_name":    req.StudentName,
		"department_name": req.DepartmentName,
		"major_name":      req.MajorName,
		"class_name":      req.ClassName,
		"birthday":        req.Birthday,
		"email":           req.Email,
		"password":        req.Password,
		"sex":             req.Sex,
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "修改失败:没有这个学生"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateStudent(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.StudentInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	class := mysql.Operator.Class

	res,err := class.Where(class.ClassName.Eq(req.ClassName)).Find()

	if err != nil || len(res)==0 {
		resp.Code = -1
		resp.Message = "专业信息错误"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Student{
		StudentName:    req.StudentName,
		MajorName:      res[0].MajorName,
		DepartmentName: res[0].DepartmentName,
		ClassName: req.ClassName,
		Email:          req.Email,
		Password:       req.Password,
		Birthday:       req.Birthday,
		Sex:            uint(req.Sex),
	}
	err = mysql.Operator.Student.Create(&toAdd)

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

func DeleteStudent(ctx *gin.Context) {
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
		resp.Message = "无此学生"
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
		resp.Message = "删除失败: 没有这个学生"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetStudentList(ctx *gin.Context) {
	var nameList []model.NameListItem
	var resp model.JSONResp
	students, err := mysql.Operator.Student.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range students {
		nameList = append(nameList, model.NameListItem{
			ID:   int(v.StudentID),
			Name: v.StudentName,
		})
	}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = nameList
	ctx.JSON(http.StatusOK, resp)
	return
}
