package repository

import (
	"errors"

	"github.com/bayuf/Golang_Authentication-System/dto"
	"github.com/bayuf/Golang_Authentication-System/model"
	"github.com/bayuf/Golang_Authentication-System/utils"
)

type UserRepository struct {
	FilePath string
}

func NewUserRepository(path string) *UserRepository {
	return &UserRepository{FilePath: path}
}

func (r *UserRepository) GetUser() ([]model.User, error) {

	var users []model.User
	if err := utils.DecoderTask(r.FilePath, &users); err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (r *UserRepository) UserRegister(req dto.RegisterDTO) error {
	newUser := []model.User{
		{
			Name:     req.Name,
			Email:    req.Email,
			Phone:    req.Phone,
			Password: req.Password,
		},
	}

	var users []model.User
	if err := utils.DecoderTask(r.FilePath, &users); err != nil {
		return err
	}

	if utils.IsEmailExist(users, newUser) {
		return errors.New("email has been used")
	}

	users = append(users, newUser...)

	if err := utils.EncoderTask(r.FilePath, users); err != nil {
		return err
	}

	return nil
}
