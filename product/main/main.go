package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"venvision/internal/app"
	"venvision/internal/config"
	"venvision/pkg/logging"
)

func main() {
	log.Print("config initialization")
	cfg := config.NewConfig(
		config.WithDebug(true),
		config.WithDevelopment(true),
	)

	log.Printf("logging initialized.")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("running Application")
	go a.Run()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-sigChan
	logger.Println("Received shutdown signal")
	a.Shutdown()
	logger.Println("Shutdown complete")
}
