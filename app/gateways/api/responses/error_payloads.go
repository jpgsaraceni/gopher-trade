package responses

type ErrorPayload struct {
	Error string `json:"error,omitempty" example:"Message for some error"`
}

var (
	// generic errors
	ErrInternalServerError = newErrorPayload("Internal server error.")
	ErrMalformedBody       = newErrorPayload("Malformed request body.")
	ErrMissingFields       = newErrorPayload("Missing required fields.")

	// exchange errors
	ErrInvalidRate    = newErrorPayload("Invalid rate. Must be an integer or point separated decimal number.")
	ErrConflictFromTo = newErrorPayload("Rate for from-to currency pair already exists.")
)

func newErrorPayload(msg string) ErrorPayload {
	return ErrorPayload{
		Error: msg,
	}
}
