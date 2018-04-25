package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/schepelin/imageresizer/pkg/storage"
	"time"
)

//go:generate go-bindata -prefix ../.. -pkg migrations -o ../migrations/migrations.go ../../migrations

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db}
}

func (ps *PostgresStorage) Create(ctx context.Context, imgModel *storage.ImageModel) error {
	var err error
	_, err = ps.DB.Exec(
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
	err := ps.DB.QueryRow("SELECT raw, created_at FROM images WHERE id=$1", id).Scan(
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
	_, err := ps.DB.Exec("DELETE FROM images WHERE id=$1", id)
	return err
}
