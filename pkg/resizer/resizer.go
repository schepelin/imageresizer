package resizer

import (
	"image"
	"time"
	"context"
)


//go:generate mockgen -source=resizer.go -destination ../mocks/mock_resizer.go -package mocks
type Image struct {
	Id  string
	Image image.Image
	CreatedAt time.Time
}

type ImageService interface {
	Create(ctx context.Context, raw []byte) (*Image, error)
	Read(ctx context.Context, imgId string) (*Image, error)
	Delete(ctx context.Context, imgId string) error
}


type Clocker interface {
	Now() time.Time
}


type Hasher interface {
	Gen(raw *[]byte) string
}

type Converter interface {
	Transform(raw *[]byte) (image.Image, error)
}