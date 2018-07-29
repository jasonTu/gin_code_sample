package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	log.Println("in db init")
	var err error
	SqlDB, err = sql.Open("mysql", "root:Trend#1..@tcp(10.206.156.128:3306)/auditlog") 
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db init success")
}
