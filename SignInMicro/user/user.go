package user

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/4IDTest/SingIn/history"
	"github.com/4IDTest/SingIn/utils"
)

func AuthUser(username, password string) error {
	var validUser, validPass bool
	users := GetUser(username)
	if len(users) == 0 {
		err := CreateUser(username, password)
		if err != nil {
			return err
		}
		err = history.AddLoginRecord(username)
		if err != nil {
			return err
		}
		return nil
	}
	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))
	for _, userData := range users {
		udNameHash := sha256.Sum256([]byte(userData.Username))
		udPassHash := sha256.Sum256([]byte(userData.Password))
		validUser = subtle.ConstantTimeCompare(udNameHash[:], usernameHash[:]) == 1
		validPass = subtle.ConstantTimeCompare(udPassHash[:], passwordHash[:]) == 1
		if validPass && validUser {
			err := history.AddLoginRecord(username)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return utils.ErrInvalidCredentials
}

func RegisterUser(username, password string) error {
	users := GetUser(username)
	if len(users) > 0 {
		return utils.ErrAlreadyExist
	}
	err := CreateUser(username, password)
	if err != nil {
		return err
	}
	return nil
}
