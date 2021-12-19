package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"time"

	"golang.org/x/sync/errgroup"
)

type Server func(context.Context) error

func main() {
	//use errgroup with context
	group, ctx := errgroup.WithContext(context.Background())

	//start some goroutine
	mux := http.NewServeMux()
	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "homepage")
	})

	//shutdown the server
	stop := make(chan bool)
	mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "shutdown")
		stop <- true
	})

	//create http server
	server := http.Server{Handler: mux, Addr: "127.0.0.1:8080"}

	//start http server
	group.Go(func() error {
		return server.ListenAndServe()
	})

	//add signak interrupt operation
	group.Go(func() error {
		timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			fmt.Println("ctx is done")
		case <-stop:
			fmt.Println("server will shutdown")
		case <-quit:
			fmt.Println("process killed")
		}

		return server.Shutdown(timeoutCtx)
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
}
