package imageservice

import (
	"context"
	"net/http"
	"encoding/json"
)

func DecodeCreateRequest( ctx context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
