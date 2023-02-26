package main

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	cfg, err := ini.Load("./conf/conf.ini")
	reportErr(err)

	addr := cfg.Section("MySQL").Key("Addr").String()
	port := cfg.Section("MySQL").Key("Port").String()
	uname := cfg.Section("MySQL").Key("Username").String()
	password := cfg.Section("MySQL").Key("Password").String()
	sqlName := "tools/db_init/db_init"
	dbName := cfg.Section("MySQL").Key("DBName").String()
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", uname, password, addr, port, dbName)

	db, err := sqlx.Open("mysql", url)
	reportErr(err)
	defer db.Close()

	sqlFile, err := os.Open("./" + sqlName + ".sql")
	reportErr(err)
	stat, err := os.Stat("./" + sqlName + ".sql")
	reportErr(err)
	buf := make([]byte, stat.Size())
	_, err = sqlFile.Read(buf)
	reportErr(err)
	prevPtr := 0
	for i := range buf {
		if buf[i] == ';' {
			cmd := string(buf[prevPtr : i+1])
			_, err = db.Exec(cmd)
			prevPtr = i + 1
			reportErr(err)
		}
	}

}

func reportErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
