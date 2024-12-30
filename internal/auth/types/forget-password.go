package auth_types

type ForgetPasswordReqBody struct {
	Email string `json:"email" validate:"required,email"`
}
