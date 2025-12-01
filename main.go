package main

import (
	"fmt"

	"github.com/bayuf/Golang_Authentication-System/handler"
	"github.com/bayuf/Golang_Authentication-System/repository"
	"github.com/bayuf/Golang_Authentication-System/service"
)

func main() {
	// init
	repo := repository.NewUserRepository("data/users.json")
	svc := service.NewAuthService(repo)
	handl := handler.NewAuthHandler(svc)

	for {
		fmt.Println("=== SIMPLE LOGIN SYSTEM ===")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Masukkan Pilihan: ")
		var num uint8
		fmt.Scan(&num)

		switch num {
		case 1:
			if err := handl.Login(); err != nil {
				fmt.Println(err)
			}
		case 2:
			handl.Register()
		case 3:
			return
		default:
			fmt.Println("masukkan angka 1-3")
		}
	}

}
