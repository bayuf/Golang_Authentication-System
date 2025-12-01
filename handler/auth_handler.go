package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bayuf/Golang_Authentication-System/dto"
	"github.com/bayuf/Golang_Authentication-System/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: svc}
}

func (h *AuthHandler) Login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- LOGIN ---")
	reader.ReadString('\n')

	fmt.Print("Email\t\t: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Password\t: ")
	password, _ := reader.ReadString('\n')
	user, err := h.Service.UserLogin(dto.LoginDTO{
		Email:    strings.TrimSpace(email),
		Password: strings.TrimSpace(password),
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Selamat datang, ", user.Name)
	}

}

func (h *AuthHandler) Register() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- REGISTER ---")
	reader.ReadString('\n')

	fmt.Print("Full Name\t: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Email\t: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Phone\t: ")
	phone, _ := reader.ReadString('\n')

	fmt.Print("Password\t: ")
	password, _ := reader.ReadString('\n')

	req := dto.RegisterDTO{
		Name:     strings.TrimSpace(name),
		Email:    strings.TrimSpace(email),
		Phone:    strings.TrimSpace(phone),
		Password: strings.TrimSpace(password),
	}

	if err := h.Service.RegisterUser(req); err != nil {
		fmt.Println("error:", err)
	}
}
