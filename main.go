package main

import (
	db "github.com/jasonTu/gin_code_sample/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8000")
}
