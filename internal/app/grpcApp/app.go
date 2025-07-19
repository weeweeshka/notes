package grpcApp

import (
	"fmt"
	grpcHandlers "github.com/weeweeshka/notes/internal/grpc/note"
	"github.com/weeweeshka/notes/internal/middleware"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type GrpcApp struct {
	gRPCServer *grpc.Server
	slog       *slog.Logger
	port       int
}

func New(port int, slog *slog.Logger, noteService grpcHandlers.Note) *GrpcApp {
	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.SimpleTraceIDInterceptor())) // создается grpc сервер с дефолт настройками
	grpcHandlers.RegisterServer(gRPCServer, noteService)                                       // регистрирует на этот сервер нашу реализацию, т.е инкапсулируем сюда

	return &GrpcApp{
		gRPCServer,
		slog,
		port}
	// возвращаем наш grpc, уже со связанным интерфейсом
}

func (a *GrpcApp) Run() error {
	const op = "grpc.Run"

	log := a.slog.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server")
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *GrpcApp) MustRun() {
	_ = a.Run()
}

func (a *GrpcApp) GracefulStop() {
	const op = "gracefulStop"
	a.slog.With(slog.String("op", op)).Info("stopping grpc server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
