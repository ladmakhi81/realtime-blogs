package users_types

type ChangePasswordReqBody struct {
	NewPassword string `json:"newPassword" validate:"required,min=8"`
}
