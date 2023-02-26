package mysql

import (
	"Back_End/conf"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

var Operator *Query

func init() {

	connection := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GlobalConfig.MySQLConf.Username,
		conf.GlobalConfig.MySQLConf.Password,
		conf.GlobalConfig.MySQLConf.Addr,
		conf.GlobalConfig.MySQLConf.Port,
		conf.GlobalConfig.MySQLConf.DatabaseName,
	)
	var err error
	DBConn, err = gorm.Open(mysql.Open(connection))
	Operator = Use(DBConn)
	if err != nil {
		panic(err)
	}
}
