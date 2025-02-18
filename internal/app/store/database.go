package store

import (
	"fmt"
	"transfeed/internal/app/model"
	"transfeed/internal/util"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(databaseURL string) {
	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		util.Logger.Error(err.Error())
	}
	DB = db
	db.Logger.LogMode(logger.Info)
	// 迁移 schema
	er := DB.AutoMigrate(&model.Feed{}, &model.User{}, &model.Entry{}, &model.Translator{})
	if err != nil {
		util.Logger.Fatal(fmt.Sprintf("failed to migrate: %v", er))
	}
}
