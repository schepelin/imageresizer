package postgres

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestImageService_Create(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, mock, err := sqlmock.New()
	defer db.Close()

	is := ImageService{
		Logger: logger,
		DB:     db,
	}

	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})

	buf := new(bytes.Buffer)
	err = png.Encode(buf, sampleImg)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	buf.Bytes()
	hasher := md5.New()
	if _, err := io.Copy(hasher, bytes.NewReader(buf.Bytes())); err != nil {
		t.Error("Something went wrong")
	}
	hash := hex.EncodeToString(hasher.Sum(nil))

	mock.ExpectExec("INSERT INTO images").WithArgs(hash, *buf)

	img, err := is.Create(buf.Bytes())
	if err != nil {
		t.Error("Error while ImageService.Create ", err)
	}

	if img.Hash != hash {
		t.Errorf("Image.Hash %s not equal to expected %s", img.Hash, hash)
	}
	if err != nil {
		t.Error("Error while creating image")
	}
}
