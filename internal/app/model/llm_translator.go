package model

import (
	"time"

	"gorm.io/gorm"
)

type Translator struct {
	gorm.Model
	Name       string     `json:"name"`
	Key        string     `json:"key"`
	Url        string     `json:"url"`
	Prompt     string     `json:"prompt"`
	Comsume    int64      `json:"comsume"`
	Role       string     `json:"role"`
	UserID     int        `json:"userId"`
	Updated    *time.Time `json:"updated"`
	CreateTime *time.Time `json:"createTime"`
	Feeds      []*Feed    `json:"feeds"`
	Lang       string     `json:"lang"`
}
