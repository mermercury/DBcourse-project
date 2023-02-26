package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAdminCoursePageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetAdminCoursesReq
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

	course := mysql.Operator.Course
	courseOp := course.Where()

	if req.DepartmentName != "" {
		courseOp = courseOp.Where(course.DepartmentName.Eq(req.DepartmentName))
	}

	if req.CourseName != "" {
		courseOp = courseOp.Where(course.CourseName.Eq(req.CourseName))
	}

	total, err := courseOp.Count()

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

func GetAdminCoursePage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetAdminCoursesReq

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

	course := mysql.Operator.Course
	courseOp := course.Where()

	if req.DepartmentName != "" {
		courseOp = courseOp.Where(course.DepartmentName.Eq(req.DepartmentName))
	}

	if req.CourseName != "" {
		courseOp = courseOp.Where(course.CourseName.Eq(req.CourseName))
	}

	res, err := courseOp.Offset(index * req.PageSize).Limit(req.PageSize).Find()

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

func GetCourseById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	course := mysql.Operator.Course
	res, err := course.Where(course.CourseID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此课程"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = *res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateCourse(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.CourseInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	Course := mysql.Operator.Course

	res, err := Course.Where(Course.CourseID.Eq(uint(req.ID))).UpdateColumns(map[string]interface{}{
		"course_name":     req.CourseName,
		"department_name": req.DepartmentName,
		"teacher_id":      req.TeacherID,
		"course_time":     req.Time,
		"grade":           req.Grade,
		"location":        req.Location,
		"credit":          req.Credit,
		"exam_date":       req.ExamDate,
		"exam_loc":        req.ExamLoc,
		"size":            req.MaxSize,
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "修改失败:没有这个课程"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateCourse(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.CourseInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.Course{
		CourseName:     req.CourseName,
		DepartmentName: req.DepartmentName,
		TeacherID:      uint(req.TeacherID),
		Grade:          req.Grade,
		CourseTime:     req.Time,
		Location:       req.Location,
		Credit:         req.Credit,
		Size:           uint(req.MaxSize),
		ExamDate:       req.ExamDate,
		ExamLoc:        req.ExamLoc,
	}
	err = mysql.Operator.Course.Create(&toAdd)

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

func DeleteCourse(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	Course := mysql.Operator.Course
	CourseCheck, err := Course.Where(Course.CourseID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(CourseCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此课程"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := Course.Where(Course.CourseID.Eq(uint(id))).Delete()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected <= 0 {
		resp.Code = -1
		resp.Message = "删除失败: 没有这个课程"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetCourseList(ctx *gin.Context) {
	var nameList []model.NameListItem
	var resp model.JSONResp
	courses, err := mysql.Operator.Course.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range courses {
		nameList = append(nameList, model.NameListItem{
			ID:   int(v.CourseID),
			Name: v.CourseName,
		})
	}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = nameList
	ctx.JSON(http.StatusOK, resp)
	return
}

func StaticCourseEvaluation(ctx *gin.Context) {
	var resp model.JSONResp
	data := make([]int, 6)
	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	scores, err := courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range scores {
		if v.EvaluateScore >= 0 && v.EvaluateScore <= 5 {
			data[v.EvaluateScore]++
		}
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
	return
}
