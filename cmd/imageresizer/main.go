package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/schepelin/imageresizer/pkg/http"
	"github.com/schepelin/imageresizer/pkg/imageservice"
	"github.com/schepelin/imageresizer/pkg/postgres"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"image"
	"image/color"
	"image/png"
	"log"
	"database/sql"
	"os"
)

func createSampleImage() []byte {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	img.Set(2, 1, color.RGBA{255, 0, 0, 255})
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Println("Failed to encode png", err)
	}
	return buf.Bytes()
}

func main() {
	var err error
	logger := log.New(os.Stdout, "", log.LstdFlags)
	const dbConnect string = "postgres://localhost/image_resizer?sslmode=disable"
	db, err := sql.Open("postgres", dbConnect)
	defer db.Close()
	if err != nil {
		logger.Panic("Could not connect to the database")
	}


	ps := postgres.NewPostgresStorage(db)
	h := resizer.HasherMD5{}
	cl := resizer.ClockUTC{}
	cnv := resizer.ConverterPNG{}

	is := imageservice.NewImageService(ps, cl, h, cnv)

	b := createSampleImage()
	ctx := context.TODO()
	_, err = is.Create(ctx, &b)
	if err != nil {
		logger.Panic("Could not write an image to the database")
	}
}

func startServer(addr string, logger *log.Logger) {
	server := http.NewServer(addr, logger)
	server.Start()
}
