package store

import (
	"strings"
	"time"
	"transfeed/internal/app/form"
	"transfeed/internal/app/model"

	"github.com/mmcdole/gofeed"
)

func GetFeed(user model.User, id int) (model.Feed, error) {
	feed := model.Feed{}
	err := DB.Preload("Entries").Preload("Translator").Where("user_id = ?", user.ID).First(&feed, id).Error
	return feed, err
}

func GetFeeds(user model.User) ([]*model.Feed, error) {
	feeds := []*model.Feed{}
	err := DB.Preload("Entries").Preload("Translator").Where("user_id = ?", user.ID).Find(feeds).Error
	return feeds, err
}

func UpdateFeed(user model.User, payload form.FeedUpdateForm) (*model.Feed, error) {
	feed, err := GetFeed(user, payload.ID)
	if err != nil {
		return nil, err
	}
	if payload.Url != feed.Url {
		feed.Url = payload.Url
	}
	if payload.Title != feed.Title {
		feed.Title = payload.Title
	}
	if payload.Description != feed.Description {
		feed.Description = payload.Description
	}
	if payload.Public != feed.Public {
		feed.Public = payload.Public
	}
	if payload.Logo != "" {
		feed.Logo = payload.Logo
	}
	if payload.FromLang != "" {
		feed.FromLang = payload.FromLang
	}
	if payload.ToLang != "" {
		feed.ToLang = payload.ToLang
	}
	if feed.TranslatorID != payload.TranslatorId {
		feed.TranslatorID = payload.TranslatorId
		trans := &model.Translator{}
		err = DB.First(trans, feed.TranslatorID).Error
		if err != nil {
			trans = nil
		}
		DB.Model(&feed).Association("Translator").Replace(trans)
	}
	err = DB.Save(&feed).Error
	return &feed, err
}

func CreateFeed(feedAddForm form.FeedAddForm) (*model.Feed, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/237.84.2.178 Safari/537.36"
	feedSrc, err := fp.ParseURL(feedAddForm.Url)
	if err != nil {
		return nil, err
	}
	translator := &model.Translator{}
	err = DB.Model(&translator).Where("id = ?", feedAddForm.TranslatorID).First(translator).Error
	if err != nil {
		translator = nil
	}
	now := time.Now()
	feedObj := model.Feed{
		Title:                feedSrc.Title,
		Url:                  feedAddForm.Url,
		Description:          feedSrc.Description,
		TranslateTitle:       feedAddForm.TranslateTitle,
		TranslateDescription: feedAddForm.TranslateDescription,
		FromLang:             feedSrc.Language,
		Public:               feedAddForm.Public,
		Translator:           translator,
		ToLang:               feedAddForm.ToLang,
		CreateTime:           &now,
	}
	if feedSrc.Author != nil {
		feedObj.Author = feedSrc.Author.Name
		feedObj.Email = feedSrc.Author.Email
	}
	if feedSrc.Image != nil {
		feedObj.Logo = feedSrc.Image.URL
	}
	if err != nil {
		return nil, err
	}

	entries := []*model.Entry{}
	for _, item := range feedSrc.Items {
		entry := model.Entry{
			Guid:    item.GUID,
			Title:   item.Title,
			Link:    item.Link,
			Summary: item.Description,
		}
		if item.Author != nil {
			entry.Author = item.Author.Name
		}
		if item.PublishedParsed != nil {
			entry.Pubdate = item.PublishedParsed
		}
		if len(item.Categories) > 0 {
			entry.Category = strings.Join(item.Categories, ",")
		}
		entries = append(entries, &entry)
	}

	feedObj.Entries = entries

	return &feedObj, nil
}
