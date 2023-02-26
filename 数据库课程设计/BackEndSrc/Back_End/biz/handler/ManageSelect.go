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

func GetSelectPageCount(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetSelectsReq
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

	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	student := mysql.Operator.Student
	courseSelectOp := courseSelect.
		LeftJoin(student, courseSelect.StudentID.EqCol(student.StudentID)).
		LeftJoin(course, course.CourseID.EqCol(course.CourseID))

	if req.ClassName != "" {
		courseSelectOp = courseSelectOp.Where(student.ClassName.Eq(req.ClassName))
	}

	if req.StudentName != "" {
		courseSelectOp = courseSelectOp.Where(student.StudentName.Eq(req.StudentName))
	}

	if req.CourseName != "" {
		courseSelectOp = courseSelectOp.Where(course.CourseName.Eq(req.CourseName))
	}

	total, err := courseSelectOp.Count()

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

func GetSelectPage(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.GetSelectsReq

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

	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	student := mysql.Operator.Student
	courseSelectOp := courseSelect.
		LeftJoin(student, courseSelect.StudentID.EqCol(student.StudentID)).
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID))

	if req.ClassName != "" {
		courseSelectOp = courseSelectOp.Where(student.ClassName.Eq(req.ClassName))
	}

	if req.StudentName != "" {
		courseSelectOp = courseSelectOp.Where(student.StudentName.Eq(req.StudentName))
	}

	if req.CourseName != "" {
		courseSelectOp = courseSelectOp.Where(course.CourseName.Eq(req.CourseName))
	}

	var res []model.CourseSelectData
	err = courseSelectOp.Offset(index*req.PageSize).Limit(req.PageSize).
		Select(
			courseSelect.CourseSelectID,
			student.StudentName,
			student.ClassName,
			course.CourseName,
			courseSelect.CourseScore,
		).Scan(&res)

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

func GetSelectById(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	courseSelect := mysql.Operator.CourseSelect
	course := mysql.Operator.Course
	student := mysql.Operator.Student
	courseSelectOp := courseSelect.
		LeftJoin(student, courseSelect.StudentID.EqCol(student.StudentID)).
		LeftJoin(course, courseSelect.CourseID.EqCol(course.CourseID))

	var res []model.CourseSelectItem

	err = courseSelectOp.Where(courseSelect.CourseSelectID.Eq(uint(id))).
		Select(
			courseSelect.CourseSelectID,
			student.StudentName,
			student.StudentID,
			course.CourseID,
			course.CourseName,
			student.ClassName,
			courseSelect.CourseScore,
		).
		Scan(&res)

	if err != nil {
		resp.Code = -1
		resp.Message = "查询失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(res) != 1 {
		resp.Code = -1
		resp.Message = "无此选课记录"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	resp.Data = res[0]
	ctx.JSON(http.StatusOK, resp)
	return
}

func UpdateSelect(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.CourseSelectInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	courseSelectCheck, err := courseSelect.Where(courseSelect.CourseID.Eq(uint(req.CourseID)), courseSelect.StudentID.Eq(uint(req.StudentID))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "参数非法:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if len(courseSelectCheck) != 1 {
		resp.Code = -1
		resp.Message = "没有相应选课记录"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toUpdate := courseSelectCheck[0]

	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {

		res, err := course.Where(course.CourseID.Eq(toUpdate.CourseID)).Update(course.Selected, course.Selected.Sub(1))

		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such course")
		}

		res, err = courseSelect.
			Where(courseSelect.CourseID.Eq(uint(req.CourseID)), courseSelect.StudentID.Eq(uint(req.StudentID))).
			UpdateColumns(map[string]interface{}{
				"course_score": req.CourseScore,
			})

		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such course")
		}

		res, err = course.Where(course.CourseID.Eq(uint(req.CourseID))).Update(course.Selected, course.Selected.Add(1))

		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such course to choose")
		}
		return nil
	})

	if err != nil {
		resp.Code = -1
		resp.Message = "修改失败:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Code = 0
	resp.Message = "success"
	return
}

func CreateSelect(ctx *gin.Context) {
	var resp model.JSONResp
	var req model.CourseSelectInfoReq

	err := ctx.ShouldBind(&req)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数解析失败:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toAdd := model.CourseSelect{
		StudentID:     uint(req.StudentID),
		CourseID:      uint(req.CourseID),
		CourseScore:   req.CourseScore,
		EvaluateScore: -1,
	}
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect

	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {
		res, err := course.Where(course.CourseID.Eq(toAdd.CourseID)).Update(course.Selected, course.Selected.Add(1))

		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such course")
		}

		err = courseSelect.Create(&toAdd)

		if err != nil {
			return err
		}

		return nil

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

func DeleteSelect(ctx *gin.Context) {
	var resp model.JSONResp
	idStr := ctx.Param(conf.ParamID)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		resp.Code = -1
		resp.Message = "参数错误:" + err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	course := mysql.Operator.Course
	courseSelect := mysql.Operator.CourseSelect
	courseSelectCheck, err := courseSelect.Where(courseSelect.CourseSelectID.Eq(uint(id))).Find()

	if err != nil {
		resp.Code = -1
		resp.Message = "删除错误:" + err.Error()
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	if len(courseSelectCheck) != 1 {
		resp.Code = -1
		resp.Message = "无此选课记录"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	toDelete := courseSelectCheck[0]

	err = mysql.Operator.Transaction(func(tx *mysql.Query) error {
		res, err := course.Where(course.CourseID.Eq(toDelete.CourseID)).Update(course.Selected, course.Selected.Sub(1))

		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such course")
		}

		res, err = courseSelect.Where(courseSelect.CourseSelectID.Eq(uint(id))).Delete()

		if err != nil {
			return err
		}

		if res.RowsAffected <= 0 {
			return errors.New("no such course")
		}

		return nil

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
