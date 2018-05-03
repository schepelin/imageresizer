package msgqueue

import (
	"context"
	"errors"
)

//go:generate mockgen -source=msgqueue.go -destination ../mocks/mock_msgqueue.go -package mocks

var (
	ErrAlreadyDone = errors.New("job has already processed")
)

type Publisher interface {
	PublishResizeJob(ctx context.Context, jobId uint64) error
}

type Consumer interface {
	ConsumeResizeJob(ctx context.Context, jobId uint64) error
}

type PublisherConsumer interface {
	Publisher
	Consumer
}
