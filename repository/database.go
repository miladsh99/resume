package repository

import (
	"database/sql"
	"resume/dto"
	"resume/utills"
)

func ConnectDatabase() (*sql.DB, *dto.ErrorHandle) {

	db, cErr := sql.Open("mysql", "root:dalim123@tcp(127.0.0.1:3306)/resume-job?parseTime=true")
	if cErr != nil {
		return nil, &dto.ErrorHandle{Type: utills.DB}
	}

	return db, nil
}
