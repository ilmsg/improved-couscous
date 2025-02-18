package cat

import "net/http"

type Cat struct {
	Name string
}

func (c *Cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(c.Name))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	cat := Cat{"Nina"}
	cat.ServeHTTP(w, r)
}
