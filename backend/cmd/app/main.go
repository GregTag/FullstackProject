package main

import (
	handler "backend/internal/handler/http"
	"backend/internal/repository"
	repository_sqlite "backend/internal/repository/sqlite"
	"backend/internal/service"
	"backend/pkg/auth"
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var wait = time.Second * 15

func main() {
	logFile, err := os.OpenFile("./log/backend.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Panicf("Error opening file: %s\n", err.Error())
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	dbPath := flag.String("db", "umo.db", "Path to SQLite database")
	jwtKeyPath := flag.String("jwt", "jwt.key", "Path to JWT key")
	flag.Parse()

	db, err := repository_sqlite.NewSQLiteDB(*dbPath)
	if err != nil {
		log.Panicf("Failed to initialize database: %s\n", err.Error())
	} else {
		log.Println("Database is initialized")
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)

	signingKey, err := os.ReadFile(*jwtKeyPath)
	if err != nil {
		log.Panicf("Failed to read JWT key: %s\n", err.Error())
	}
	auth := auth.NewAuthManager(signingKey)

	h := handler.NewHandler(service, auth)

	server := &http.Server{
		Addr: "0.0.0.0:8000",
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
