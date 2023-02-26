package excel

import (
	"Back_End/model"
	"errors"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
)

const ExcelBashPath = "./temp/"

const (
	StudentID  = "学号"
	TeacherID  = "工号"
	Name       = "姓名"
	Class      = "班级"
	Password   = "密码"
	Email      = "邮箱"
	Birthday   = "生日"
	Sex        = "性别"
	Department = "学院"
	Major      = "专业"
	Phone      = "电话"
)

func init() {
	os.Mkdir(ExcelBashPath, 0777)
}

func getFirstTableRows(filePath string) (*excelize.Rows, error) {
	f, err := excelize.OpenFile(filePath)

	if err != nil {
		return nil, err
	}

	sheets := f.GetSheetList()

	if len(sheets) != 1 {
		return nil, errors.New("incorrect table format")
	}

	rows, err := f.Rows(sheets[0])

	return rows, err
}

func ReadTeacherFromExcel(filePath string) ([]model.Teacher, error) {
	rows, err := getFirstTableRows(filePath)
	var ans []model.Teacher
	var headerMap []string
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("empty table")
	}

	headers, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for _, v := range headers {
		headerMap = append(headerMap, v)
	}

	for rows.Next() {
		cols, err := rows.Columns()

		if err != nil {
			return nil, err
		}
		var newItem model.Teacher
		correct := true
		for i, v := range cols {
			switch headerMap[i] {
			case TeacherID:
				id, err := strconv.Atoi(v)
				if err != nil {
					correct = false
					break
				}
				newItem.TeacherID = uint(id)
			case Name:
				newItem.TeacherName = v
			case Department:
				newItem.DepartmentName = v
			case Password:
				newItem.Password = v
			case Phone:
				newItem.Phone = v
			}
		}
		if correct {
			ans = append(ans, newItem)
		}
	}
	return ans, nil
}

func ReadStudentFromExcel(filePath string) ([]model.Student, error) {
	rows, err := getFirstTableRows(filePath)
	var ans []model.Student
	var headerMap []string
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New("empty table")
	}

	headers, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for _, v := range headers {
		headerMap = append(headerMap, v)
	}

	for rows.Next() {
		cols, err := rows.Columns()

		if err != nil {
			return nil, err
		}
		var newItem model.Student
		correct := true
		for i, v := range cols {
			switch headerMap[i] {
			case StudentID:
				id, err := strconv.Atoi(v)
				if err != nil {
					correct = false
					break
				}
				newItem.StudentID = uint(id)
			case Name:
				newItem.StudentName = v
			case Class:
				newItem.ClassName = v
			case Password:
				newItem.Password = v
			case Department:
				newItem.DepartmentName = v
			case Major:
				newItem.MajorName = v
			case Email:
				newItem.Email = v
			case Birthday:
				newItem.Birthday = v
			case Sex:
				if strings.Contains(v, "男") {
					newItem.Sex = 1
				} else {
					newItem.Sex = 0
				}
			}
		}
		if correct {
			ans = append(ans, newItem)
		}
	}
	return ans, nil
}
