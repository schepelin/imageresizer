package imageservice

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"image"
	"time"
)

type createImageRequest struct {
	Raw []byte `json:"raw"`
}

type createImageResponse struct {
	Id  string `json:"id"`
	Err string `json:"err,omitempty"`
}

type getImageRequest struct {
	Id string `json:"id"`
}

type getImageResponse struct {
	Img image.Image
	Err error
}

type deleteImageRequest struct {
	Id string `json:"id"`
}

type deleteImageResponse struct {
	Err string `json:"err,omitempty"`
}

type createResizeRequest struct {
	ImgId  string
	Width  uint `json:"width"`
	Height uint `json:"height"`
}

type createResizeResponse struct {
	Err       string    `json:"err,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func MakeCreateImageEndpoint(svc resizer.ImageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createImageRequest)
		imgObj, err := svc.Create(ctx, &req.Raw)
		if err != nil {
			return createImageResponse{"", err.Error()}, nil
		}
		return createImageResponse{imgObj.Id, ""}, nil
	}
}

func MakeGetImageEndpoint(svc resizer.ImageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getImageRequest)
		imgObj, err := svc.Read(ctx, req.Id)
		if err != nil {
			return getImageResponse{nil, err}, nil
		}
		return getImageResponse{imgObj.Image, nil}, nil
	}
}

func MakeDeleteImageEndpoint(svc resizer.ImageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteImageRequest)
		err := svc.Delete(ctx, req.Id)
		if err != nil {
			return deleteImageResponse{err.Error()}, nil
		}
		return deleteImageResponse{""}, nil
	}
}

func MakeScheduleResizeJobEndpoint(svc resizer.ImageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createResizeRequest)
		job, err := svc.ScheduleResizeJob(ctx, req.ImgId, req.Width, req.Height)
		if err != nil {
			return createResizeResponse{err.Error(), resizer.StatusFailed, time.Time{}}, nil
		}
		return createResizeResponse{"", job.Status, job.CreatedAt}, nil
	}
}

type Endpoints struct {
	CreateImageEndpoint       endpoint.Endpoint
	GetImageEndpoint          endpoint.Endpoint
	DeleteImageEndpoint       endpoint.Endpoint
	ScheduleResizeJobEndpoint endpoint.Endpoint
}

func MakeServerEndpoint(svc resizer.ImageService) Endpoints {
	return Endpoints{
		CreateImageEndpoint:       MakeCreateImageEndpoint(svc),
		GetImageEndpoint:          MakeGetImageEndpoint(svc),
		DeleteImageEndpoint:       MakeDeleteImageEndpoint(svc),
		ScheduleResizeJobEndpoint: MakeScheduleResizeJobEndpoint(svc),
	}
}
