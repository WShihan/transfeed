package worker

import (
	"sync"
	"transfeed/internal/util"
)

type Pool struct {
	MaxWorker int
	WG        *sync.WaitGroup
}

func (p *Pool) processJob(sem chan struct{}, wg *sync.WaitGroup, data *interface{}, handler func(*interface{})) {
	defer func() {
		wg.Done()
		sem <- struct{}{} // 获取信号量
		<-sem             // 释放信号量
		util.Logger.Info("--- Mircojob done ---")
	}()
	util.Logger.Info("--- Mircojob start ---")
	handler(data)
}

func (p *Pool) StartJob(dataCollection *[]interface{}, handler func(*interface{})) {

	sem := make(chan struct{}, p.MaxWorker)
	for _, data := range *dataCollection {
		p.WG.Add(1)
		go p.processJob(sem, p.WG, &data, handler)
	}
	p.WG.Wait()
	util.Logger.Info("== job finish ===")

}
func NewPool(maxWorker int) *Pool {
	return &Pool{
		MaxWorker: maxWorker,
		WG:        &sync.WaitGroup{},
	}
}
