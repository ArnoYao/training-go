package main

import (
	"log"
	"time"

	. "github.com/arnoyao/training-go/05week/hystrix"
)

func main() {
	window := NewSlidingWindow()
	for _, request := range []int64{1, 2, 3, 4, 5} {
		window.Increment(request)
		time.Sleep(1 * time.Second)
	}
	log.Println(window.Avg(time.Now()))
}
