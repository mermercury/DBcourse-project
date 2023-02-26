package conf

import (
	"log"
	"github.com/go-ini/ini"
)

type MySQLConfig struct {
	Addr string `ini:"Addr"`
	Port int	`ini:"Port"`
	Username string `ini:"Username"`
	Password string `ini:"Password"`
	DatabaseName string `ini:"DatabaseName"`
}

type RedisConfig struct {
	Addr string `ini:"Addr"`
	Port int `ini:"Port"`
	Password string `ini:"Password"`
	DBIndex int `ini:"DBIndex"`
}

type AuthConfig struct {
	KeySize int `ini:"KeySize"`
	ExpireTime int64 `ini:"Expire"`
}

type AppConfig struct {
	MySQLConf MySQLConfig `ini:"MySQL"`
	RedisConf RedisConfig `ini:"Redis"`
	AuthConf AuthConfig `ini:"Auth"`
}

var GlobalConfig AppConfig

func init() {
	conf,err := ini.Load("./conf/conf.ini")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	//Load MySQL Config
	err = conf.MapTo(&GlobalConfig)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	GlobalConfig.AuthConf.ExpireTime *= 1000000000

}

func ReadMySQL(path string) (MySQLConfig,error) {

	var resConf MySQLConfig

	conf,err := ini.Load(path)
	if err != nil {
		return resConf,err
	}
	//Load MySQL Config
	mysqlSection := conf.Section("MySQL")
	resConf.Addr = mysqlSection.Key("Addr").String()
	resConf.Port = mysqlSection.Key("Port").MustInt(3306)
	resConf.Username = mysqlSection.Key("Username").String()
	resConf.Password = mysqlSection.Key("Password").String()
	resConf.DatabaseName = mysqlSection.Key("DatabaseName").String()

	return resConf,nil
}
