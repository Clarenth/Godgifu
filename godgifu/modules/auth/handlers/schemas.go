package handlers

// signupSchema models the request data to /auth/signup
type signupSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=16,lte=512"`
}
