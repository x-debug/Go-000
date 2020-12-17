package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/x-debug/Go-000/Week04/api/myapp"
	"github.com/x-debug/Go-000/Week04/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//init db layer

	_, cleanup, _ := internal.InitializeDB()
	defer cleanup()

	//init biz layer

	router := gin.Default()

	//register routers
	router.GET("/user/:id", myapp.GetUser)

	srv := &http.Server{
		Addr:    ":8282",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
