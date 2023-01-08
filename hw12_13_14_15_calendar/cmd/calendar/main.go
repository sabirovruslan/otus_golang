package main

import (
	"context"
	"flag"
	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/migrate"
	sqlstorage "github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/storage/sql"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/app"
	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/config"
	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/server/http"
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

	conf := config.NewConfig()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed read config file: %v", configFile)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	logg := logger.New(conf.Logger.Level)

	migrator, err := migrate.NewPgMigrate(&conf.Storage.Database)
	if err != nil {
		log.Fatal(err)
	}
	if err := migrator.Run(); err != nil {
		log.Fatal(err)
	}

	store, err := sqlstorage.New(&conf.Storage.Database)
	if err != nil {
		log.Fatal(err)
	}

	calendar := app.New(logg, store)

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
