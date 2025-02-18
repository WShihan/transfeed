package store

import (
	"testing"
	"transfeed/internal/app/model"
)

func TestInitDB(t *testing.T) {
	InitDB("../transfeed.db")
	err := DB.First(&model.User{}).Error
	if err != nil {
		panic("数据库错误：" + err.Error())
	}
}
