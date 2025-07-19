package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

type Storage struct {
	db *pgx.Conn
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.New"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	_, err = conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS notes (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    content TEXT NOT NULL)`)
	if err != nil {
		slog.Info("error creating notes table")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	slog.Info("notes table created")
	return &Storage{conn}, nil
}

func (s *Storage) CreateNote(ctx context.Context, content string) (int64, error) {
	const op = "storage.CreateNote"

	var id int64
	slog.With(slog.String("op", op)).Info("creating note")
	err := s.db.QueryRow(ctx, `INSERT INTO notes (content) VALUES ($1) RETURNING id`, content).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) DeleteNote(ctx context.Context, id int64) (string, error) {
	const op = "storage.DeleteNote"

	slog.With(slog.String("op", op)).Info("deleting note")
	err := s.db.QueryRow(ctx, `DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return "success", nil
}

func (s *Storage) ReadNote(ctx context.Context, id int64) (string, error) {
	const op = "storage.ReadNote"
	var content string

	slog.With(slog.String("op", op)).Info("reading note")
	err := s.db.QueryRow(ctx, `SELECT content FROM notes WHERE id = $1`, id).Scan(&content)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return fmt.Sprintf("%v", content), nil
}
