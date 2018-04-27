package resizer

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/png"
	"time"
)

//go:generate mockgen -source=resizer.go -destination ../mocks/mock_resizer.go -package mocks
type Image struct {
	Id        string
	Image     image.Image
	CreatedAt time.Time
}

type ResizeJob struct {
	Id uint64
	ImageId string
	Status string
	Image image.Image
	CreatedAt time.Time
	// TODO: extend the service object if necessary
}

type ImageService interface {
	Create(ctx context.Context, raw *[]byte) (*Image, error)
	Read(ctx context.Context, imgId string) (*Image, error)
	Delete(ctx context.Context, imgId string) error
	ScheduleResizeJob(ctx context.Context, imgId string, width, height uint32) (*ResizeJob, error)
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

type ClockUTC struct{}

func (c ClockUTC) Now() time.Time {
	return time.Now().UTC()
}

type ConverterPNG struct{}

func (cnv ConverterPNG) Transform(raw *[]byte) (image.Image, error) {
	buf := bytes.NewBuffer(*raw)
	img, err := png.Decode(buf)
	if err != nil {
		return nil, err
	}
	return img, nil
}

type HasherMD5 struct {}

func (h HasherMD5) Gen(raw *[]byte) string {
	hash := md5.New()
	hash.Write(*raw)

	return hex.EncodeToString(hash.Sum(nil))
}
