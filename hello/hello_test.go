package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	// if w.Code != http.StatusOK { }
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", resp.StatusCode)
	}

	exp := "hello"
	if w.Body.String() != exp {
		t.Errorf("expected '%s', got '%s'", exp, w.Body.String())
	}

}
