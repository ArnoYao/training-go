package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	group, ctx := errgroup.WithContext(context.Background())

	// 1、server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	// 模拟server的请求
	shutdown := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		shutdown <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	group.Go(func() error {
		return server.ListenAndServe()
	})

	// 2、控制server退出行为
	group.Go(func() error {

		select {
		case <-ctx.Done():
			log.Println("ctx done 2")
		case <-shutdown:
			log.Println("server shutdown")
		}

		log.Println("server即将退出")

		return server.Shutdown(ctx)
	})

	// 3、接收信号量
	group.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit)

		for {
			select {
			case <-ctx.Done():
				log.Println("ctx done 3")
				return ctx.Err()
			case sig := <-quit:
				log.Printf("信号量：%v \n", sig) // lsof -i:8080 & kill -1 pid
				return errors.Errorf("信号量退出", sig)
			}
		}
	})

	if err := group.Wait(); err != nil {
		log.Println("group error: ", err)
	}

	log.Println("ok")
}
