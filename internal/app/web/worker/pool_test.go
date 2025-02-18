package worker

import (
	"fmt"
	"testing"
	"time"
)

type point struct {
	X int
	Y int
}

func job(data *interface{}) {
	time.Sleep(3 * time.Second)
	pt := (*data).(point)
	sum := pt.X + pt.Y
	fmt.Println(sum)

}

func TestPool(t *testing.T) {
	data := []interface{}{}
	for i := 0; i <= 20; i++ {
		data = append(data, point{i, i})
	}
	pool := NewPool(5)
	pool.StartJob(&data, job)

	// for _, d := range data {
	// 	job(&d)
	// }

}
