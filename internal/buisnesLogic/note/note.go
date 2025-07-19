package note

import (
	"context"
	"fmt"
	"github.com/weeweeshka/notes/internal/middleware"
	"log/slog"
)

type Note struct {
	log         *slog.Logger
	noteCreater NoteCreater
	noteReader  NoteReader
	noteDeleter NoteDeleter
}

type NoteCreater interface {
	CreateNote(ctx context.Context, content string) (int64, error)
}

type NoteReader interface {
	ReadNote(ctx context.Context, id int64) (string, error)
}

type NoteDeleter interface {
	DeleteNote(ctx context.Context, id int64) (string, error)
}

func New(log *slog.Logger, noteCreater NoteCreater, noteReader NoteReader, notDeleter NoteDeleter) *Note {
	return &Note{
		log:         log,
		noteCreater: noteCreater,
		noteReader:  noteReader,
		noteDeleter: notDeleter,
	}
} // пока эта функция не создаст объект, все методы недоступны

func (n *Note) CreateNote(ctx context.Context, content string) (int64, error) {
	const op = "buisnesLogic.note.CreateNote"
	traceID := middleware.TraceIDFromContext(ctx)

	log := slog.With(slog.String("op", op), slog.String("trace_id", traceID))

	log.Info("Creating note")

	noteID, err := n.noteCreater.CreateNote(ctx, content)
	if err != nil {
		log.Info("failed to create note")
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("note created")

	return noteID, nil

}

func (n *Note) ReadNote(ctx context.Context, id int64) (string, error) {
	const op = "buisnesLogic.note.GetNote"
	traceID := middleware.TraceIDFromContext(ctx)

	log := slog.With(slog.String("op", op), slog.String("trace_id", traceID))

	log.Info("Getting note")

	note, err := n.noteReader.ReadNote(ctx, id)
	if err != nil {
		log.Info("failed to get note")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("note found")
	return note, nil
}

func (n *Note) DeleteNote(ctx context.Context, id int64) (string, error) {
	const op = "buisnesLogic.note.DeleteNote"
	traceID := middleware.TraceIDFromContext(ctx)

	log := slog.With(slog.String("op", op), slog.String("trace_id", traceID))

	log.Info("Deleting note")

	_, err := n.noteDeleter.DeleteNote(ctx, id)
	if err != nil {
		log.Info("failed to delete note")
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("note deleted")

	return "delete successful", nil
}
