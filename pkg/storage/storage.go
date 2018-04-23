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
}

type ImageModel struct { // rename
	Id        string
	Raw       []byte
	CreatedAt time.Time
}
