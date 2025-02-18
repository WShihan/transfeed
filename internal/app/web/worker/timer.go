package worker

import (
	"fmt"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"
	"transfeed/internal/util"

	"github.com/robfig/cron/v3"
)

type Timer struct {
	Cron  *cron.Cron
	Hours int
}

func (t *Timer) Start() {
	c := t.Cron
	c.AddFunc(fmt.Sprintf("CRON_TZ=Asia/Shanghai 0 */%d * * *", t.Hours), func() {
		util.Logger.Info("===refresh all feed==")
		feeds := []*model.Feed{}
		store.DB.Preload("Translator").Find(&feeds)
		for _, feed := range feeds {
			_, err := RefreshFeed(feed)
			if err != nil {
				fmt.Println(err)
			}
		}
		util.Logger.Info("===refresh all feed Finished==")
	})

	c.Start()
	util.Logger.Info("== timer active, execute at every " + fmt.Sprintf("%d", t.Hours) + " hours ==")
}

func NewTimer(hours int) *Timer {
	return &Timer{
		Cron:  cron.New(),
		Hours: hours,
	}
}
