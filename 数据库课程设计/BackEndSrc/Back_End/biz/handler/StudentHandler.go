package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/conf"
	"Back_End/model"
	"Back_End/utils/redis"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSelectedCourse(ctx *gin.Context) {
	var resp model.JSONResp
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	teacher := mysql.Operator.Teacher
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	var rawData []model.StudentCourseRawData
	err = courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		LeftJoin(teacher, course.TeacherID.EqCol(teacher.TeacherID)).
		Where(courseSelect.StudentID.Eq(uint(studentID))).
		Select(
			courseSelect.CourseSelectID,
			course.CourseID,
			course.CourseName,
			teacher.TeacherID,
			teacher.TeacherName,
			course.Credit,
			courseSelect.CourseScore,
			course.ExamDate,
			course.ExamLoc,
		).
		Scan(&rawData)

	if err != nil {
		resp.Code = -1
		resp.Message = "获取课程信息失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	var respData []model.StudentCourseData

	for _, v := range rawData {
		respData = append(respData, model.StudentCourseData{
			CourseSelectID: v.CourseSelectID,
			CourseName:     v.CourseName,
			TeacherName:    v.TeacherName,
			Credit:         v.Credit,
			CourseScore:    v.CourseScore,
		})
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = respData
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetExamList(ctx *gin.Context) {
	var resp model.JSONResp
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	teacher := mysql.Operator.Teacher
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	var rawData []model.StudentCourseRawData
	err = courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		LeftJoin(teacher, course.TeacherID.EqCol(teacher.TeacherID)).
		Where(courseSelect.StudentID.Eq(uint(studentID))).
		Select(
			courseSelect.CourseSelectID,
			course.CourseID,
			course.CourseName,
			teacher.TeacherID,
			teacher.TeacherName,
			course.Credit,
			courseSelect.CourseScore,
			course.ExamDate,
			course.ExamLoc,
		).
		Scan(&rawData)

	if err != nil {
		resp.Code = -1
		resp.Message = "获取考试信息失败"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	var respData []model.ExamData

	for _, v := range rawData {
		respData = append(respData, model.ExamData{
			CourseSelectID: v.CourseSelectID,
			CourseName:     v.CourseName,
			TeacherName:    v.TeacherName,
			ExamDate:       v.ExamDate,
			ExamLoc:        v.ExamLoc,
		})
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = respData
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetStudentInfo(ctx *gin.Context) {
	var resp model.JSONResp
	student := mysql.Operator.Student
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	infos, err := student.Where(student.StudentID.Eq(uint(studentID))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(infos) != 1 {
		resp.Code = -1
		resp.Message = "没有这个学生"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = infos[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateStudentInfo(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.StudentInfoReq
	student := mysql.Operator.Student
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	err = ctx.ShouldBind(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := student.Where(student.StudentID.Eq(uint(studentID))).UpdateColumns(map[string]interface{}{
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

	if res.RowsAffected != 1 {
		resp.Code = -1
		resp.Message = "修改失败: 没有这个学生"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetNotEvaluatedCourse(ctx *gin.Context) {
	var resp model.JSONResp
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	teacher := mysql.Operator.Teacher
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	var rawData []model.StudentCourseRawData
	err = courseSelect.
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID)).
		LeftJoin(teacher, course.TeacherID.EqCol(teacher.TeacherID)).
		Where(courseSelect.StudentID.Eq(uint(studentID)), courseSelect.EvaluateScore.Lt(0)).
		Select(
			courseSelect.CourseSelectID,
			course.CourseID,
			course.CourseName,
			teacher.TeacherID,
			teacher.TeacherName,
			course.Credit,
			courseSelect.CourseScore,
			course.ExamDate,
			course.ExamLoc,
		).
		Scan(&rawData)

	if err != nil {
		resp.Code = -1
		resp.Message = "获取课程信息失败"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	var respData []model.NotEvaluateCourseDate

	for _, v := range rawData {
		respData = append(respData, model.NotEvaluateCourseDate{
			CourseID:    v.CourseID,
			CourseName:  v.CourseName,
			TeacherName: v.TeacherName,
			Credit:      v.Credit,
		})
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = respData
	ctx.JSON(http.StatusOK, resp)
}

func SubmitEvaluation(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.SubmitEvaluationReq
	courseSelect := mysql.Operator.CourseSelect
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	err = ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := courseSelect.
		Where(courseSelect.CourseID.Eq(uint(req.CourseID)), courseSelect.StudentID.Eq(uint(studentID))).
		UpdateColumns(map[string]interface{}{
			"evaluate_score": req.Star,
			"evaluation":     req.Evaluation,
		})

	if err != nil {
		resp.Code = -1
		resp.Message = "提交评价失败" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if res.RowsAffected != 1 {
		resp.Code = -1
		resp.Message = "没有满足条件的选课记录"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetCoursePageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetCoursesReq

	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
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

	course := mysql.Operator.Course

	courseSelect := mysql.Operator.CourseSelect
	var courseIDs []model.CourseIDs
	var ids []uint
	err = courseSelect.Where(courseSelect.StudentID.Eq(uint(studentID))).Select(courseSelect.CourseID).Scan(&courseIDs)
	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, v := range courseIDs {
		ids = append(ids, v.CourseID)
	}
	courseOp := course.Where(course.CourseID.NotIn(ids...))

	if req.CourseName != "" {
		courseOp = courseOp.Where(course.CourseName.Eq(req.CourseName))
	}

	if req.TeacherName != "" {
		teacher := mysql.Operator.Teacher
		courseOp = courseOp.
			LeftJoin(teacher, course.TeacherID.
				EqCol(teacher.TeacherID)).
			Where(teacher.TeacherName.
				Eq(req.TeacherName))
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

func GetCoursePage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetCoursesReq
	var respData []model.GetCourseData

	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info["id"])

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
		index--
	}

	course := mysql.Operator.Course
	teacher := mysql.Operator.Teacher
	courseSelect := mysql.Operator.CourseSelect
	var courseIDs []model.CourseIDs
	var ids []uint
	err = courseSelect.Where(courseSelect.StudentID.Eq(uint(studentID))).Select(courseSelect.CourseID).Scan(&courseIDs)
	if err != nil {
		resp.Code = -1
		resp.Message = "查询错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	for _, v := range courseIDs {
		ids = append(ids, v.CourseID)
	}
	courseOp := course.Where(course.CourseID.NotIn(ids...))

	if req.CourseName != "" {
		courseOp = courseOp.Where(course.CourseName.Eq(req.CourseName))
	}

	if req.TeacherName != "" {
		courseOp = courseOp.
			LeftJoin(teacher, course.TeacherID.
				EqCol(teacher.TeacherID)).
			Where(teacher.TeacherName.
				Eq(req.TeacherName))
	} else {
		courseOp = courseOp.
			LeftJoin(teacher, course.TeacherID.
				EqCol(teacher.TeacherID))
	}

	err = courseOp.Offset(req.PageSize * index).Limit(req.PageSize).Scan(&respData)

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

func SelectCourse(ctx *gin.Context) {
	var resp model.JSONResp
	var selectReq model.OperateCourseReq
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	ck, _ := ctx.Cookie("auth_session")
	info, _ := redis.GetHash(ck)
	studentID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "登录失效，请重试"
		ctx.SetCookie(conf.AuthCookieName, "", -1, "/", "*", true, true)
		redis.UnsetKey(ck)
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	err = ctx.ShouldBind(&selectReq)

	if err != nil {
		resp.Code = -2
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = mysql.Use(mysql.DBConn).Transaction(
		func(tx *mysql.Query) error {
			res, err := course.Where(course.CourseID.Eq(uint(selectReq.CourseID))).Update(course.Selected, course.Selected.Add(1))
			if err != nil {
				return err
			}
			if res.RowsAffected <= 0 {
				return errors.New("no such course")
			}

			var selectItem = model.CourseSelect{
				StudentID: uint(studentID),
				CourseID:  uint(selectReq.CourseID),
			}

			err = courseSelect.Create(&selectItem)
			return err
		},
	)

	if err != nil {
		resp.Code = -1
		resp.Message = "选课失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func UnselectCourse(ctx *gin.Context) {
	var resp model.JSONResp
	// var unselectReq model.OperateCourseReq
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	ck, _ := ctx.Cookie("auth_session")
	info, _ := redis.GetHash(ck)
	studentID, err := strconv.Atoi(info[conf.RedisIDField])

	if err != nil {
		resp.Code = -1
		resp.Message = "登录失效，请重试"
		ctx.SetCookie(conf.AuthCookieName, "", -1, "/", "*", true, true)
		redis.UnsetKey(ck)
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}

	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := courseSelect.
		Where(courseSelect.CourseSelectID.Eq(uint(id))).
		Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "退课失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "没有对应选课记录"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if res[0].StudentID != uint(studentID) {
		resp.Code = -1
		resp.Message = "不能退选别人的课"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = mysql.Use(mysql.DBConn).Transaction(
		func(tx *mysql.Query) error {

			res, err := courseSelect.
				Where(courseSelect.CourseSelectID.Eq(uint(id))).
				Delete()

			if err != nil {
				return err
			}

			if res.RowsAffected <= 0 {
				return errors.New("haven't select this course")
			}

			_, err = course.Where(course.CourseID.Eq(uint(uint(id)))).Update(course.Selected, course.Selected.Sub(1))

			return err
		},
	)

	if err != nil {
		resp.Code = -1
		resp.Message = "退课失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func GetStudentCourseTable(ctx *gin.Context) {
	var resp model.JSONResp
	var respData [7][10]model.CourseTableItem
	courseMap := make(map[int]model.CourseTableItem)
	course := mysql.Operator.Course
	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info[conf.RedisIDField])
	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}
	teacher := mysql.Operator.Teacher
	courseSelect := mysql.Operator.CourseSelect
	courses, err := course.LeftJoin(courseSelect, course.CourseID.EqCol(courseSelect.CourseID)).
		Where(courseSelect.StudentID.Eq(uint(studentID))).Find()

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
			tcInfo, _ := teacher.Where(teacher.TeacherID.Eq(v.TeacherID)).Find()
			var tcName string
			if len(tcInfo) > 0 {
				tcName = tcInfo[0].TeacherName
			}
			_, ok := courseMap[key]
			if ok {
				courseMap[key] = model.CourseTableItem{
					CourseName:  "冲突课程",
					TeacherName: "冲突",
					Location:    "冲突",
				}
			} else {
				courseMap[key] = model.CourseTableItem{
					CourseName:  v.CourseName,
					TeacherName: tcName,
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

func StaticScore(ctx *gin.Context) {
	var resp model.JSONResp
	data := make([]int, 3)

	ck, _ := ctx.Cookie("auth_session")

	info, _ := redis.GetHash(ck)

	studentID, err := strconv.Atoi(info[conf.RedisIDField])
	if err != nil {
		resp.Code = -1
		resp.Message = "认证失败:" + err.Error()
		ctx.JSON(http.StatusUnauthorized, resp)
		return
	}
	courseSelect := mysql.Operator.CourseSelect
	scores, err := courseSelect.Where(courseSelect.StudentID.Eq(uint(studentID))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for _, v := range scores {
		if v.CourseScore >= 0 && v.CourseScore < 60 {
			data[0]++
		} else if v.CourseScore >= 60 && v.CourseScore < 80 {
			data[1]++
		} else if v.CourseScore >= 80 && v.CourseScore <= 100 {
			data[2]++
		}
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
	return
}
