package services

type ErrorResponse struct {
	ErrorMessage string      `json:"errorMessage"`
	Error        interface{} `json:"error"`
}

type SignInRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type SignUpRequest struct {
	Name       *string `json:"name"`
	LastName   *string `json:"lastName"`
	Occupation *string `json:"occupation"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
}

type CreateCategoryRequest struct {
	Name *string `json:"name"`
}
