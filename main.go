package main

import (
	"log"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
var db *sql.DB

type Auditlog struct {
	id        int       `json:"id"`
	log_time  string    `json:"log_time"`
	admin     string    `json:"admin"`
	ip        string    `json:"ip"`
	action    string    `json:"action"`
	result    string    `json:"result"`
	detail    string    `json:"detail"`
}
*/


func main() {
	db, err := sql.Open("mysql", "root:Trend#1..@tcp(10.206.156.128:3306)/auditlog")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	router.Run(":8000")
}
