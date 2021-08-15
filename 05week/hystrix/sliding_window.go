package hystrix

import (
	"log"
	"sync"
	"time"
)

type SlidingWindow struct {
	Windows map[int64]*window
	Mutex   *sync.RWMutex
}

type window struct {
	Value int64
}

func NewSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		Windows: make(map[int64]*window),
		Mutex:   &sync.RWMutex{},
	}
}

func (sw *SlidingWindow) getCurrentWindow() *window {
	now := time.Now().Unix()

	var w *window
	var ok bool
	if w, ok = sw.Windows[now]; !ok {
		w = &window{}
		sw.Windows[now] = w
	}

	return w
}

func (sw *SlidingWindow) removeOldWindows() {
	now := time.Now().Unix() - 10
	for timestamp := range sw.Windows {
		if timestamp <= now {
			delete(sw.Windows, timestamp)
		}
	}
}

func (sw *SlidingWindow) Increment(i int64) {
	if i == 0 {
		return
	}

	sw.Mutex.Lock()
	defer sw.Mutex.Unlock()

	b := sw.getCurrentWindow()
	b.Value += i
	sw.removeOldWindows()
}

func (sw *SlidingWindow) Sum(now time.Time) (sum int64) {
	sw.Mutex.Lock()
	defer sw.Mutex.Unlock()

	for timestamp, window := range sw.Windows {
		if timestamp > now.Unix()-10 {
			log.Println(sum)
			sum += window.Value
		}
	}

	return sum
}

func (sw *SlidingWindow) Avg(now time.Time) int64 {
	return sw.Sum(now) / 10
}
