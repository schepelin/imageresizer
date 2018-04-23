package storage

import (
	"context"
	"time"
)

//go:generate mockgen -source=storage.go -destination ../mocks/mock_storage.go -package mocks

type Storage interface {
	Create(ctx context.Context, model *ImageModel) error
}



type ImageModel struct {
	Id string
	Raw []byte
	CreatedAt time.Time
}