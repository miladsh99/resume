package service

import (
	"database/sql"
	"errors"
	"net/http"
	"resume/entity"
	"resume/repository"
	"resume/utills"
)

func CheckEmail(e string, db *sql.DB, w http.ResponseWriter) utills.ErrorType {
	var nEmail string
	cErr := db.QueryRow(`SELECT email from users where email=?`, e).Scan(&nEmail)
	if cErr != nil {
		err := sql.ErrNoRows
		if errors.As(cErr, &err) {
			return utills.NotExist
		}
		return utills.DB
	}
	return utills.Exist
}

func InsertUserDataInDB(user entity.User) (entity.User, error) {

	db := repository.ConnectDatabase()
	defer db.Close()

	res, rErr := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.GetPassword())
	if rErr != nil {
		return entity.User{}, rErr
	}
	// insert into db ->
	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil
}
