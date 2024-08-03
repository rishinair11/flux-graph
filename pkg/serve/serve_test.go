package serve

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleGraphSVG(t *testing.T) {
	// test server
	req := httptest.NewRequest(http.MethodGet, "/graph.html", nil)
	w := httptest.NewRecorder()

	handleGraphSVG(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// response body should match the static file
	expectedBody, err := os.ReadFile(filepath.Join("static", "graph.html"))
	assert.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, string(expectedBody), string(body))
}
