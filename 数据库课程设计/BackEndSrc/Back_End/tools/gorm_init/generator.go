package main

import (
	"Back_End/conf"
	"Back_End/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	mysqlConf, err := conf.ReadMySQL("./conf/conf.ini")

	if err != nil {
		panic(err)
	}

	connection := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Addr,
		mysqlConf.Port,
		mysqlConf.DatabaseName,
	)
	db, err := gorm.Open(mysql.Open(connection))

	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./biz/dal/mysql",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})

	g.UseDB(db)
	g.ApplyBasic(
		model.Course{},
		model.CourseSelect{},
		model.Department{},
		model.Student{},
		model.Teacher{},
		model.Admin{},
		model.Major{},
		model.Class{},
	)
	g.Execute()
}
