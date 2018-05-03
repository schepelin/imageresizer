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
	Publisher msgqueue.Publisher
}

func NewResizeService(
	s storage.ResizeStorage, cnv resizer.Converter, pub msgqueue.Publisher) *ResizeService {
	return &ResizeService{s, cnv, pub}
}

func (rs *ResizeService) ResizeAsync(ctx context.Context, req *resizer.ResizeServiceRequest) error {
	var err error

	err = rs.Publisher.PublishResizeJob(ctx, req.JobId)
	//img, err := rs.Converter.Transform(&req.RawImg)
	//if err != nil {
	//	return err
	//}
	//resizeRaw, err := rs.Converter.Resize(&img, req.Width, req.Height)
	//if err != nil {
	//	return err
	//}
	//err = rs.Storage.WriteResizeJobResult(ctx, &storage.ResizeResultRequest{req.JobId, resizeRaw})
	return err
}
