package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReq(t *testing.T) {
	mux := NewRouter()

	for _, method := range []string{"GET", "PUT", "POST", "DELETE"} {
		req, err := http.NewRequest(method, "/req", nil)
		require.NoError(t, err, method)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		resp := rr.Result()

		if method != "GET" {
			assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode, method)
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		resp.Body.Close()

		assert.Equal(t, "GET /req HTTP/1.1\r\nRemoteaddr: \r\n\r\n", string(body), method)
	}
}
