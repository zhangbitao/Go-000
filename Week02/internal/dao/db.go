package dao

import (
	"database/sql"
	"fmt"
)

func NewDB() *sql.DB {
	auth := fmt.Sprintf("tcp:%s*%s/%s/%s", "127.0.0.1:3306", "learning", "root", "root@password")
	db, err := sql.Open("mysql", auth)
	if err != nil {
		panic(fmt.Sprintf("failed to connect db, %s", err))
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.Ping()

	return db
}
