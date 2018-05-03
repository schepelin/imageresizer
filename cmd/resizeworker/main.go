package main

import (
	"context"
	"database/sql"
	"github.com/schepelin/imageresizer/pkg/postgres"
	"github.com/schepelin/imageresizer/pkg/rabbitmq"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"github.com/schepelin/imageresizer/pkg/resizesvc"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

func main() {
	const dbConnect string = "postgres://localhost/image_resizer?sslmode=disable"
	const mqConnect string = "amqp://guest:guest@localhost:5672/"

	logger := log.New(os.Stdout, "", log.LstdFlags)

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
	cnv := resizer.ConverterPNG{}
	rSvc := resizesvc.NewResizeService(ps, cnv, pubSub)

	ch := make(chan uint64)
	rSvc.RunResizeWorker(context.Background(), ch)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	close(ch)

}
