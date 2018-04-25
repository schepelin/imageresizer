package imageservice

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/schepelin/imageresizer/pkg/resizer"
)

type createRequest struct {
	Raw []byte `json:"raw"`
}

type createResponse struct {
	Id  string `json:"id"`
	Err string `json:"err,omitempty"`
}

func MakeCreateEndpoint(svc resizer.ImageService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		imgObj, err := svc.Create(ctx, &req.Raw)
		if err != nil {
			return createResponse{"", err.Error()}, nil
		}
		return createResponse{imgObj.Id, ""}, nil
	}
}

type Endpoints struct {
	CreateImageEndpoint endpoint.Endpoint
}

func MakeServerEndpoint(svc resizer.ImageService) Endpoints {
	return Endpoints{
		CreateImageEndpoint: MakeCreateEndpoint(svc),
	}
}
