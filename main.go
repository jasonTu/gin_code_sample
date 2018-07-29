package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*
var db *sql.DB
*/

type AuditLog struct {
	Id        int       `json:"id"` 
	LogTime  int       `json:"log_time" binding:"required"`
	Admin     string    `json:"admin" binding:"required"`
	Ip        string    `json:"ip" binding:"required"`
	Action    string    `json:"action" binding:"required"`
	Result    string    `json:"result" binding:"required"`
	Detail    string    `json:"detail"`
}

/*
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
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

	router.POST("/auditlog", func(c *gin.Context) {
		/*
        var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		*/
		var audit AuditLog
		if err := c.ShouldBindJSON(&audit); err != nil {
		    fmt.Println("auditlog input error: ", err)
			c.JSON(http.StatusOK, gin.H{"status": "input data invalid"})
		}
		fmt.Println("auditlog log_time: ", audit.LogTime)
		fmt.Println("auditlog admin: ", audit.Admin)
		rs, err := db.Exec(
			"insert into cloud(log_time, admin, ip, action, result, details) values(?, ?, ?, ?, ?, ?)",
			audit.LogTime, audit.Admin, audit.Ip, audit.Action, audit.Result, "")
		if err != nil {
			log.Fatalln(err)
		}
		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert auditlog id {}", id)
		c.JSON(http.StatusOK, gin.H{"msg": "create auditlog ok"})
	})

	router.Run(":8000")
}
