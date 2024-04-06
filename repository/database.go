package repository

import (
	"database/sql"
	"fmt"
)

func ConnectDatabase() *sql.DB {
	db, cErr := sql.Open("mysql", "root:dalim123@tcp(127.0.0.1:3306)/resume-job")
	if cErr != nil {
		fmt.Println(cErr)
	}
	return db
}
