package main

import (
	"errors"
	"fmt"
)

var UserInfo map[string] string
func init() {
	UserInfo = map[string] string {
		"Joe": "5566",
	}
}

func CheckUserIsExist(username string) bool {
	a, isExist := UserInfo[username]
	fmt.Println(a, isExist, "check user is exist")
	return isExist
}

func CheckPassword(pd1 string, pd2 string) error {
	if pd1 == pd2 {
		return nil
	} else {
		return errors.New("password is error!")
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserInfo[username], password)
	} else {
		return errors.New("user is not exist!")
	}
}