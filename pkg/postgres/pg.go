package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/schepelin/imageresizer/pkg/storage"
	"time"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(connect string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db}, nil
}

func (ps *PostgresStorage) Create(ctx context.Context, imgModel *storage.ImageModel) error {
	_, err := ps.DB.Exec(
		`INSERT INTO images(id, raw, created_at) VALUES($1, $2, $3)`,
		imgModel.Id,
		string(imgModel.Raw),
		imgModel.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStorage) Get(ctx context.Context, id string) (*storage.ImageModel, error) {
	var raw string
	var createdAt time.Time
	err := ps.DB.QueryRow("SELECT raw, created_at FROM images WHERE id=?", id).Scan(
		&raw,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}
	return &storage.ImageModel{
		Id:        id,
		Raw:       []byte(raw),
		CreatedAt: createdAt,
	}, nil
}

func (ps *PostgresStorage) Delete(ctx context.Context, id string) error {
	_, err := ps.DB.Exec("DELETE FROM images WHERE id=?", id)
	return err
}
