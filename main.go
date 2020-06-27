package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charlires/go-app/controller"
	"github.com/charlires/go-app/router"
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
)

func main() {

	var configFile string
	flag.StringVar(&configFile, "config-file",
		"/etc/config/config.yml", "Path to config file")
	flag.Parse()

	config := viper.New()
	config.SetConfigFile(configFile)
	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Failed to load config: %w", err)
	}

	demoController := controller.NewDemo(render.New())

	httpRouter := router.Setup(
		demoController,
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	// Build HTTP server object
	server := http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%s", config.GetString("http_port")),
		Handler:           handlers.LoggingHandler(os.Stdout, httpRouter),
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	// Launch server start in a separate goroutine
	go func() {
		log.Printf("starting server in address, %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("starting server: %w", err)
		}
	}()

	// Catch application's interrupt signals (Kill, Hang up and Interrupt)
	<-stop

	// Gracefully shutdown using "new" .Shutdown() method
	log.Println("shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("shutting down server gracefully: %w", err)
		return
	}
	log.Println("server gracefully stopped")
}
