package storage

import (
	"context"
	"errors"
	"time"
)

//go:generate mockgen -source=storage.go -destination ../mocks/mock_storage.go -package mocks

var (
	ErrNoImageFound = errors.New("no image found")
)

type ResizeStorage interface {
	WriteResizeJobResult(ctx context.Context, req *ResizeResultRequest) error
	GetResizeJob(ctx context.Context, req *ResizeGetRequest) (*ResizeJobResponse, error)
}

type Storage interface {
	ResizeStorage

	Create(ctx context.Context, model *ImageModel) error
	Get(ctx context.Context, id string) (*ImageModel, error)
	Delete(ctx context.Context, id string) error
	CreateResizeJob(ctx context.Context, req *ResizeJobRequest) (*ResizeJobResponse, error)
}

type ImageModel struct {
	Id        string
	Raw       []byte
	CreatedAt time.Time
}

type ResizeJobRequest struct {
	ImgId  string
	Width  uint
	Height uint
}

type ResizeJobResponse struct {
	Id        uint64
	Status    string
	CreatedAt time.Time
	RawImg    []byte
}

type ResizeGetRequest struct {
	JobId uint64
}

type ResizeResultRequest struct {
	JobId uint64
	Raw   []byte
}
