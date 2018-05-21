package rabbitmq

import (
	"context"
	"encoding/binary"
	"github.com/streadway/amqp"
)

type Config struct {
	Queue    string
	Exchange string
}

type PubSub struct {
	Channel *amqp.Channel
	Cfg     *Config
}

func NewPubSub(ch *amqp.Channel, cfg *Config) *PubSub {
	return &PubSub{ch, cfg}
}

func (p *PubSub) PublishResizeJob(ctx context.Context, jobId uint64) error {
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

func (p *PubSub) ConsumeResizeJobs(ctx context.Context, ch chan<- uint64) error {
	msgs, err := p.Channel.Consume(
		p.Cfg.Queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		jobId := binary.LittleEndian.Uint64(msg.Body)
		ch <- jobId
	}
	return nil
}
