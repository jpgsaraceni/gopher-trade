package exchanges

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var testContext = context.Background()

func newTestPostRequest(t *testing.T, target string, body CreateExchangeRequest) *http.Request {
	t.Helper()

	reqPayload, err := json.Marshal(body)
	assert.NoError(t, err)
	reader := bytes.NewReader(reqPayload)
	request, err := http.NewRequestWithContext(testContext, http.MethodPost, target, reader)
	assert.NoError(t, err)

	return request
}

func newTestPostResponse(h http.HandlerFunc, req *http.Request, target string) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	router.Post(target, h)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	return res
}

func newTestGetRequest(t *testing.T, target string, params map[string]string) *http.Request {
	t.Helper()

	if len(params) > 0 {
		separater := "?"
		for key, value := range params {
			target += separater + key + "=" + value
			separater = "&"
		}
	}
	request, err := http.NewRequestWithContext(testContext, http.MethodGet, target, nil)
	assert.NoError(t, err)

	return request
}

func newTestGetResponse(h http.HandlerFunc, req *http.Request, target string) *httptest.ResponseRecorder {
	router := chi.NewRouter()
	router.Get(target, h)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	return res
}

func assertResponse(t *testing.T, wantStatus int, wantBody json.RawMessage, res *httptest.ResponseRecorder) {
	t.Helper()

	assert.Equal(t, wantStatus, res.Code)
	gotBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, string(wantBody), string(gotBody), "got payload: %s", string(gotBody))
}
