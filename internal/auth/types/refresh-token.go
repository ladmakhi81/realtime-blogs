package auth_types

type RefreshTokenReqBody struct {
	AccessToken  string `json:"accessToken" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}
