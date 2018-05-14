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
	"strconv"
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

func DecodeGetResizeJobRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, InconsistentRoute
	}
	jobId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, InconsistentRoute
	}

	return getResizeJobRequest{jobId}, nil
}

func EncodeGetResizeJobResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "image/png")
	resp := response.(getResizeJobResponse)
	if resp.Err != nil {
		return resp.Err
	}
	err := png.Encode(w, resp.Img)
	return err
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

func DecodeScheduleResizeJobRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	imgId, ok := vars["id"]
	if !ok {
		return nil, InconsistentRoute
	}
	var req createResizeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	req.ImgId = imgId
	return req, err
}

func MakeHTTPHandler(is resizer.ImageService) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoint(is)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/images").Handler(httptransport.NewServer(
		e.CreateImageEndpoint,
		DecodeCreateImageRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
	r.Methods("GET").Path("/images/{id}").Handler(httptransport.NewServer(
		e.GetImageEndpoint,
		DecodeGetImageRequest,
		EncodeImageResponse,
		options...,
	))
	r.Methods("DELETE").Path("/images/{id}").Handler(httptransport.NewServer(
		e.DeleteImageEndpoint,
		DecodeDeleteImageRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
	r.Methods("POST").Path("/images/{id}/jobs").Handler(httptransport.NewServer(
		e.ScheduleResizeJobEndpoint,
		DecodeScheduleResizeJobRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))
	r.Methods("GET").Path("/jobs/{id}").Handler(httptransport.NewServer(
		e.GetResizeJobEndpoint,
		DecodeGetResizeJobRequest,
		EncodeGetResizeJobResponse,
		options...,
	))
	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case resizer.ErrNoImage:
		return http.StatusNotFound
	case InconsistentRoute:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
