package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/schepelin/imageresizer/pkg/storage"
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

