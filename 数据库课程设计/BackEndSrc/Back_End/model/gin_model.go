package model

type LoginInfo struct {
	UserID   int    `json:"userid"`
	Password string `json:"password"`
	Type     int    `json:"usertype"`
}
type JSONResp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type LoginStatusResp struct {
	LoggedIn   bool   `json:"loggedIn"`
	UserId     string `json:"userid"`
	Username   string `json:"username"`
	Usertype   int    `json:"usertype"`
	Permission int    `json:"permission"`
}

type GetCoursesReq struct {
	CourseName  string `json:"courseName" form:"courseName"`
	TeacherName string `json:"teacherName" form:"teacherName"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}

type GetAdminCoursesReq struct {
	DepartmentName string `json:"departmentName" form:"departmentName"`
	MajorName      string `json:"majorName" form:"majorName"`
	CourseName     string `json:"courseName" form:"courseName"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
}

type GetMajorsReq struct {
	DepartmentName string `json:"departmentName" form:"departmentName"`
	MajorName      string `json:"name" form:"name"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
}

type GetDepartmentsReq struct {
	DepartmentName string `json:"name" form:"name"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
}

type GetClassReq struct {
	ClassName      string `json:"name" form:"name"`
	MajorName      string `json:"majorName" form:"majorName"`
	DepartmentName string `json:"departmentName" form:"departmentName"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
}

type GetStudentsReq struct {
	StudentName string `json:"name" form:"name"`
	MajorName   string `json:"majorName" form:"majorName"`
	ClassName   string `json:"className" form:"className"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}

type GetTeachersReq struct {
	TeacherName    string `json:"name" form:"name"`
	DepartmentName string `json:"departmentName" form:"departmentName"`
	PageSize       int    `json:"pageSize" form:"pageSize"`
}

type GetSelectsReq struct {
	StudentName string `json:"studentName" form:"studentName"`
	ClassName   string `json:"className" form:"className"`
	CourseName  string `json:"courseName" form:"courseName"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}

type SubmitEvaluationReq struct {
	CourseID   int    `json:"courseId"`
	Star       int    `json:"star"`
	Evaluation string `json:"text"`
}

type OperateCourseReq struct {
	CourseID int `json:"courseId"`
}

type CountResp struct {
	Count int `json:"count"`
}

type CourseSelectData struct {
	CourseSelectID int     `json:"id"`
	ClassName      string  `json:"className"`
	StudentName    string  `json:"studentName"`
	CourseName     string  `json:"courseName"`
	MajorName      string  `json:"majorName"`
	CourseScore    float64 `json:"score"`
}

type CourseSelectItem struct {
	CourseSelectID int     `json:"id"`
	StudentName    string  `json:"studentName"`
	StudentID      int     `json:"studentId"`
	CourseID       int     `json:"courseId"`
	CourseName     string  `json:"courseName"`
	MajorName      string  `json:"majorName"`
	ClassName      string  `json:"className"`
	CourseScore    float64 `json:"score"`
}

type GetCourseData struct {
	CourseID       int    `json:"courseId"`
	CourseName     string `json:"courseName"`
	TeacherName    string `json:"teacherName"`
	Credit         int    `json:"credit"`
	Time           string `json:"time"`
	Selected       int    `json:"selectedCount"`
	MaxSize        int    `json:"maxSize"`
	DepartmentName string `json:"departmentName"`
	Grade          int    `json:"grade"`
}

type NameListItem struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name"`
}

type MajorInfoReq struct {
	ID             int    `json:"id"`
	MajorName      string `json:"name" form:"name"`
	DepartmentName string `json:"departmentName"`
}

type DepartmentInfoReq struct {
	ID             int    `json:"id"`
	DepartmentName string `json:"name" form:"name"`
	MajorCount     int    `json:"majorCount"`
	TeacherCount   int    `json:"teacherCount"`
}

type TeacherInfoReq struct {
	ID             int    `json:"id" form:"id"`
	TeacherName    string `json:"name" form:"name"`
	DepartmentName string `json:"departmentName" form:"departmentName"`
	Phone          string `json:"number" form:"number"`
	Password       string `json:"password" form:"password"`
}

type StudentInfoReq struct {
	ID             int    `json:"id"`
	StudentName    string `json:"name" form:"name"`
	ClassName      string `json:"className"`
	MajorName      string `json:"majorName"`
	DepartmentName string `json:"departmentName"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Birthday       string `json:"birthday"`
	Sex            int    `json:"sex"`
}

type CourseSelectInfoReq struct {
	ID          int `json:"id"`
	StudentID   int `json:"studentId"`
	CourseID    int `json:"courseId"`
	CourseScore int `json:"score"`
}

type AdminInfoReq struct {
	ID        int    `json:"adminID" form:"adminID"`
	AdminName string `json:"adminName" form:"adminName"`
	Privilege int    `json:"privilege" form:"privilege"`
	Password  string `json:"password" form:"password"`
}

type ClassInfoReq struct {
	ID             int    `json:"id"`
	ClassName      string `json:"name" form:"name"`
	Grade          int    `json:"grade"`
	DepartmentName string `json:"departmentName"`
	MajorName      string `json:"majorName"`
}

type CourseInfoReq struct {
	ID             int    `json:"id"`
	TeacherID      int    `json:"teacherId"`
	CourseName     string `json:"courseName"`
	DepartmentName string `json:"departmentName"`
	Grade          int    `json:"grade"`
	Time           string `json:"time"`
	Location       string `json:"location"`
	Credit         int    `json:"credit"`
	MaxSize        int    `json:"size"`
	ExamDate       string `json:"examDate"`
	ExamLoc        string `json:"examLocation"`
}

type StudentCourseRawData struct {
	CourseSelectID int
	CourseID       int
	CourseName     string
	TeacherID      int
	TeacherName    string
	Credit         int
	CourseScore    int
	ExamDate       string
	ExamLoc        string
}

type StudentCourseData struct {
	CourseSelectID int    `json:"studentCourseId"`
	CourseName     string `json:"courseName"`
	TeacherName    string `json:"teacherName"`
	Credit         int    `json:"credit"`
	CourseScore    int    `json:"score"`
}

type ExamData struct {
	CourseSelectID int    `json:"studentCourseId"`
	CourseName     string `json:"courseName"`
	TeacherName    string `json:"teacherName"`
	ExamDate       string `json:"examDate"`
	ExamLoc        string `json:"examLocation"`
}

type NotEvaluateCourseDate struct {
	CourseID    int    `json:"courseId"`
	CourseName  string `json:"courseName"`
	TeacherName string `json:"teacherName"`
	Credit      int    `json:"credit"`
}

type TeacherGradeReq struct {
	CourseName  string `json:"courseName" form:"courseName"`
	StudentName string `json:"studentName" form:"studentName"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}

type ScoreItem struct {
	CourseSelectID int `json:"studentCourseId"`
	Score          int `json:"score"`
}

type TeacherCourseSelectItem struct {
	CourseSelectID int    `json:"studentCourseId"`
	CourseName     string `json:"courseName"`
	StudentID      int    `json:"studentId"`
	StudentName    string `json:"studentName"`
	CourseScore    int    `json:"score"`
}

type CourseTableItem struct {
	CourseName  string `json:"courseName"`
	TeacherName string `json:"teacherName"`
	Location    string `json:"location"`
}

type CourseIDs struct {
	CourseID uint 
}
