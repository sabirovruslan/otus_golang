package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	host, port := flag.Arg(0), flag.Arg(1)
	conn := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)
	if err := conn.Connect(); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		err := conn.Send()
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err := conn.Receive()
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
}
