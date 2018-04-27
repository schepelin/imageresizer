package resizer

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"time"
)

//go:generate mockgen -source=resizer.go -destination ../mocks/mock_resizer.go -package mocks

var (
	ErrImgDecode   = errors.New("could not decode image")
	StatusCreated  = "CREATED"
	StatusFinished = "FINISHED"
	StatusFailed   = "FAILED"
)

type Image struct {
	Id        string
	Image     image.Image
	CreatedAt time.Time
}

type ResizeJob struct {
	Id        uint64
	Status    string
	Image     image.Image
	CreatedAt time.Time
	// TODO: extend the service object if necessary
}

type ResizeServiceRequest struct {
	JobId  uint64
	RawImg []byte
	Width  uint
	Height uint
}

type ImageService interface {
	Create(ctx context.Context, raw *[]byte) (*Image, error)
	Read(ctx context.Context, imgId string) (*Image, error)
	Delete(ctx context.Context, imgId string) error
	ScheduleResizeJob(ctx context.Context, imgId string, width, height uint) (*ResizeJob, error)
	GetResizeJob(ctx context.Context, jobId uint64) (*ResizeJob, error)
}

type Clocker interface {
	Now() time.Time
}

type Hasher interface {
	Gen(raw *[]byte) string
}

type Converter interface {
	Transform(raw *[]byte) (image.Image, error)
	Resize(img *image.Image, width, height uint) ([]byte, error)
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
		return nil, ErrImgDecode
	}
	return img, nil
}

func (cnv ConverterPNG) Resize(img *image.Image, width, height uint) ([]byte, error) {
	resizedImg := resize.Resize(width, height, *img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err := png.Encode(buf, resizedImg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type HasherMD5 struct{}

func (h HasherMD5) Gen(raw *[]byte) string {
	hash := md5.New()
	hash.Write(*raw)

	return hex.EncodeToString(hash.Sum(nil))
}

type ResizeService interface {
	ResizeAsync(ctx context.Context, req *ResizeServiceRequest) error
	// ResizeSync(ctx context.Context, raw *[]byte) (image.Image, error)
}
