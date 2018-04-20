package postgres

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"image/png"
	"io"

	"bytes"
	_ "github.com/lib/pq"
	"github.com/schepelin/imageresizer/pkg/imageresizer"

	"log"
)

type ImageService struct {
	DB     *sql.DB
	Logger *log.Logger
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
	is.DB.Query(`INSERT INTO images(hash, data) VALUES($1, $2)`, hash, r)

	return &imgObj, nil
}

func New(dbConnect string, logger *log.Logger) *ImageService {
	db, err := sql.Open("postgres", dbConnect)
	if err != nil {
		logger.Fatal("could not connect to the database", err)
	}
	return &ImageService{
		DB:     db,
		Logger: logger,
	}
}
