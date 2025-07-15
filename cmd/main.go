package cmd

import (
	"log/slog"
	"notes/internal/config"
	"os"
)

func main() {
	config.MustLoad()
	slog.Info("Config loaded")

	slogger := SetupLogger()
	slog.Info("Logger loaded")

}

func SetupLogger() *slog.Logger {
	var log *slog.Logger

	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
