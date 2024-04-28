package service

import (
	"database/sql"
	"errors"
	"net/http"
	"resume/entity"
	"resume/utills"
)

func FindUserByEmail(req entity.User, db *sql.DB, w http.ResponseWriter) entity.User {
	row := db.QueryRow(`SELECT * from users where email=?`, req.Email)
	var user entity.User
	var password string
	sErr := row.Scan(&user.ID, &user.Name, &user.Email, &password, &user.CreatedAt, &user.UpdatedAt)
	if sErr != nil {
		err := sql.ErrNoRows
		if errors.As(sErr, &err) {
			utills.ErrorManagement(w, utills.NotExist)
			return entity.User{}
		}
		utills.ErrorManagement(w, utills.DB)
		return entity.User{}
	}
	user.SetPassword(password)
	return user
}
