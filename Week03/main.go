package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppServer struct{}

func (serv *AppServer) ServeHTTP(http.ResponseWriter, *http.Request) {
	//ignore something here ...
}

type DebugServer struct {}

func (serv *DebugServer) ServeHTTP(http.ResponseWriter, *http.Request) {
	//ignore something here ...
}

func main() {
	g, _ := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 8)
	defer cancel()

	//this is app http server
	appMux := &http.Server{
		Addr: ":8080",
		Handler: &AppServer{},
	}

	g.Go(func() error {
		return appMux.ListenAndServe()
	})

	//this is debug http server
	debugMux := &http.Server{
		Addr:    ":8081",
		Handler: &DebugServer{},
	}

	g.Go(func() error {
		return debugMux.ListenAndServe()
	})

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	_ = appMux.Shutdown(ctx)
	_ = debugMux.Shutdown(ctx)
	_ = g.Wait()
}
