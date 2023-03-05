package user

import (
	"github.com/4IDTest/SingIn/initializers"
	"github.com/4IDTest/SingIn/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func GetUser(username string) []User {
	var users []User
	initializers.DB.Where("username= ?", username).Find(&users)
	if len(users) > 0 {
		return users
	}
	return []User{}
}

func CreateUser(username, password string) error {
	newUser := User{Username: username, Password: password}
	result := initializers.DB.Omit("UpdatedAt").Create(&newUser)
	if result.RowsAffected == 0 {
		return utils.ErrInsertData
	}
	return nil
}
