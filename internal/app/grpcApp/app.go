package grpcApp

import (
	grpcHandlers "github.com/weeweeshka/notes/internal/grpc/note"
	"google.golang.org/grpc"
	"log/slog"
)

type grpcApp struct {
	gRPCServer *grpc.Server
	slog       *slog.Logger
	port       int
}

func New(port int, slog *slog.Logger, noteService grpcHandlers.Note) *grpcApp {}
