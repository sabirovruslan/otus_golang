package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/app"
	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/spf13/viper"
)

func main() {
	var (
		configFile = flag.String("config", "/etc/calendar/config.toml", "Path to configuration file")
	)
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed read config file: %v", configFile)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	logg := logger.New(config.Logger.Level)

	storage := memorystorage.New()
	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
		logg.Info("calendar shutdown...")
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
