package handler

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bayuf/Golang_Authentication-System/dto"
	"github.com/bayuf/Golang_Authentication-System/service"
	"github.com/bayuf/Golang_Authentication-System/utils"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{Service: svc}
}

func (h *AuthHandler) Login() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- LOGIN ---")
	reader.ReadString('\n')

	fmt.Print("Email\t\t: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Password\t: ")
	password, _ := reader.ReadString('\n')

	validEmail, err := utils.EmailValid(email)
	if err != nil {
		return err
	}

	validPassword, err := utils.PasswordValid(password)
	if err != nil {
		return err
	}

	user, err := h.Service.UserLogin(dto.LoginDTO{
		Email:    validEmail,
		Password: validPassword,
	})
	if err != nil {
		return err
	} else {
		fmt.Println(Green, "Login Berhasil, Selamat datang", user.Name, Reset)
	}

	return nil

}

func (h *AuthHandler) Register() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- REGISTER ---")
	reader.ReadString('\n')

	fmt.Print("Full Name\t: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Email\t\t: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Phone\t\t: ")
	phone, _ := reader.ReadString('\n')

	fmt.Print("Password\t: ")
	password, _ := reader.ReadString('\n')

	validName, err := utils.UserNameValid(name)
	if err != nil {
		return err
	}
	validEmail, err := utils.EmailValid(email)
	if err != nil {
		return err
	}
	validPhone, err := utils.PhoneNumberValid(phone)
	if err != nil {
		return err
	}
	validPassword, err := utils.PasswordValid(password)
	if err != nil {
		return err
	}

	req := dto.RegisterDTO{
		Name:     validName,
		Email:    validEmail,
		Phone:    validPhone,
		Password: validPassword,
	}

	if err := h.Service.RegisterUser(req); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(Green, "Registrasi Berhasil, data tersimpan di users.json", Reset)
	}

	return nil
}
