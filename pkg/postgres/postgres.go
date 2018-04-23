package postgres

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"image/png"
	"log"

	_ "github.com/lib/pq"
	"github.com/schepelin/imageresizer/pkg/imageresizer"
)

type ImageService struct {
	DB     *sql.DB
	Logger *log.Logger
}

// TODO: Do I need *[]byte here
func (is *ImageService) Create(b []byte) (*imageresizer.Image, error) {
	buf := bytes.NewBuffer(b)
	img, err := png.Decode(buf)
	if err != nil {
		is.Logger.Panic("Could not decode png ", err)
		return nil, err
	}
	hash := md5.New()
	if _, err := hash.Write(b); err != nil {
		return nil, err
	}
	imgObj := imageresizer.Image{
		Hash:  hex.EncodeToString(hash.Sum(nil)),
		Image: img,
	}
	is.DB.Query(`INSERT INTO images(hash, data) VALUES($1, $2)`, hash, b)

	return &imgObj, nil
}

func New(dbConnect string, logger *log.Logger) *ImageService {
	db, err := sql.Open("postgres", dbConnect)
	if err != nil {
		logger.Fatal("could not connect to the database ", err)
	}
	return &ImageService{
		DB:     db,
		Logger: logger,
	}
}
