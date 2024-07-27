package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"project1/dto"
	"project1/entity"
	"project1/utills"
)

func InsertUserDataInDB(db *sql.DB, user *entity.User) (*entity.User, *dto.ErrorHandle) {

	res, rErr := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.GetPassword())
	if rErr != nil {
		fmt.Println(rErr)
		return nil, &dto.ErrorHandle{Type: utills.FailedGetDataFromDB}
	}
	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil
}

func FindUserByEmail(db *sql.DB, e string) (*entity.User, *dto.ErrorHandle) {

	var user entity.User
	var password string

	row := db.QueryRow(`SELECT * from users where email=?`, e)
	cErr := row.Scan(&user.ID, &user.Name, &user.Email, &password, &user.CreatedAt, &user.UpdatedAt, &user.UserTypeId)
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
