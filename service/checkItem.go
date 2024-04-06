package service

import (
	"database/sql"
	"errors"
	"fmt"
)

func CheckEmail(e string, db *sql.DB) (statusCode int) {

	rows, wErr := db.Query("SELECT email FROM users")
	if wErr != nil {
		fmt.Println(wErr)
	}
	defer rows.Close()

	for rows.Next() {
		var columnValue string
		if rErr := rows.Scan(&columnValue); rErr != nil {
			fmt.Println(rErr)
		}
		if columnValue == e {
			return 1
		}
	}
	return 2
}

func CheckPassword(email string, password1 string, db *sql.DB) (statusCode int) {
	var password2 string
	dErr := db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&password2)
	if dErr != nil {
		fmt.Println(dErr)
	}
	if password1 == password2 {
		return 1
	} else {
		return 2
	}
}

func CheckRePassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		fmt.Println("The passwords do not match")
		return errors.New("err")
	}
}
