package server

import (
	"encoding/gob"
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	Error string `json:"error"`
}

type resp struct {
	w http.ResponseWriter
}

func (r *resp) Code(code int) *resp {
	r.w.WriteHeader(code)

	return r
}

func (r *resp) Error(err error) *resp {
	ResponseError(r.w, err)

	return r
}

func (r *resp) JsonBody(data any) *resp {
	r.w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(r.w).Encode(data)

	return r
}

func (r *resp) Body(data any) *resp {
	enc := gob.NewEncoder(r.w)
	enc.Encode(data)

	return r
}

func Response(w http.ResponseWriter) *resp {
	return &resp{w}
}

func ResponseError(w http.ResponseWriter, err error) {
	if err != nil {
		r := ErrResponse{Error: err.Error()}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(r)
	}
}

func ResponseCode(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

func ResponseData(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
