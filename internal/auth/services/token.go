package auth_services

import (
	auth_contracts "github.com/ladmakhi81/realtime-blogs/internal/auth/contracts"
	auth_entities "github.com/ladmakhi81/realtime-blogs/internal/auth/entities"
	users_entities "github.com/ladmakhi81/realtime-blogs/internal/users/entities"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type TokenService struct {
	TokenRepo auth_contracts.TokenRepositoryContract
}

func NewTokenService(tokenRepo auth_contracts.TokenRepositoryContract) TokenService {
	return TokenService{TokenRepo: tokenRepo}
}

func (tokenService TokenService) CreateToken(user *users_entities.User) (*auth_entities.Token, error) {
	if deleteTokenErr := tokenService.TokenRepo.DeleteTokensByUserId(user.ID); deleteTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in deleting remaining token from user",
			"TokenService.CreateToken.DeleteTokenByUserId",
			deleteTokenErr.Error(),
		)
	}
	accessToken, accessTokenErr := pkg_utils.GenerateToken(user.ID, user.Email)
	if accessTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in generating access token",
			"AuthService.Login.GenerateToken",
			accessTokenErr.Error(),
		)
	}
	refreshToken, refreshTokenErr := pkg_utils.GenerateCode(100)
	if refreshTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in generating refresh token",
			"AuthService.Login.GenerateRefreshToken",
			refreshTokenErr.Error(),
		)
	}
	token := auth_entities.NewToken(accessToken, refreshToken, user)
	if createTokenErr := tokenService.TokenRepo.CreateToken(token); createTokenErr != nil {
		return nil, pkg_types.NewServerError(
			"error in saving token on database",
			"AuthService.Login.CreateToken",
			createTokenErr.Error(),
		)
	}
	return token, nil
}
