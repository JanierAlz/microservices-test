package history

import (
	"github.com/4IDTest/SingIn/initializers"
	"github.com/4IDTest/SingIn/utils"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	Username string
	IsLogged bool
}

func GetLogs(args map[string]any) []History {
	var log []History
	initializers.DB.Where(args).Find(&log)
	if len(log) > 0 {
		return log
	}
	return []History{}
}

func CreateLog(username string) error {
	newLog := History{Username: username, IsLogged: true}
	result := initializers.DB.Omit("UpdatedAt").Create(&newLog)
	if result.RowsAffected == 0 {
		return utils.ErrInsertData
	}
	return nil
}

func UpdateLog(username string, isLogged bool) error {
	logs := GetLogs(map[string]any{"username": username, "is_logged": true})
	if len(logs) == 0 {
		return utils.ErrLoggedOut
	}
	result := initializers.DB.Model(&logs[0]).Where("is_logged=?", true).Update("is_logged", &isLogged)
	if result.RowsAffected == 0 {
		return utils.ErrInsertData
	}
	return nil
}
