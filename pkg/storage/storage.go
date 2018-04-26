package storage

import (
	"context"
	"time"
)

//go:generate mockgen -source=storage.go -destination ../mocks/mock_storage.go -package mocks

type Storage interface {
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
	ImgId string
	Width uint32
	Height uint32
}

type ResizeJobResponse struct {
	Id uint64
	Status string
	CreatedAt time.Time
}
