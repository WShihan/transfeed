package worker

import (
	"sort"
	"strings"
	"time"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

func RefreshFeed(feed *model.Feed) (*model.Feed, error) {
	fp := gofeed.NewParser()
	feedSub, err := fp.ParseURL(feed.Url)
	if err != nil {
		return nil, err
	}
	entries := []*model.Entry{}
	for _, item := range feedSub.Items {
		entry := model.Entry{
			Guid:    item.GUID,
			Title:   item.Title,
			Link:    item.Link,
			Summary: item.Content,
			Pubdate: item.PublishedParsed,
		}
		err = store.DB.First(&entry, "guid = ?", item.GUID).Error
		if err == nil {
			continue
		}
		if item.Author != nil {
			entry.Author = item.Author.Name
		}
		if len(item.Categories) > 0 {
			entry.Category = strings.Join(item.Categories, ",")
		}
		entries = append(entries, &entry)
	}
	if len(entries) > 0 {
		PostProcessEntries(feed, entries, 10)

		for _, entry := range entries {
			store.DB.Model(feed).Association("Entries").Append(entry)
		}
	}

	refreshTime := time.Now()
	feed.RefreshTime = &refreshTime
	store.DB.Save(feed)

	return feed, nil
}

func GenFeedRss(id string) (*feeds.Feed, error) {
	items := []*feeds.Item{}
	feed := model.Feed{}
	err := store.DB.Preload("Entries").First(&feed, id).Error
	if err != nil {
		return nil, err
	}
	entires := feed.Entries
	sort.Slice(entires, func(i, j int) bool {
		return entires[i].ID < entires[j].ID
	})
	for _, entry := range entires {
		item := feeds.Item{
			Title:       entry.Title,
			Link:        &feeds.Link{Href: entry.Link},
			Description: entry.Summary,
			Content:     entry.Summary,
			Author:      &feeds.Author{Name: entry.Author},
			Created:     entry.CreatedAt,
			Updated:     *entry.Pubdate,
		}
		if entry.Pubdate != nil {
			item.Created = *entry.Pubdate
		}
		items = append(items, &item)

	}

	var updateTIme time.Time
	if feed.RefreshTime != nil {
		updateTIme = *feed.RefreshTime
	} else {
		updateTIme = feed.CreatedAt
	}
	rss := &feeds.Feed{
		Title:       feed.Title,
		Link:        &feeds.Link{Href: feed.Url},
		Description: feed.Description,
		Author:      &feeds.Author{Name: feed.Author, Email: ""},
		Created:     feed.CreatedAt,
		Updated:     updateTIme,
	}

	rss.Items = items

	return rss, nil

}
