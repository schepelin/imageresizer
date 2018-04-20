package postgres

import (
	"bytes"
	"image"
	"image/png"
	"testing"
)

func TestImageService_Create(t *testing.T) {
	is := ImageService{}
	sampleImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	var b bytes.Buffer
	png.Encode(&b, sampleImg)

	_, err := is.Create(&b)
	if err != nil {
		t.Error("Error while creating image")
	}
}
