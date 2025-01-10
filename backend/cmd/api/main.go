package main

import (
	"backend/internals/config"
	"backend/internals/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Println("Server shutdown successfully")
	done <- true
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	config := config.GetConfig()

	db, err := database.New(
		config.DBURI,
		config.DBMaxOpenConns,
		config.DBMaxIdleConns,
		config.DBMaxIdleTime,
	)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	app := NewApplication(db)

	server := http.Server{
		Addr:         config.ServerAddr,
		Handler:      app.RegisterRoutes(),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
		IdleTimeout:  time.Minute,
	}

	done := make(chan bool, 1)

	go gracefulShutdown(&server, done)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}
