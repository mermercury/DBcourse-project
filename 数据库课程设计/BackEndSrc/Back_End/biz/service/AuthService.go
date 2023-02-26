package service

import (
	"Back_End/biz/dal/mysql"
	"Back_End/model"
)

func ServeLogin(userId uint, password string, loginType model.UserType) (int, model.User) {

	student := mysql.Operator.Student
	admin := mysql.Operator.Admin
	teacher := mysql.Operator.Teacher
	var userInfo model.User
	var err error

	switch loginType {
	case model.TypeAdmin:
		var adminInfos []*model.Admin
		adminInfos, err = admin.Where(admin.AdminID.Eq(userId)).Find()
		if err != nil {
			return -3, nil
		}
		if len(adminInfos) <= 0 {
			return -1, nil
		}
		userInfo = *adminInfos[0]
	case model.TypeTeacher:
		var stuInfos []*model.Teacher
		stuInfos, err = teacher.Where(teacher.TeacherID.Eq(userId)).Find()
		if err != nil {
			return -3, nil
		}
		if len(stuInfos) <= 0 {
			return -1, nil
		}
		userInfo = *stuInfos[0]
	case model.TypeStudent:
		var stuInfos []*model.Student
		stuInfos, err = student.Where(student.StudentID.Eq(userId)).Find()
		if err != nil {
			return -3, nil
		}
		if len(stuInfos) <= 0 {
			return -1, nil
		}
		userInfo = *stuInfos[0]
	}

	if userInfo.GetPassword() == "" || userInfo.GetPassword() != password {
		return -2, nil
	}

	return 0, userInfo
}
