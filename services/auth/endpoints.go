package auth

import (
	"fmt"
	"html"
	"net/http"
)

type Endpoints struct {
}

func (o *Endpoints) test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (o *Endpoints) Handle() {
	http.HandleFunc("/test", o.test)
}

func NewEndpoints() *Endpoints {
	return &Endpoints{}
}
