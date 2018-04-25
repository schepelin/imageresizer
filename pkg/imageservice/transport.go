package imageservice

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/schepelin/imageresizer/pkg/resizer"
	"net/http"
)

func DecodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeHTTPHandler(is resizer.ImageService) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoint(is)
	r.Methods("POST").Path("/images").Handler(httptransport.NewServer(
		e.CreateImageEndpoint,
		DecodeCreateRequest,
		EncodeResponse,
	))
	return r
}
