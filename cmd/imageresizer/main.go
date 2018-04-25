package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/schepelin/imageresizer/pkg/imageservice"
	"github.com/schepelin/imageresizer/pkg/postgres"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
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

	rawImage := createSampleImage()
	fmt.Println("Image:", rawImage)

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

	handler := imageservice.MakeHTTPHandler(is)
	logger.Fatal(http.ListenAndServe(":8080", handler))

}
