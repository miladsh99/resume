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
		user, gErr := FindUserByEmail(email, repository.ConnectDatabase())
		if gErr != nil {
			fmt.Println(gErr)
			continue
		}

		fmt.Print("Enter password: ")
		password, _ := gopass.GetPasswd()
		password1 := ModifyPassword(password)

		if user.GetPassword() != password1 {
			fmt.Println("email or password is not correct")
			continue
		}
		fmt.Println("login successfully")
		break
	}

}
