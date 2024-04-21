package service

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"resume/entity"
	"resume/repository"
	"strings"
	"time"
)

func ModifyValue(v string) string {
	newValue := strings.TrimSpace(strings.ToLower(v))
	return newValue
}

func ModifyPassword(p []byte) string {
	newValue := strings.TrimSpace(string(p))
	return newValue
}

func ReadRequest(r *http.Request) entity.User {
	data, _ := io.ReadAll(r.Body)
	req := RegisterRequest{}
	json.Unmarshal(data, &req)

	user := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user.SetPassword(string(pass))

	return user
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
