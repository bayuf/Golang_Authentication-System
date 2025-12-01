package dto

type RegisterDTO struct {
	Email    string
	Password string
	Name     string
	Phone    string
}

type LoginDTO struct {
	Email    string
	Password string
}
