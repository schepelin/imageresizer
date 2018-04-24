package postgres

import (
	"bytes"
	"context"
	"fmt"
	"github.com/schepelin/imageresizer/pkg/storage"
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"image/png"
	"testing"
	"time"
	"v/gopkg.in/DATA-DOG/go-sqlmock.v1@v1.3.0-gopkgin-v1.3.0"
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
		Id:        "100500",
		Raw:       getRawImageSample(),
		CreatedAt: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	mock.ExpectExec("INSERT INTO images").WithArgs(imgModel.Id, string(imgModel.Raw), imgModel.CreatedAt)
	ps.Create(ctx, &imgModel)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresStorage_Get(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	ps := PostgresStorage{db}
	ctx := context.TODO()
	imgId := "42"
	expectedRaw := []byte{10, 42, 15}
	expectedCreatedAt := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	rows := sqlmock.NewRows([]string{"raw", "created_at"}).AddRow(
		string(expectedRaw),
		expectedCreatedAt,
	)
	mock.ExpectQuery(
		"SELECT raw, created_at FROM images WHERE id=?",
	).WithArgs(imgId).WillReturnRows(rows)

	imgModel, err := ps.Get(ctx, imgId)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, imgModel.Id, imgId)
	assert.Equal(t, imgModel.Raw, expectedRaw)
	assert.Equal(t, imgModel.CreatedAt, expectedCreatedAt)
}

func TestPostgresStorage_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	ps := PostgresStorage{db}
	ctx := context.TODO()
	imgId := "42"

	mock.ExpectExec(
		"DELETE FROM images WHERE id=?",
	).WithArgs(imgId).WillReturnResult(sqlmock.NewResult(100500, 1))

	ps.Delete(ctx, imgId)
	assert.NoError(t, mock.ExpectationsWereMet())
}
