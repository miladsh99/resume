package repository

import (
	"database/sql"
	"project1/dto"
	"project1/utills"
)

func ConnectDB(dsn string) (*sql.DB, *dto.ErrorHandle) {

	db, cErr := sql.Open("mysql", dsn)
	if cErr != nil {
		return nil, &dto.ErrorHandle{Type: utills.DB}
	}

	return db, nil
}
