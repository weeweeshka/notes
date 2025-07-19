package buildApp

import (
	"fmt"
	runGrpc "github.com/weeweeshka/notes/internal/app/grpcApp"
	buisnessLogic "github.com/weeweeshka/notes/internal/buisnesLogic/note"
	"github.com/weeweeshka/notes/internal/storage"
	"log/slog"
)

type App struct {
	GRPCServer *runGrpc.GrpcApp
}

func NewApp(port int, storagePath string, slog *slog.Logger) (error, *App) {
	const op = "buildApp.New"

	postgres, err := storage.New(storagePath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err), nil
	}

	noteService := buisnessLogic.New(slog, postgres, postgres, postgres)
	grpcAPP := runGrpc.New(port, slog, noteService) // + реализация с бизнес логики

	return nil, &App{GRPCServer: grpcAPP}
}
