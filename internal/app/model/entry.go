package model

import (
	"time"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Guid       string     `json:"guid"`
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	Link       string     `json:"llink"`
	Summary    string     `json:"summary"`
	Category   string     `json:"category"`
	FeedID     uint       `json:"feed_id"`
	CreateTime *time.Time `json:"createTime"`
	Pubdate    *time.Time `json:"pubDate"`
}
