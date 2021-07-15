package services

import (
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
	DbConfig *string `json:"dbConfig"`
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
	data := SignInRequest{}
	tkt.CheckErr(tkt.ParseParamOrBody(r, &data))
	currentUser := models.User{}
	if err := tx.Where("email = ?", *data.Email).First(&currentUser).Error; err != nil {
		panic(ErrorResponse{ErrorMessage: "Invalid email or password"})
	}
	err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(*data.Password))
	if err != nil {
		panic(ErrorResponse{ErrorMessage: "Invalid email or password"})
	}
	token, err := jwtAuth.CreateToken(currentUser.ID)
	if err != nil {
		panic("err")
	}
	tkt.JsonResponse(LoginResponse{Token: &token}, w)
}

func (o *Endpoints) Handle() {
	tkt.TransactionalLoggable("/signUp", o.DbConfig, o.signUp)
	tkt.TransactionalLoggable("/signIn", o.DbConfig, o.signIn)
}

func JsonResponse(i interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	tkt.JsonEncode(i, w)
}

func NewEndpoints(dbConfig *string) *Endpoints {
	return &Endpoints{DbConfig: dbConfig}
}
