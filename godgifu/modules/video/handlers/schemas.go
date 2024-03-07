package handlers

// signupSchema models the request data to /auth/signup
type signupSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
