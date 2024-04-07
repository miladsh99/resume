package service

import (
	"database/sql"
	"errors"
	"fmt"
	"resume/entity"
)

func CheckEmail(e string, db *sql.DB) (bool, error) {
	var email string
	row := db.QueryRow(`SELECT email from users where email=?`, e)
	sErr := row.Scan(&email)
	if sErr != nil {
		err := sql.ErrNoRows
		if errors.As(sErr, &err) {
			return false, nil
		}
		return false, errors.New("something went wrong")
	}

	return true, nil
}

func CheckPassword(email string, password1 string, db *sql.DB) (bool, error) {
	var password2 string
	sErr := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&password2)
	if sErr != nil {
		err := sql.ErrNoRows
		if errors.As(sErr, &err) {
			return false, nil
		}
		return false, errors.New("something went wrong")
	}

	return true, nil
}

func CheckRePassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		fmt.Println("The passwords do not match")
		return errors.New("err")
	}
}

func FindUserByEmail(email string, db *sql.DB) (entity.User, error) {
	row := db.QueryRow(`SELECT * from users where email=?`, email)
	var user entity.User
	var password string
	sErr := row.Scan(&user.ID, &user.Name, &user.Email, &password, &user.CreatedAt, &user.UpdatedAt)
	if sErr != nil {
		err := sql.ErrNoRows
		if errors.As(sErr, &err) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, errors.New("something went wrong")
	}

	user.SetPassword(password)

	return user, nil
}
