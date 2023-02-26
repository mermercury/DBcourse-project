package test

import (
	"Back_End/utils/excel"
	"fmt"
	"testing"
)

func TestParseStudentTable(t *testing.T) {
	datas, err := excel.ReadStudentFromExcel("./testStudent.xlsx")
	if err != nil {
		t.Error(err)
	}

	for _, v := range datas {
		fmt.Println(v)
	}

	tc, err := excel.ReadTeacherFromExcel("./testTeacher.xlsx")
	if err != nil {
		t.Error(err)
	}

	for _, v := range tc {
		fmt.Println(v)
	}
}
