package imageresizer

import (
	"bytes"
	"image"
)

type Image struct {
	Hash  string
	Image image.Image
	// TODO: Add CreatedAt time.Time.Date,
}

type ImageService interface {
	Create(r *bytes.Reader) (*Image, error)
	//Read(imgId ImageId) (*Image, error)
	//Delete(imgId ImageId) bool
}
