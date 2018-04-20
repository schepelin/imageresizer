package postgres

import (
	"crypto/md5"
	"encoding/hex"
	"image/png"
	"io"

	"bytes"
	"github.com/schepelin/imageresizer/pkg/imageresizer"
)

type ImageService struct {
	//DB *sql.DB
}

func (is *ImageService) Create(r *bytes.Buffer) (*imageresizer.Image, error) {
	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}

	hash := md5.New()
	if _, err := io.Copy(hash, r); err != nil {
		return nil, err
	}
	imgObj := imageresizer.Image{
		Hash:  hex.EncodeToString(hash.Sum(nil)),
		Image: img,
	}
	// TODO: Write to database

	return &imgObj, nil
}

func New(dbConnect string) *ImageService {
	return &ImageService{}
}
