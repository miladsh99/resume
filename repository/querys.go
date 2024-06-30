package repository

import (
	"database/sql"
	"errors"
	"resume/dto"
	"resume/entity"
	"resume/utills"
)

func InsertUserDataInDB(user *entity.User) (*entity.User, *dto.ErrorHandle) {

	db, dErrType := ConnectDatabase()
	if dErrType != nil {
		return nil, dErrType
	}
	defer db.Close()

	res, rErr := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.GetPassword())
	if rErr != nil {
		return nil, &dto.ErrorHandle{Type: utills.FailedGetDataFromDB}
	}
	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil
}

func CheckUserByEmail(e string) (*entity.User, *dto.ErrorHandle) {

	var user entity.User
	var password string

	db, dErrType := ConnectDatabase()
	if dErrType != nil {
		return nil, dErrType
	}
	defer db.Close()

	row := db.QueryRow(`SELECT * from users where email=?`, e)
	cErr := row.Scan(&user.ID, &user.Name, &user.Email, &password, &user.CreatedAt, &user.UpdatedAt)
	user.SetPassword(password)
	if cErr != nil {
		err := sql.ErrNoRows
		if errors.As(cErr, &err) {
			return &user, &dto.ErrorHandle{Type: utills.NotExistEmail}
		}
		return &user, &dto.ErrorHandle{Type: utills.FailedGetDataFromDB}
	}
	return &user, &dto.ErrorHandle{Type: utills.ExistEmail}

}
