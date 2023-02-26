package handler

import (
	"Back_End/biz/dal/mysql"
	"Back_End/model"
	"Back_End/utils/excel"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const FileFieldName = "file"

func HandleUploadStudent(ctx *gin.Context) {
	var resp model.JSONResp
	file, err := ctx.FormFile(FileFieldName)
	if err != nil {
		resp.Code = -1
		resp.Message = "接收上传文件失败"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fileUUID, _ := uuid.NewUUID()
	filePath := excel.ExcelBashPath + fileUUID.String() + ".xlsx"
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		resp.Code = -1
		resp.Message = "接收上传文件失败"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	students, err := excel.ReadStudentFromExcel(filePath)
	if err != nil {
		resp.Code = -1
		resp.Message = "添加学生信息失败"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	for i := range students {
		mysql.Operator.Student.Create(&students[i])
	}
	os.RemoveAll(filePath)
	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}

func HandleUploadTeacher(ctx *gin.Context) {
	var resp model.JSONResp
	file, err := ctx.FormFile(FileFieldName)
	if err != nil {
		resp.Code = -1
		resp.Message = "接收上传文件失败"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fileUUID, _ := uuid.NewUUID()
	filePath := excel.ExcelBashPath + fileUUID.String() + ".xlsx"
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		resp.Code = -1
		resp.Message = "接收上传文件失败"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	students, err := excel.ReadTeacherFromExcel(filePath)
	if err != nil {
		resp.Code = -1
		resp.Message = "添加老师信息失败"
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}
	department := mysql.Operator.Department
	for i := range students {
		mysql.Operator.Transaction(func(tx *mysql.Query) error {
			err := mysql.Operator.Teacher.Create(&students[i])

			if err != nil {
				return nil
			}

			res,err := department.
			Where(department.DepartmentName.Eq(students[i].DepartmentName)).
			Update(department.TeacherCount,department.TeacherCount.Add(1))

			if err != nil {
				return err
			}

			if res.RowsAffected <= 0 {
				return errors.New("no such department")
			}
			return nil
		})
		
	}
	os.RemoveAll(filePath)
	resp.Code = 0
	resp.Message = "success"
	ctx.JSON(http.StatusOK, resp)
	return
}
