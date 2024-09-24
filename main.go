package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/szh/naturewatch-fetcher/pkg/agent"
	"github.com/szh/naturewatch-fetcher/pkg/util"
)

func main() {
	loadConfig()
	startAgent()
}

func loadConfig() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal("Cannot get path to executable:", err)
	}
	exPath := filepath.Dir(ex)

	_, err = util.LoadConfig(exPath)
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	err = util.ValidateConfig()
	if err != nil {
		log.Fatal("Invalid config:", err)
	}
}

func startAgent() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		os.Exit(0)
	}()

	agent.Start()
}
