package store

import (
	"transfeed/internal/app/model"
	"transfeed/internal/util"
)

func ResetUser(username string, password string) error {
	user := model.User{}
	DB.Where("username = ?", username).Find(&user)
	user.Token = nil
	hashedPass, err := util.HashMessage(password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	return DB.Save(&user).Error
}
