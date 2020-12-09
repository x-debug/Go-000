package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppServer struct{}

func (serv *AppServer) ServeHTTP(reply http.ResponseWriter, req *http.Request) {
	//ignore something here ...
	_, _ = reply.Write([]byte("AppServer"))
}

type DebugServer struct {}

func (serv *DebugServer) ServeHTTP(reply http.ResponseWriter, req *http.Request) {
	//ignore something here ...
	_, _ = reply.Write([]byte("DebugServer"))
}

func main() {
	g, _ := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 8)
	defer cancel()

	//this is app http server
	appMux := &http.Server{
		Addr: ":1234",
		Handler: &AppServer{},
	}

	//this is debug http server
	debugMux := &http.Server{
		Addr:    ":1235",
		Handler: &DebugServer{},
	}

	g.Go(func() error {
		err := debugMux.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
		_ = appMux.Shutdown(ctx)
		return err
	})

	g.Go(func() error {
		err := appMux.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
		_ = debugMux.Shutdown(ctx)
		return err
	})

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		_ = appMux.Shutdown(ctx)
		_ = debugMux.Shutdown(ctx)
	}()

	_ = g.Wait()
}
