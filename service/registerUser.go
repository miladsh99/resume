package service

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"resume/entity"
	"resume/repository"
	"time"
)

func RegisterUser(db *sql.DB) {

	currentTime := time.Now()
	newuser := GetRegisterInfo()

	newUser := entity.User{
		Name:       newuser.Name,
		Email:      newuser.Email,
		Password:   newuser.Password,
		Created_at: currentTime,
		Updated_at: currentTime,
	}

	_, rErr := db.Exec("INSERT INTO users (ID, Name, Email, Password, Created_at, Updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		newUser.ID, newUser.Name, newUser.Email, newUser.Password, newUser.Created_at, newUser.Updated_at)
	if rErr != nil {
		fmt.Println(rErr)
	}
	fmt.Println("Register successfully")

}

func GetRegisterInfo() entity.User {
	var user entity.User
	reader := bufio.NewReader(os.Stdin)

	//get name
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	user.Name = ModifyValue(name)

	//get email
	for {
		fmt.Print("Enter email: ")
		email, _ := reader.ReadString('\n')
		email = ModifyValue(email)
		statusCode := CheckEmail(email, repository.ConnectDatabase())
		if statusCode == 1 {
			fmt.Println("This email has already been used , Try again")
			continue
		} else if statusCode == 2 {
			user.Email = email
			break
		}
	}

	//get password
	for {
		fmt.Print("Enter password: ")
		password1, _ := gopass.GetPasswd()
		password11 := ModifyPassword(password1)
		fmt.Print("Repeat password: ")
		password2, _ := gopass.GetPasswd()
		password22 := ModifyPassword(password2)
		pErr := CheckRePassword(password11, password22)
		if pErr != nil {
			fmt.Println("Try again")
			continue
		}
		user.Password = password11
		break
	}

	return user
}
