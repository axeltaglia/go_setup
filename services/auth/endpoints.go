package auth

import (
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"html"
	"net/http"
)

type Endpoints struct {
}

func (o *Endpoints) private(w http.ResponseWriter, r *http.Request) {
	tokenAuth, err := jwtAuth.ExtractTokenMetadata(r)
	if err != nil {
		panic("unauthorized")
	}
	userId, err := jwtAuth.FetchAuth(tokenAuth)

	fmt.Fprintf(w, "USER ID: , %d, TOKEN: %s", userId, tokenAuth)
}

func (o *Endpoints) login(w http.ResponseWriter, r *http.Request) {
	token, err := jwtAuth.CreateToken(6)
	if err != nil {
		panic("err")
	}
	fmt.Fprintf(w, "TOKEN: , %q", html.EscapeString(token))
}

func (o *Endpoints) Handle() {
	http.HandleFunc("/login", o.login)
	http.HandleFunc("/private", o.private)
}

func NewEndpoints() *Endpoints {
	return &Endpoints{}
}
