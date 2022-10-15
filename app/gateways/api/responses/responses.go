package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Payload interface{}
	Writer  http.ResponseWriter
	Error   error
	Status  int
}

func BadRequest(w http.ResponseWriter, payload ErrorPayload, err error) {
	r := Response{
		Writer:  w,
		Status:  http.StatusBadRequest,
		Error:   err,
		Payload: payload,
	}

	r.sendJSON()
}

func InternalServerError(w http.ResponseWriter, err error) {
	r := Response{
		Writer:  w,
		Status:  http.StatusInternalServerError,
		Error:   err,
		Payload: ErrInternalServerError,
	}

	r.sendJSON()
}

func Created(w http.ResponseWriter, payload interface{}) {
	r := Response{
		Writer:  w,
		Status:  http.StatusCreated,
		Payload: payload,
	}

	r.sendJSON()
}

func (r Response) sendJSON() {
	r.Writer.Header().Set("Content-Type", "application/json")
	r.Writer.WriteHeader(r.Status)
	if r.Error != nil {
		log.Println(r.Error)
	}
	if err := json.NewEncoder(r.Writer).Encode(r.Payload); err != nil {
		log.Printf("failed to encode http response: %s", err)
	}
}
