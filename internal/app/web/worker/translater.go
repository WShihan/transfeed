package worker

import (
	"fmt"
	"sync"
	"transfeed/internal/app/model"
	translate "transfeed/internal/app/translate"
	"transfeed/internal/util"
)

func PostProcessEntries(feed *model.Feed, entries []*model.Entry, maxWorker int) error {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok := r.(error)
			if !ok {
				util.Logger.Error(fmt.Errorf("pkg: %v", r).Error())
				util.Logger.Error(err.Error())
			}
			return
		}
	}()

	var wg sync.WaitGroup
	agent := feed.Translator
	if agent == nil {
		return fmt.Errorf("translator not found")
	}
	translator := translate.LLMTranslator{
		Agent: *agent,
	}

	sem := make(chan struct{}, maxWorker)
	for _, entry := range entries {
		wg.Add(1)
		go translator.Translate(sem, &wg, feed, entry)
	}
	wg.Wait()
	util.Logger.Info("finish")

	return nil
}
