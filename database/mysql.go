package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql",
		"adm_user1:12345678@tcp(127.0.0.1:3306)/mydb?parseTime=true&loc=Local&charset=utf8")
	if err != nil {
		log.Println(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Println(err.Error())
	}
}
