package model

import (
	"time"

	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	Author               string      `json:"author"`
	Url                  string      `json:"url"`
	Public               bool        `json:"public"`
	Logo                 string      `json:"logo"`
	Email                string      `json:"email"`
	FromLang             string      `json:"fromLang"`
	ToLang               string      `json:"toLang"`
	UserID               int         `json:"userId"`
	TranslatorID         int         `json:"translatorId"`
	Translator           *Translator `json:"translator"`
	Updated              *time.Time  `json:"updated"`
	CreateTime           *time.Time  `json:"createTime"`
	RefreshTime          *time.Time  `json:"refreshTime"`
	TranslateTitle       bool        `json:"translateTitle"`
	TranslateDescription bool        `json:"translateDescription"`
	EntryCount           int         `json:"entrCount"`
	Entries              []*Entry    `json:"entries"`
}
