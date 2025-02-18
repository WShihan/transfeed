package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int        `json:"id"`
	Username    string     `gorm:"unique" json:"username"`
	Password    string     `json:"password"`
	Token       *string    `json:"token"`
	LastLogin   *time.Time `json:"lastLogin"`
	Feeds       []*Feed
	Translators []*Translator
	Admin       bool `json:"admin"`
}
