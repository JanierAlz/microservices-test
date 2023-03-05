package history

import (
	"fmt"
)

type LoginHistory struct {
	Username   string `json:"sername"`
	IsLogged   bool   `json:"isLogged"`
	LoginDate  int64  `json:"loginDate"`
	LogoutDate int64  `json:"logoutDate"`
}

func AddLoginRecord(username string) error {
	err := CreateLog(username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateLogRecord(username string) error {
	err := UpdateLog(username, false)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
