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
	//const address string = "0.0.0.0:8081"

	const dbConnect string = "postgres://schepelin:topsecret@localhost/image_resizer?sslmode=disable"

	// logger := log.New(os.Stdout, "", log.LstdFlags)

	ps, err := postgres.NewPostgresStorage(dbConnect)
	if err != nil {
		fmt.Println("Can not connect to the database", dbConnect)
	}
	h := resizer.HasherMD5{}
	cl := resizer.ClockUTC{}
	cnv := resizer.ConverterPNG{}

	is := imageservice.NewImageService(ps, cl, h, cnv)

	b := createSampleImage()
	ctx := context.TODO()
	imgObj, err := is.Create(ctx, &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("IMG****", imgObj)

}

func startServer(addr string, logger *log.Logger) {
	server := http.NewServer(addr, logger)
	server.Start()
}
