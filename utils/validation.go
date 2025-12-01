package utils

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/bayuf/Golang_Authentication-System/model"
)

func UserNameValid(username string) (string, error) {
	re := regexp.MustCompile("^[A-Za-z]+$")

	username = strings.TrimSpace(username)
	if len(username) < 6 {
		return "", errors.New("username must be more than 5 character")
	}

	if !re.MatchString(username) {
		return "", errors.New("username must be character A-z")
	}

	return username, nil
}

func IsEmailExist(users, newUser []model.User) bool {

	userMap := makeMapUser(users)

	if _, exist := userMap[newUser[0].Email]; exist {
		return true
	}

	return false
}

func makeMapUser(users []model.User) map[string]model.User {
	userMap := make(map[string]model.User)

	for _, user := range users {
		userMap[strings.ToLower(user.Email)] = user
	}

	return userMap
}

func EmailValid(email string) (string, error) {
	re := regexp.MustCompile(`.+@.+\..+`)

	email = strings.TrimSpace(email)

	if !re.MatchString(email) {
		return "", errors.New("email format not valid")
	}

	return email, nil
}

func PasswordValid(pass string) (string, error) {

	pass = strings.TrimSpace(pass)

	if len(pass) < 6 {
		return "", errors.New("password must more then 5 character")
	}

	return pass, nil
}

func PhoneNumberValid(phoneNumber string) (string, error) {

	phoneNumber = strings.TrimSpace(phoneNumber)

	for _, num := range phoneNumber {
		if !unicode.IsDigit(num) {
			return "", errors.New("phone must be only number")
		}
	}

	if len(phoneNumber) < 10 || len(phoneNumber) > 15 {
		return "", errors.New("Phone Number is too long or too short")
	}

	return phoneNumber, nil
}
