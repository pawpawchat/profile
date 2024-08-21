package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawpawchat/profile/config"
	"github.com/pawpawchat/profile/internal/app"
)

func main() {
	flag.Parse()

	// read the config file
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err = config.ConfigureLogger(cfg); err != nil {
		log.Fatal(err)
	}

	// main application context
	ctx, cancel := context.WithCancel(context.Background())

	// catch the signal of the shutdown programm
	// to correctly terminate the app
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	// run the application
	err = app.Run(ctx, cfg)

	if err != nil {
		fmt.Fprintf(os.Stderr, "app was terminated with an error: %s", err.Error())
		os.Exit(1)
	}
}
