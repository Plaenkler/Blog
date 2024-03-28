package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/plaenkler/blog/pkg/server"
)

func main() {
	server.Start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	server.Stop()
}
