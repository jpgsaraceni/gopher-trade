package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Payload any
	Writer  http.ResponseWriter
	Error   error
	Status  int
}

func BadRequest(w http.ResponseWriter, payload ErrorPayload, err error) {
	errResponse(w, payload, err, http.StatusBadRequest).sendJSON()
}

func NotFound(w http.ResponseWriter, payload ErrorPayload, err error) {
	errResponse(w, payload, err, http.StatusNotFound).sendJSON()
}

func Conflict(w http.ResponseWriter, payload ErrorPayload, err error) {
	errResponse(w, payload, err, http.StatusConflict).sendJSON()
}

func InternalServerError(w http.ResponseWriter, err error) {
	errResponse(w, ErrInternalServerError, err, http.StatusInternalServerError).sendJSON()
}

func BadGateway(w http.ResponseWriter, payload ErrorPayload, err error) {
	errResponse(w, payload, err, http.StatusBadGateway).sendJSON()
}

func Created(w http.ResponseWriter, payload any) {
	successResponse(w, payload, http.StatusCreated).sendJSON()
}

func OK(w http.ResponseWriter, payload any) {
	successResponse(w, payload, http.StatusOK).sendJSON()
}

func successResponse(w http.ResponseWriter, payload any, status int) Response {
	return Response{
		Writer:  w,
		Status:  status,
		Payload: payload,
	}
}

func errResponse(w http.ResponseWriter, payload ErrorPayload, err error, status int) Response {
	return Response{
		Writer:  w,
		Error:   err,
		Status:  status,
		Payload: payload,
	}
}

func (r Response) sendJSON() {
	r.Writer.Header().Set("Content-Type", "application/json")
	r.Writer.WriteHeader(r.Status)
	if r.Error != nil {
		log.Println(r.Error.Error())
	}
	if err := json.NewEncoder(r.Writer).Encode(r.Payload); err != nil {
		log.Printf("failed to encode http response: %s", err)
	}
}
