package model

type UserType int

const (
	TypeStudent = 1
	TypeTeacher = 2
	TypeAdmin   = 3
)

type User interface {
	GetID() uint
	GetName() string
	GetType() UserType
	GetPassword() string
	GetPerm() int
}

type Admin struct {
	AdminID   uint   `gorm:"admin_id" json:"adminID"`
	AdminName string `gorm:"admin_name" json:"adminName"`
	Privilege int    `gorm:"privilege" json:"privilege"`
	Password  string `gorm:"password" json:"password"`
}

func (Admin) TableName() string {
	return "admin"
}

func (a Admin) GetName() string {
	return a.AdminName
}

func (a Admin) GetID() uint {
	return a.AdminID
}

func (a Admin) GetPassword() string {
	return a.Password
}

func (Admin) GetType() UserType {
	return TypeAdmin
}

func (a Admin) GetPerm() int {
	return a.Privilege
}

type Student struct {
	StudentID      uint   `gorm:"student_id" json:"id"`
	StudentName    string `gorm:"student_name" json:"name"`
	DepartmentName string `gorm:"department_name" json:"departmentName"`
	MajorName      string `gorm:"major_name" json:"majorName"`
	ClassName      string `gorm:"class_name" json:"className"`
	Email          string `gorm:"email" json:"email"`
	Birthday       string `gorm:"birthday" json:"birthday"`
	Sex            uint   `gorm:"sex" json:"sex"`
	Password       string `gorm:"password" json:"password"`
}

func (Student) TableName() string {
	return "student"
}

func (s Student) GetName() string {
	return s.StudentName
}

func (s Student) GetID() uint {
	return s.StudentID
}

func (s Student) GetPassword() string {
	return s.Password
}

func (Student) GetType() UserType {
	return TypeStudent
}

func (Student) GetPerm() int {
	return 0
}

type Teacher struct {
	TeacherID      uint   `gorm:"teacher_id" json:"id"`
	TeacherName    string `gorm:"teacher_name" json:"name"`
	DepartmentName string `gorm:"department_name" json:"departmentName"`
	Phone          string `gorm:"phone" json:"number"`
	Password       string `gorm:"password" json:"password"`
}

func (Teacher) TableName() string {
	return "teacher"
}

func (t Teacher) GetName() string {
	return t.TeacherName
}

func (t Teacher) GetID() uint {
	return t.TeacherID
}

func (t Teacher) GetPassword() string {
	return t.Password
}

func (Teacher) GetType() UserType {
	return TypeTeacher
}

func (Teacher) GetPerm() int {
	return 0
}

type Class struct {
	ClassID        uint   `gorm:"class_id" json:"id"`
	ClassName      string `gorm:"class_name" json:"name"`
	Grade          uint   `gorm:"grade" json:"grade"`
	DepartmentName string `gorm:"department_name" json:"departmentName"`
	MajorName      string `gorm:"major_name" json:"majorName"`
}

func (Class) TableName() string {
	return "class"
}

type Major struct {
	MajorID        uint   `gorm:"major_id" json:"id"`
	MajorName      string `gorm:"major_name" json:"name"`
	DepartmentName string `gorm:"department_name" json:"departmentName"`
}

func (Major) TableName() string {
	return "major"
}

type Department struct {
	DepartmentID   uint   `gorm:"department_id" json:"id"`
	DepartmentName string `gorm:"department_name" json:"departmentName"`
	MajorCount     int    `gorm:"major_count" json:"majorCount"`
	TeacherCount   int    `gorm:"teacher_count" json:"teacherCount"`
}

func (Department) TableName() string {
	return "department"
}

type Course struct {
	CourseID       uint   `gorm:"course_id" json:"id"`
	CourseName     string `gorm:"course_name" json:"courseName"`
	Grade          int    `gorm:"grade" json:"grade"`
	TeacherID      uint   `gorm:"teacher_id" json:"teacherID"`
	DepartmentName string `gorm:"department_name" json:"departmentName"`
	Credit         int    `gorm:"credit" json:"credit"`
	CourseTime     string `gorm:"course_time" json:"courseTime"`
	Location       string `gorm:"location" json:"location"`
	Selected       int    `gorm:"selected" json:"selected"`
	Size           uint   `gorm:"size" json:"size"`
	ExamDate       string `gorm:"exam_date" json:"examDate"`
	ExamLoc        string `gorm:"exam_loc" json:"examLocation"`
}

func (Course) TableName() string {
	return "course"
}

type CourseSelect struct {
	CourseSelectID uint   `gorm:"course_select_id"`
	StudentID      uint   `gorm:"student_id"`
	CourseID       uint   `gorm:"course_id"`
	CourseScore    int    `gorm:"course_score"`
	EvaluateScore  int    `gorm:"evaluate_score"`
	Evaluation     string `gorm:"evaluation"`
}

func (CourseSelect) TableName() string {
	return "course_select"
}
