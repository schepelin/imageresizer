package resizesvc

import (
	"context"
	"github.com/schepelin/imageresizer/pkg/msgqueue"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"github.com/schepelin/imageresizer/pkg/storage"
)

type ResizeService struct {
	Storage   storage.ResizeStorage
	Converter resizer.Converter
	PubSub    msgqueue.PublisherConsumer
}

func NewResizeService(
	s storage.ResizeStorage, cnv resizer.Converter, pubsub msgqueue.PublisherConsumer) *ResizeService {
	return &ResizeService{s, cnv, pubsub}
}

func (rs *ResizeService) SendResizeJob(ctx context.Context, req *resizer.ResizeServiceRequest) error {
	var err error

	err = rs.PubSub.PublishResizeJob(ctx, req.JobId)
	return err
}

func (rs *ResizeService) RunResizeWorker(ctx context.Context, ch chan uint64) error {
	go rs.PubSub.ConsumeResizeJobs(ctx, ch)

	for jobId := range ch {
		resp, err := rs.Storage.GetResizeJobForUpdate(ctx, jobId)
		if err != nil {
			return err
		}
		img, err := rs.Converter.Transform(&resp.RawImg)
		if err != nil {
			return err
		}
		raw, err := rs.Converter.Resize(&img, resp.Width, resp.Height)

		if err != nil {
			return err
		}
		rs.Storage.WriteResizeJobResult(ctx, &storage.ResizeResultRequest{
			jobId,
			raw,
		})

	}

	return nil
}
