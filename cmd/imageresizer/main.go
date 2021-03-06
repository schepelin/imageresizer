package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/schepelin/imageresizer/pkg/imageservice"
	"github.com/schepelin/imageresizer/pkg/postgres"
	"github.com/schepelin/imageresizer/pkg/rabbitmq"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"github.com/schepelin/imageresizer/pkg/resizesvc"
	"github.com/streadway/amqp"
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
	var err error
	logger := log.New(os.Stdout, "", log.LstdFlags)
	const dbConnect string = "postgres://localhost/image_resizer?sslmode=disable"
	const mqConnect string = "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(mqConnect)
	if err != nil {
		logger.Panic("Could not connect to the amqp server")
	}
	defer conn.Close()
	mqChannel, err := conn.Channel()
	if err != nil {
		logger.Panic("Could not create the channel")
	}
	mqCfg := rabbitmq.Config{
		Queue:    "test",
		Exchange: "",
	}
	_, err = mqChannel.QueueDeclare(
		mqCfg.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Panic("Could not declare a queue")
	}

	pubSub := rabbitmq.NewPubSub(mqChannel, &mqCfg)

	db, err := sql.Open("postgres", dbConnect)
	if err != nil {
		logger.Panic("Could not connect to the database")
	}
	defer db.Close()

	ps := postgres.NewPostgresStorage(db)
	h := resizer.HasherMD5{}
	cl := resizer.ClockUTC{}
	cnv := resizer.ConverterPNG{}
	rSvc := resizesvc.NewResizeService(ps, cnv, pubSub)
	is := imageservice.NewImageService(ps, cl, h, cnv, rSvc)

	handler := imageservice.MakeHTTPHandler(is)
	logger.Fatal(http.ListenAndServe(":8080", handler))
}
