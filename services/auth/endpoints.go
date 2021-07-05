package auth

import (
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"go_setup_v1/lib/tkt"
	"go_setup_v1/models"
	"gorm.io/gorm"
	"net/http"
)

type LoginResponse struct {
	Token *string `json:"token"`
}

type Endpoints struct {
	db *gorm.DB `json:"db"`
}

func (o *Endpoints) private(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "private")
}

func (o *Endpoints) signUp(tx *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token, err := jwtAuth.CreateToken(6)
	if err != nil {
		panic("err")
	}
	tkt.JsonResponse(LoginResponse{Token: &token}, w)
}

func (o *Endpoints) signIn(tx *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tx.Create(&models.Category{
		Name: "hola mundo 2",
	})
	token, err := jwtAuth.CreateToken(6)
	if err != nil {
		panic("err")
	}
	tkt.JsonResponse(LoginResponse{Token: &token}, w)
}

func (o *Endpoints) Handle() {
	tkt.TransactionalLoggable("/signUp", o.db, o.signUp)
	tkt.TransactionalLoggable("/signIn", o.db, o.signIn)
	tkt.AuthenticatedEndpoint("/private", o.private)
}

func JsonResponse(i interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	tkt.JsonEncode(i, w)
}

func NewEndpoints(db *gorm.DB) *Endpoints {
	return &Endpoints{db: db}
}
