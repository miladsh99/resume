package service

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"resume/repository"
)

func LoginUser() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter email: ")
		email, _ := reader.ReadString('\n')
		email = ModifyValue(email)
		statusCode := CheckEmail(email, repository.ConnectDatabase())
		if statusCode == 2 {
			fmt.Println("Email is not available , Try again")
			continue
		}
		fmt.Print("Enter password: ")
		password, _ := gopass.GetPasswd()
		password1 := ModifyPassword(password)
		statusCode = CheckPassword(email, password1, repository.ConnectDatabase())
		if statusCode == 2 {
			fmt.Println("Password is wrong")
			continue
		} else if statusCode == 1 {
			fmt.Println("Login successfully")
			break
		}

	}

}
