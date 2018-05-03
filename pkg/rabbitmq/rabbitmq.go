package rabbitmq

import (
	"context"
	"encoding/binary"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"github.com/schepelin/imageresizer/pkg/storage"
	"github.com/streadway/amqp"
)

type Config struct {
	Queue    string
	Exchange string
}

type Publisher struct {
	Channel *amqp.Channel
	Cfg     *Config
}

type Consumer struct {
	Channel   *amqp.Channel
	Storage   storage.ResizeStorage
	Converter resizer.Converter
	Cfg       *Config
}

func NewPublisher(ch *amqp.Channel, cfg *Config) *Publisher {
	return &Publisher{ch, cfg}
}

func (p *Publisher) PublishResizeJob(ctx context.Context, jobId uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, jobId)
	err := p.Channel.Publish(
		p.Cfg.Exchange,
		p.Cfg.Queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	return err
}
