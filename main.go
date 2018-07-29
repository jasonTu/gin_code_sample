package main

import (
	"os"
	"io"
	"github.com/gin-gonic/gin"
	db "github.com/jasonTu/gin_code_sample/database"
)

func main() {
	// We don't need console color while writing to log file.
    gin.DisableConsoleColor()

	// Specific the log file.
	f, _ := os.Create("auditlog.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// Also write log to console.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
