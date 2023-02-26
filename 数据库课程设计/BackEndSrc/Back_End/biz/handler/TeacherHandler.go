package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"Back_End/utils/redis"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetGradePageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.TeacherGradeReq
	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	student := mysql.Operator.Student
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	err = ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if req.PageSize <= 0 {
		resp.Code = -1
		resp.Message = "pageSize不能为0"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	operator := courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		LeftJoin(student, courseSelect.StudentID.EqCol(student.StudentID)).
		Where(course.TeacherID.Eq(uint(teacherID)))
	if req.CourseName != "" {
		operator = operator.Where(course.CourseName.Eq(req.CourseName))
	}

	if req.StudentName != "" {
		operator = operator.Where(student.StudentName.Eq(req.StudentName))
	}

	total, err := operator.Count()

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
		Count: int(ans),
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetGradePage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.TeacherGradeReq
	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	student := mysql.Operator.Student
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

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
		resp.Message = "参数解析错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if index > 0 {
		index--
	}

	operator := courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		LeftJoin(student, courseSelect.StudentID.EqCol(student.StudentID)).
		Where(course.TeacherID.Eq(uint(teacherID)))
	if req.CourseName != "" {
		operator = operator.Where(course.CourseName.Eq(req.CourseName))
	}

	if req.StudentName != "" {
		operator = operator.Where(student.StudentName.Eq(req.StudentName))
	}
	var respData []model.TeacherCourseSelectItem

	err = operator.Offset(req.PageSize*index).Limit(req.PageSize).
		Select(
			courseSelect.CourseSelectID,
			course.CourseName,
			student.StudentID,
			student.StudentName,
			courseSelect.CourseScore,
		).
		Scan(&respData)

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = respData
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetTeacherCourseTable(ctx *gin.Context) {
	var resp model.JSONResp
	var respData [7][10]model.CourseTableItem
	courseMap := make(map[int]model.CourseTableItem)
	course := mysql.Operator.Course
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])
	teacherName, ok := info[conf.RedisNameField]
	if err != nil || !ok {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	courses, err := course.Where(course.TeacherID.Eq(uint(teacherID))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range courses {
		var weekday, start, length int
		_, err := fmt.Sscanf(v.CourseTime, "%d-%d-%d", &weekday, &start, &length)
		if err != nil || weekday > 6 || weekday < 0 || start < 0 || start+length > 9 {
			resp.Code = -1
			resp.Message = "数据格式错误"
			ctx.JSON(http.StatusInternalServerError, resp)
			return
		}

		for i := start; i < start+length; i++ {
			key := weekday*10 + i
			_, ok := courseMap[key]
			if ok {
				courseMap[key] = model.CourseTableItem{
					CourseName:  "冲突课程",
					TeacherName: teacherName,
					Location:    "冲突",
				}
			} else {
				courseMap[key] = model.CourseTableItem{
					CourseName:  v.CourseName,
					TeacherName: teacherName,
					Location:    v.Location,
				}
			}
		}
	}

	for k, v := range courseMap {
		weekday := k / 10
		start := k % 10
		respData[weekday][start] = v
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = respData
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetTeacherCourse(ctx *gin.Context) {
	var resp model.JSONResp

	course := mysql.Operator.Course
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	courses, err := course.Where(course.TeacherID.Eq(uint(teacherID))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = courses
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetStudentGrade(ctx *gin.Context) {
	var resp model.JSONResp
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	studentIDStr := ctx.Param(conf.ParamID)
	studentID, err := strconv.Atoi(studentIDStr)
	if err != nil {
		resp.Code = -1
		resp.Message = "参数非法:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	operator := courseSelect.Where(courseSelect.CourseSelectID.Eq(uint(studentID))).
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).Where(course.TeacherID.Eq(uint(teacherID)))

	res, err := operator.Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无记录"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = model.ScoreItem{
		CourseSelectID: int(res[0].CourseSelectID),
		Score:          res[0].CourseScore,
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateStudentGrade(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.ScoreItem
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	err = ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数非法:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course

	courseCheck, err := courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		Where(course.TeacherID.Eq(uint(teacherID)), courseSelect.CourseSelectID.Eq(uint(req.CourseSelectID))).
		Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "修改错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(courseCheck) < 1 {
		resp.Code = -1
		resp.Message = "没有修改的权限"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := courseSelect.Where(courseSelect.CourseSelectID.Eq(uint(req.CourseSelectID))).Update(courseSelect.CourseScore, req.Score)
	if err != nil {
		resp.Code = -1
		resp.Message = "修改错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected != 1 {
		resp.Code = -1
		resp.Message = "没有对应的选课记录,或修改前后值一致"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return

}

func StaticEvaluate(ctx *gin.Context) {
	var resp model.JSONResp
	data := make([]int, 6)

	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	teacherID, err := strconv.Atoi(info[conf.RedisIDField])
	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}
	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	scores, err := courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		Where(course.TeacherID.Eq(uint(teacherID))).Find()

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
