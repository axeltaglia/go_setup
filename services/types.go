package services

type SignUpRequest struct {
	Name       *string `json:"name"`
	LastName   *string `json:"lastName"`
	Occupation *string `json:"occupation"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
}
