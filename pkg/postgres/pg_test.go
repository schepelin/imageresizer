package postgres

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"v/gopkg.in/DATA-DOG/go-sqlmock.v1@v1.3.0-gopkgin-v1.3.0"
	"image"
	"image/color"
	"bytes"
	"image/png"
	"fmt"
	"github.com/schepelin/imageresizer/pkg/storage"
	"time"
	"context"
)

func getRawImageSample() []byte {
	sampleImg := image.NewRGBA(image.Rect(0, 0, 10, 10))
	sampleImg.Set(1, 1, color.RGBA{255, 0, 0, 255})

	buf := new(bytes.Buffer)
	err := png.Encode(buf, sampleImg)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}

	return buf.Bytes()
}

func TestNewPostgresStorage(t *testing.T) {
	_, err := NewPostgresStorage("postgres://test:test@localhost/test?sslmode=disable")
	assert.NoError(t, err)
	// TODO: How to check ps
}


func TestPostgresStorage_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
		t.Error()
	}
	defer db.Close()
	ps := PostgresStorage{db}
	ctx := context.TODO()
	imgModel := storage.ImageModel{
		Id: "100500",
		Raw: getRawImageSample(),
		CreatedAt: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	mock.ExpectExec("INSERT INTO images").WithArgs(imgModel.Id, string(imgModel.Raw), imgModel.CreatedAt)
	ps.Create(ctx, &imgModel)

	assert.NoError(t, mock.ExpectationsWereMet())
}
