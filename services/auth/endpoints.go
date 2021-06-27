package auth

import (
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"go_setup_v1/lib/tkt"
	"gorm.io/gorm"
	"net/http"
)

type LoginResponse struct {
	Token *string `json:"token"`
}

type Endpoints struct {
}

func (o *Endpoints) private(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "private")
}

func (o *Endpoints) login(w http.ResponseWriter, r *http.Request) {
	token, err := jwtAuth.CreateToken(6)
	if err != nil {
		panic("err")
	}
	tkt.JsonResponse(LoginResponse{Token: &token}, w)
}

func (o *Endpoints) Handle() {
	tkt.TransactionalLoggable("/login", o.login)
	tkt.AuthenticatedEndpoint("/private", o.private)
}

func JsonResponse(i interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	tkt.JsonEncode(i, w)
}

func NewEndpoints(db *gorm.DB) *Endpoints {
	return &Endpoints{}
}
