package store

import (
	"fmt"
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
	err = DB.Save(&user).Error
	if err != nil {
		fmt.Printf("reset user error:%s\n", err.Error())
	}
	if user.Username == "" {
		fmt.Printf("user: %s not exist!\n", username)
		return fmt.Errorf("user not exist")
	}
	fmt.Printf("reset user: %s success\n", username)
	return err
}
