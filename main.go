package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"resume/repository"
	"resume/service"
)

func main() {
	db := repository.ConnectDatabase()
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

OuterLoop:
	for {
		fmt.Print("What is your request? (login or register or exit) : ")
		input, _ := reader.ReadString('\n')
		input = service.ModifyValue(input)
		switch input {
		case "login":
			service.LoginUser()
		case "register":
			service.RegisterUser(db)
		case "exit":
			break OuterLoop
		default:
			fmt.Println("Wrong entry")
		}
	}
}
