package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-flight-search/internal/app"
)

func main() {
	// Create context that listens for OS signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize the app (DI, router, etc.) and get HTTP server
	server := app.Run()

	log.Println("go-flight-search API started on :8080")

	// Run HTTP server in goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	// Wait until we receive a shutdown signal
	<-ctx.Done()
	log.Println("shutdown signal received")

	// Create context with timeout for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try to gracefully shut down the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("server forced to shutdown:", err)
	} else {
		log.Println("server shutdown gracefully")
	}
}
