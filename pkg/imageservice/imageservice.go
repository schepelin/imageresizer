package imageservice

import (
	"context"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"github.com/schepelin/imageresizer/pkg/storage"
)

type ImageService struct {
	Storage   storage.Storage
	Clock     resizer.Clocker
	Hash      resizer.Hasher
	Converter resizer.Converter
	ResizeSvc resizer.ResizeService
}

func NewImageService(s storage.Storage, cl resizer.Clocker,
	h resizer.Hasher, c resizer.Converter, rs resizer.ResizeService) *ImageService {

	return &ImageService{
		Storage:   s,
		Clock:     cl,
		Hash:      h,
		Converter: c,
		ResizeSvc: rs,
	}
}

func (is *ImageService) Create(ctx context.Context, raw *[]byte) (*resizer.Image, error) {
	imgModel := storage.ImageModel{
		Id:        is.Hash.Gen(raw),
		Raw:       *raw,
		CreatedAt: is.Clock.Now(),
	}
	err := is.Storage.Create(ctx, &imgModel)
	if err != nil {
		return nil, err
	}
	img, err := is.Converter.Transform(raw)
	if err != nil {
		return nil, err
	}
	return &resizer.Image{
		Id:        imgModel.Id,
		Image:     img,
		CreatedAt: imgModel.CreatedAt,
	}, nil

}

func (is *ImageService) Read(ctx context.Context, id string) (*resizer.Image, error) {
	imgModel, err := is.Storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	img, err := is.Converter.Transform(&imgModel.Raw)
	if err != nil {
		return nil, err
	}
	return &resizer.Image{
		Id:        imgModel.Id,
		Image:     img,
		CreatedAt: imgModel.CreatedAt,
	}, nil

}

func (is *ImageService) Delete(ctx context.Context, id string) error {
	err := is.Storage.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (is *ImageService) ScheduleResizeJob(ctx context.Context, imgId string, width, height uint) (*resizer.ResizeJob, error) {
	var err error
	req := storage.ResizeJobRequest{
		ImgId:  imgId,
		Width:  width,
		Height: height,
	}
	resp, err := is.Storage.CreateResizeJob(ctx, &req)
	if err != nil {
		return nil, err
	}

	is.ResizeSvc.ResizeAsync(ctx, &resizer.ResizeServiceRequest{
		JobId:  resp.Id,
		RawImg: resp.RawImg,
		Width:  req.Width,
		Height: req.Height,
	})
	return &resizer.ResizeJob{
		Id:        resp.Id,
		ImageId:   imgId,
		Status:    resp.Status,
		Image:     nil,
		CreatedAt: resp.CreatedAt,
	}, nil

}
