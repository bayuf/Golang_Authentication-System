package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/bayuf/Golang_Authentication-System/model"
)

func UserNameValid(username string) (bool, error) {
	re := regexp.MustCompile("^[A-Za-z]+$")

	username = strings.TrimSpace(username)
	if len(username) < 6 {
		return false, errors.New("username must be more than 5 character")
	}

	if !re.MatchString(username) {
		return false, errors.New("username must be character A-z")
	}

	return true, nil
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
