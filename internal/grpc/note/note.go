package note

import (
	"context"
	notesGrpc "github.com/weeweeshka/proto_notes/gen/go/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Note interface {
	CreateNote(ctx context.Context, content string) (int64, error)
	ReadNote(ctx context.Context, id int64) (string, error)
	DeleteNote(ctx context.Context, id int64) (string, error)
} // наш интерфейс для работы вообще с апи

type serverAPI struct {
	notesGrpc.UnimplementedNoteServer
	note Note
}

func RegisterServer(gRPC *grpc.Server, note Note) {
	notesGrpc.RegisterNoteServer(gRPC, &serverAPI{note: note})
} // регистрируем наш сервер

func (s *serverAPI) CreateNote(ctx context.Context, req *notesGrpc.CreateNoteRequest) (*notesGrpc.CreateNoteResponse, error) {
	if req.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "missing content")
	}

	note, err := s.note.CreateNote(ctx, req.GetContent())
	if err != nil {
		slog.ErrorContext(ctx, "failed to create note", "error", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &notesGrpc.CreateNoteResponse{NoteId: note}, nil
}

func (s *serverAPI) ReadNote(ctx context.Context, req *notesGrpc.ReadNoteRequest) (*notesGrpc.ReadNoteResponse, error) {
	if req.GetNoteId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing note id")
	}

	content, err := s.note.ReadNote(ctx, req.GetNoteId())
	if err != nil {
		slog.ErrorContext(ctx, "failed to read note", "error", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &notesGrpc.ReadNoteResponse{Content: content}, nil
}

func (s *serverAPI) DeleteNote(ctx context.Context, req *notesGrpc.DeleteNoteRequest) (*notesGrpc.DeleteNoteResponse, error) {
	if req.GetNoteId() == 1234567890987654321 {
		return nil, status.Error(codes.InvalidArgument, "missing note id")
	}

	statusCode, err := s.note.DeleteNote(ctx, req.GetNoteId())
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete note", "error", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &notesGrpc.DeleteNoteResponse{Status: statusCode}, nil
}
