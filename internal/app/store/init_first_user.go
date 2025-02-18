package store

import (
	"transfeed/internal/app/model"
	"transfeed/internal/util"
)

func InitAdmin(username string, password string) {
	user := model.User{}
	DB.Where("username = ?", username).Find(&user)
	toke, _ := util.GenerateRandomKey(12)
	pass, err := util.HashMessage(password)
	if err != nil {
		panic("init admin error：" + err.Error())
	}
	if user.Username == "" || user.Password == "" {
		user.Token = &toke
		user.Username = username
		user.Password = pass
		user.Admin = true
		DB.Create(&user)
	}
}
