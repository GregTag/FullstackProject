package main

import (
	handler "backend/internal/handler/http"
	"backend/internal/repository"
	repository_sqlite "backend/internal/repository/sqlite"
	"backend/internal/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var wait = time.Second * 15

func main() {
	db, err := repository_sqlite.NewSQLiteDB("umo.db")
	if err != nil {
		log.Panicf("Failed to initialize database: %s\n", err.Error())
	} else {
		log.Println("Database is initialized")
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	h := handler.NewHandler(service)

	server := &http.Server{
		Addr: "127.0.0.1:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)

	// Block until we receive our signal.
	<-channel

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Shutting down")
	os.Exit(0)
}
