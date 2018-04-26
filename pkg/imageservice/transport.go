package imageservice

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"image/png"
	"net/http"
)

var (
	OperationError    = errors.New("service respond with error")
	InconsistentRoute = errors.New("required route params not found in URI")
)

func DecodeCreateImageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request createImageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetImageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return nil, InconsistentRoute
	}
	return getImageRequest{
		Id: id,
	}, nil
}

func DecodeDeleteImageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, InconsistentRoute
	}
	return deleteImageRequest{id}, nil
}

func EncodeImageResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "image/png")
	resp := response.(getImageResponse)
	if resp.Err != nil {
		return resp.Err
	}
	err := png.Encode(w, resp.Img)
	return err
}

func MakeHTTPHandler(is resizer.ImageService) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoint(is)
	r.Methods("POST").Path("/images").Handler(httptransport.NewServer(
		e.CreateImageEndpoint,
		DecodeCreateImageRequest,
		httptransport.EncodeJSONResponse,
	))
	r.Methods("GET").Path("/images/{id}").Handler(httptransport.NewServer(
		e.GetImageEndpoint,
		DecodeGetImageRequest,
		EncodeImageResponse,
	))
	r.Methods("DELETE").Path("/images/{id}").Handler(httptransport.NewServer(
		e.DeleteImageEndpoint,
		DecodeDeleteImageRequest,
		httptransport.EncodeJSONResponse,
	))
	return r
}
