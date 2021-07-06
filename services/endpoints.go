package services

import (
	"fmt"
	"go_setup_v1/lib/jwtAuth"
	"go_setup_v1/lib/tkt"
	"go_setup_v1/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type LoginResponse struct {
	Token *string `json:"token"`
}

type Endpoints struct {
	dbConfig *string `json:"dbConfig"`
}

func (o *Endpoints) private(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "private")
}

func (o *Endpoints) signUp(tx *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := SignUpRequest{}
	tkt.CheckErr(tkt.ParseParamOrBody(r, &data))
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(*data.Password), bcrypt.DefaultCost)
	tkt.CheckErr(err)
	user := models.User{
		Name:       *data.Name,
		LastName:   *data.LastName,
		Occupation: *data.Occupation,
		Email:      *data.Email,
		Password:   string(encodedPassword),
	}
	tx.Create(&user)
	token, err := jwtAuth.CreateToken(user.ID)
	tkt.CheckErr(err)
	tkt.JsonResponse(LoginResponse{Token: &token}, w)
}

func (o *Endpoints) signIn(tx *gorm.DB, w http.ResponseWriter, r *http.Request) {
	token, err := jwtAuth.CreateToken(6)
	if err != nil {
		panic("err")
	}
	tkt.JsonResponse(LoginResponse{Token: &token}, w)
}

func (o *Endpoints) Handle() {
	tkt.TransactionalLoggable("/signUp", o.dbConfig, o.signUp)
	tkt.TransactionalLoggable("/signIn", o.dbConfig, o.signIn)
	tkt.AuthenticatedEndpoint("/private", o.private)
}

func JsonResponse(i interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	tkt.JsonEncode(i, w)
}

func NewEndpoints(dbConfig *string) *Endpoints {
	return &Endpoints{dbConfig: dbConfig}
}
