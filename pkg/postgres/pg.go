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

func (ps *PostgresStorage) CreateResizeJob(ctx context.Context,
	req *storage.ResizeJobRequest) (*storage.ResizeJobResponse, error) {
	var err error
	var rawImg string
	var createdAt time.Time
	var jobId uint64
	tx, err := ps.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.QueryRow("SELECT raw FROM images where id=$1", req.ImgId).Scan(&rawImg)
	switch {
	case err == sql.ErrNoRows:
		return nil, storage.ErrNoImageFound
	case err != nil:
		return nil, err
	}
	err = tx.QueryRow(
		`INSERT INTO resize_jobs(image_id, status, width, height)
				VALUES($1, $2, $3, $4) RETURNING id, created_at`,
		req.ImgId, storage.StatusCreated, req.Width, req.Height,
	).Scan(&jobId, &createdAt)

	if err != nil {
		return nil, err
	}

	return &storage.ResizeJobResponse{
		Id:        jobId,
		Status:    storage.StatusCreated,
		CreatedAt: createdAt,
		RawImg:    []byte(rawImg),
	}, nil

}

func (ps *PostgresStorage) WriteResizeJobResult(ctx context.Context, req *storage.ResizeResultRequest) error {
	var err error
	_, err = ps.DB.Exec(
		"UPDATE resize_jobs SET status = $1, raw = $2 WHERE id=$3",
		storage.StatusFinished,
		string(req.Raw),
		req.JobId,
	)
	return err
}
