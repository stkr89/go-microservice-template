package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/shopr-org/grpc-service-template/common"
	"github.com/shopr-org/grpc-service-template/endpoints"
	"github.com/shopr-org/grpc-service-template/pb"
	"net/http"
)

type errorWrapper struct {
	Error string `json:"error"`
}

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/api/v1/register", httptransport.NewServer(
		endpoints.Add,
		decodeHTTPAddRequest,
		encodeHTTPGenericResponse,
	))

	return m
}

func err2code(err error) int {
	switch err {
	case common.InvalidRequestBody:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.MathRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}
