package main

import (
	"Back_End/biz/handler"
	"Back_End/biz/middleware"

	"github.com/gin-gonic/gin"
)

func register(engine *gin.Engine) {
	engine.Use(middleware.GetCors())
	v1 := engine.Group("/api/v1")
	v1.GET("/paint-egg/", handler.PaintEgg)

	upload := v1.Group("/upload", middleware.AuthAsAdmin())
	upload.POST("/studentTable", handler.HandleUploadStudent)
	upload.POST("/teacherTable", handler.HandleUploadTeacher)

	auth := v1.Group("/auth")
	auth.POST("/login", handler.HandleLogin)
	auth.POST("/logout", handler.HandleLogout)
	auth.GET("/status", handler.HandleGetStatus)

	student := v1.Group("/student", middleware.AuthAsStudent())
	student.POST("/course/select", handler.SelectCourse)
	student.GET("/course", handler.GetSelectedCourse)
	student.DELETE("/course/:id", handler.UnselectCourse)
	student.GET("/course/page/count", handler.GetCoursePageCount)
	student.GET("/course/page/:index", handler.GetCoursePage)
	student.GET("/exam", handler.GetExamList)
	student.GET("/info", handler.GetStudentInfo)
	student.PUT("/info", handler.UpdateStudentInfo)
	student.GET("/evaluateCourse", handler.GetNotEvaluatedCourse)
	student.PUT("/evaluateCourse", handler.SubmitEvaluation)
	student.GET("/timetable", handler.GetStudentCourseTable)
	student.GET("/static/score", handler.StaticScore)

	teacher := v1.Group("/teacher", middleware.AuthAsTeacher())

	teacher.GET("/course/list", handler.GetTeacherCourse)
	teacher.GET("/grade/:id", handler.GetStudentGrade)
	teacher.PUT("/grade", handler.UpdateStudentGrade)
	teacher.GET("/grade/page/count", handler.GetGradePageCount)
	teacher.GET("/grade/page/:index", handler.GetGradePage)
	teacher.GET("/timetable", handler.GetTeacherCourseTable)
	teacher.GET("/static/evaluation", handler.StaticEvaluate)

	admin := v1.Group("/admin", middleware.AuthAsAdmin())

	admin.GET("/major/page/count", handler.GetMajorPageCount)
	admin.GET("/major/page/:index", handler.GetMajorPage)
	admin.GET("/major/:id", handler.GetMajorById)
	admin.POST("/major", handler.CreateMajor)
	admin.PUT("/major", handler.UpdateMajor)
	admin.DELETE("/major/:id", handler.DeleteMajor)
	admin.GET("/major/names", handler.GetMajorList)

	admin.GET("/department/page/count", handler.GetDepartmentPageCount)
	admin.GET("/department/page/:index", handler.GetDepartmentPage)
	admin.GET("/department/:id", handler.GetDepartmentById)
	admin.POST("/department", handler.CreateDepartment)
	admin.PUT("/department", handler.UpdateDepartment)
	admin.DELETE("/department/:id", handler.DeleteDepartment)
	admin.GET("/department/names", handler.GetDepartmentList)

	admin.GET("/student/page/count", handler.GetStudentPageCount)
	admin.GET("/student/page/:index", handler.GetStudentPage)
	admin.GET("/student/:id", handler.GetStudentById)
	admin.POST("/student", handler.CreateStudent)
	admin.PUT("/student", handler.UpdateStudent)
	admin.DELETE("/student/:id", handler.DeleteStudent)
	admin.GET("/student/names", handler.GetStudentList)

	admin.GET("/student/course/page/count", handler.GetSelectPageCount)
	admin.GET("/student/course/page/:index", handler.GetSelectPage)
	admin.GET("/student/course/:id", handler.GetSelectById)
	admin.POST("/student/course", handler.CreateSelect)
	admin.PUT("/student/course", handler.UpdateSelect)
	admin.DELETE("/student/course/:id", handler.DeleteSelect)

	admin.GET("/admin/:id", handler.GetAdminById)
	admin.POST("/admin", handler.CreateAdmin)
	admin.DELETE("/admin/:id", handler.DeleteAdmin)
	admin.PUT("/admin", handler.UpdateAdmin)
	admin.GET("/admin", handler.GetAdminList)

	admin.GET("/class/page/count", handler.GetClassPageCount)
	admin.GET("/class/page/:index", handler.GetClassPage)
	admin.GET("/class/:id", handler.GetClassById)
	admin.POST("/class", handler.CreateClass)
	admin.PUT("/class", handler.UpdateClass)
	admin.DELETE("/class/:id", handler.DeleteClass)
	admin.GET("/class/name/list", handler.GetClassList)

	admin.GET("/course/page/count", handler.GetAdminCoursePageCount)
	admin.GET("/course/page/:index", handler.GetAdminCoursePage)
	admin.GET("/course/:id", handler.GetCourseById)
	admin.POST("/course", handler.CreateCourse)
	admin.PUT("/course", handler.UpdateCourse)
	admin.DELETE("/course/:id", handler.DeleteCourse)
	admin.GET("/course/name/list", handler.GetCourseList)
	admin.GET("/course/static/evaluation", handler.StaticCourseEvaluation)

	admin.GET("/teacher/page/count", handler.GetTeacherPageCount)
	admin.GET("/teacher/page/:index", handler.GetTeacherPage)
	admin.GET("/teacher/:id", handler.GetTeacherById)
	admin.POST("/teacher", handler.CreateTeacher)
	admin.PUT("/teacher", handler.UpdateTeacher)
	admin.DELETE("/teacher/:id", handler.DeleteTeacher)
	admin.GET("/teacher/name/list", handler.GetTeacherList)

	admin.GET("/teacher/names", func(context *gin.Context) {
		context.Request.URL.Path = "/api/v1/admin/department/names"
		engine.HandleContext(context)
	})

	admin.GET("/class/names", func(context *gin.Context) {
		context.Request.URL.Path = "/api/v1/admin/major/names"
		engine.HandleContext(context)
	})

	admin.GET("/course/names", func(context *gin.Context) {
		context.Request.URL.Path = "/api/v1/admin/teacher/name/list"
		engine.HandleContext(context)
	})

	// GET /api/v1/admin/course/names 是获取所有的教师的信息
	// GET /api/v1/admin/teacher/names 和 GET /api/v1/admin/department/names 都是获取系的信息
	// GET /api/v1/admin/class/names 和 GET /api/v1/admin/major/names 都是获取专业的信息
}
