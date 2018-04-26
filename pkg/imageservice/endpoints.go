package imageservice

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"image"
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

type Endpoints struct {
	CreateImageEndpoint endpoint.Endpoint
	GetImageEndpoint    endpoint.Endpoint
	DeleteImageEndpoint endpoint.Endpoint
}

func MakeServerEndpoint(svc resizer.ImageService) Endpoints {
	return Endpoints{
		CreateImageEndpoint: MakeCreateImageEndpoint(svc),
		GetImageEndpoint:    MakeGetImageEndpoint(svc),
		DeleteImageEndpoint: MakeDeleteImageEndpoint(svc),
	}
}
