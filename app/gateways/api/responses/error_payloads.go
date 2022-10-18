package responses

type ErrorPayload struct {
	Error string `json:"error,omitempty" example:"Message for some error"`
}

var (
	// generic errors
	ErrInternalServerError = newErrorPayload("Internal server error.")
	ErrMalformedBody       = newErrorPayload("Malformed request body.")
	ErrMissingFields       = newErrorPayload("Missing required fields.")
	ErrMissingParams       = newErrorPayload("Missing required params.")

	// currency errors
	ErrInvalidRate          = newErrorPayload("Invalid rate. Must be an integer or point separated decimal number.")
	ErrInvalidAmount        = newErrorPayload("Invalid amount. Must be an integer or point separated decimal number.")
	ErrCurrenciesNotFound   = newErrorPayload("Currency pair not found.")
	ErrCurrencyCodeNotFound = newErrorPayload("Currency code not found.")
	ErrCurrencyAPI          = newErrorPayload("External API not available.")
	ErrIsDefaultRate        = newErrorPayload("Code belongs to a default rate.")
)

func newErrorPayload(msg string) ErrorPayload {
	return ErrorPayload{
		Error: msg,
	}
}
