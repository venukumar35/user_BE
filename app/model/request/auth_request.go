package request

type SendotpRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyAuthOtp struct {
	Email string `json:"email" binding:"required,email"`
	Otp   string `json:"otp" binding:"required"`
}
