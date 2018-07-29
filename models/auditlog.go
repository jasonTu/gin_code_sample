package models

import (
	"log"
	db "github.com/jasonTu/gin_code_sample/database"
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

func (audit *AuditLog) AddAuditLog() (id int64, err error) {
	log.Println(audit.LogTime)
	log.Println(audit)
	rs, err := db.SqlDB.Exec(
		"insert into cloud(log_time, admin, ip, action, result, details) values(?, ?, ?, ?, ?, ?)",
		audit.LogTime, audit.Admin, audit.Ip, audit.Action, audit.Result, audit.Details)
	if err != nil {
	    log.Println("sql execute fail")
		return
	}
	id, err = rs.LastInsertId()
	return
}
