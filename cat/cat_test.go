package cat

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	cat := Cat{Name: "Kum Rai"}

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	cat.ServeHTTP(w, r)

	exp := "Kum Rai"
	if w.Body.String() != exp {
		t.Errorf("expected '%s', got '%s'", exp, w.Body.String())
	}

}
