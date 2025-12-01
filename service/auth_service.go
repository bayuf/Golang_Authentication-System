package service

import (
	"errors"

	"github.com/bayuf/Golang_Authentication-System/dto"
	"github.com/bayuf/Golang_Authentication-System/model"
	"github.com/bayuf/Golang_Authentication-System/repository"
	"github.com/bayuf/Golang_Authentication-System/utils"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) UserLogin(req dto.LoginDTO) (model.User, error) {
	// Logic Here
	users, err := s.Repo.GetUser()
	if err != nil {
		return model.User{}, err
	}

	// find email
	for _, v := range users {
		if v.Email == req.Email {
			// check password
			if utils.CheckPassword(v.Password, req.Password) {
				return model.User{
					Name: v.Name,
				}, nil
			} else {
				return model.User{}, errors.New("password not match")
			}

		}

	}

	return model.User{}, errors.New("email not found")
}

func (s *AuthService) RegisterUser(req dto.RegisterDTO) error {
	// Logic Here
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	newUser := dto.RegisterDTO{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPass,
	}
	if err := s.Repo.UserRegister(newUser); err != nil {
		return err
	}
	return nil
}
