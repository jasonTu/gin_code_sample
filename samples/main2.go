/*
package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type AuditLog struct {
	Id        int       `json:"id"` 
	LogTime   int       `json:"log_time" binding:"required"`
	Admin     string    `json:"admin" binding:"required"`
	Ip        string    `json:"ip" binding:"required"`
	Action    string    `json:"action" binding:"required"`
	Result    string    `json:"result" binding:"required"`
	Details    string    `json:"details"`
}

var SqlDB *sql.DB

func init() {
	log.Println("in db init")
	var err error
	SqlDB, err := sql.Open("mysql", "root:Trend#1..@tcp(10.206.156.128:3306)/auditlog") 
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db init success")
}

func (audit *AuditLog) AddAuditLog() (id int64, err error) {
	log.Println(audit.LogTime)
	log.Println(audit)
	rs, err := SqlDB.Exec(
		"insert into cloud(log_time, admin, ip, action, result, details) values(?, ?, ?, ?, ?, ?)",
		audit.LogTime, audit.Admin, audit.Ip, audit.Action, audit.Result, audit.Details)
	if err != nil {
	    log.Println("sql execute fail")
		return
	}
	id, err = rs.LastInsertId()
	return
}

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddAuditlogApi(c *gin.Context) {
	var audit AuditLog
	if err := c.ShouldBindJSON(&audit); err != nil {
		fmt.Println("auditlog input error: ", err)
		c.JSON(http.StatusOK, gin.H{"status": "input data invalid"})
		return
	}
	ra, err := audit.AddAuditLog()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("insert auditlog successful {}", ra)
	msg := fmt.Sprintln("insert auditlog successful %d", ra)
	c.JSON(http.StatusOK, gin.H{"msg": msg})
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/auditlog", AddAuditlogApi)

	return router
}


func main() {
	log.Println("in db init")
	var err error
	SqlDB, err := sql.Open("mysql", "root:Trend#1..@tcp(10.206.156.128:3306)/auditlog") 
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db init success")

	router := initRouter()
	router.Run(":8000")
}
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB
var TestInt int

type AuditLog struct {
	Id        int       `json:"id"` 
	LogTime  int       `json:"log_time" binding:"required"`
	Admin     string    `json:"admin" binding:"required"`
	Ip        string    `json:"ip" binding:"required"`
	Action    string    `json:"action" binding:"required"`
	Result    string    `json:"result" binding:"required"`
	Details    string    `json:"details"`
}

func (audit *AuditLog) AddAuditLog() (id int64, err error) {
	log.Println(audit.LogTime)
	log.Println(audit)
	log.Println(TestInt)
	log.Println(SqlDB)
	// 如果是函数参数的形式传入，则不会有问题
	rs, err := SqlDB.Exec(
		"insert into cloud(log_time, admin, ip, action, result, details) values(?, ?, ?, ?, ?, ?)",
		audit.LogTime, audit.Admin, audit.Ip, audit.Action, audit.Result, audit.Details)
	if err != nil {
	    log.Println("sql execute fail")
		return
	}
	id, err = rs.LastInsertId()
	return
}

func testFunc() {
	log.Println(TestInt)
	log.Println(SqlDB)
}

func main() {
	log.Println("in db init")
	var err error
	SqlDB, err = sql.Open("mysql", "root:Trend#1..@tcp(10.206.156.128:3306)/auditlog") 
	if err != nil {
		log.Fatal(err.Error())
	}
	TestInt = 5

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db init success")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		testFunc()
		c.String(http.StatusOK, "It works")
	})

	log.Println(SqlDB)
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
		ra, err := audit.AddAuditLog()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert auditlog id {}", ra)
		c.JSON(http.StatusOK, gin.H{"msg": "create auditlog ok"})
	})

	router.Run(":8000")
}
