package main

import (
	"github.com/weeweeshka/notes/internal/app/buildApp"
	"github.com/weeweeshka/notes/internal/config"
	"log/slog"
	"os/signal"
	"syscall"

	"os"
)

func main() {
	cfg := config.MustLoad()
	slog.Info("Config loaded")

	slogger := SetupLogger()
	slog.Info("Logger loaded")

	err, app := buildApp.NewApp(cfg.Port, cfg.StoragePath, slogger)
	if err != nil {
		panic(err)
	}

	go app.GRPCServer.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	slog.Info("Gracefully shutting down...")
	<-stop
	slog.Info("App stopped")

}

func SetupLogger() *slog.Logger {
	var log *slog.Logger

	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
