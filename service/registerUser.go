package service

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/howeyc/gopass"
	"log"
	"os"
	"resume/entity"
	"resume/repository"
)

func RegisterUser(db *sql.DB, user entity.User) (entity.User, error) {

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
		exist, err := CheckEmail(email, repository.ConnectDatabase())
		if err != nil {
			log.Fatalln(err)
		}
		if exist {
			fmt.Println("This email has already been used , Try again")
			continue
		} else {
			user.Email = email
			break
		}
	}

	//get password
	for {
		fmt.Print("Enter password: ")
		pswByte, _ := gopass.GetPasswd()
		pswStr := ModifyPassword(pswByte)
		fmt.Print("Repeat password: ")
		confPswByte, _ := gopass.GetPasswd()
		confPswStr := ModifyPassword(confPswByte)
		pErr := CheckRePassword(pswStr, confPswStr)
		if pErr != nil {
			fmt.Println("Try again")
			continue
		}
		user.SetPassword(pswStr)
		break
	}

	return user
}
