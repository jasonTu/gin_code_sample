package apis

import (
	"net/http"
	"log"
	"fmt"
	//"strconv"
	"github.com/gin-gonic/gin"
	. "github.com/jasonTu/gin_code_sample/models"
)

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
