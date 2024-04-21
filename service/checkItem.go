package service

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"resume/entity"
)

func CheckEmail(e string, db *sql.DB) (statusCode int) {
	email, eErr := CheckRegexEmail(e)
	if eErr != nil {
		return 3
	}
	var nEmail string
	cErr := db.QueryRow(`SELECT email from users where email=?`, email).Scan(&nEmail)
	if cErr != nil {
		err := sql.ErrNoRows
		if errors.As(cErr, &err) {
			return 0
		}
		return 1
	}
	return 2
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

func CheckRegexEmail(input string) (string, error) {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	if re.MatchString(input) {
		return input, nil
	} else {
		return "", errors.New("input is not a valid email")
	}
}
