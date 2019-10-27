package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	mux := NewRouter()
	payload := []byte("Lorem ipsum dolor sit amet, consectetur adipisicing elit")

	for _, method := range []string{"GET", "PUT", "POST", "DELETE"} {
		req, err := http.NewRequest(method, "/echo", bytes.NewReader(payload))
		require.NoError(t, err, method)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		resp := rr.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode, method)

		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err, method)
		assert.Equal(t, payload, body, method)

		resp.Body.Close()
	}
}
