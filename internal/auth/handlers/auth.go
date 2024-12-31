package auth_handlers

import (
	"encoding/json"
	"net/http"

	auth_contracts "github.com/ladmakhi81/realtime-blogs/internal/auth/contracts"
	auth_types "github.com/ladmakhi81/realtime-blogs/internal/auth/types"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type AuthHandler struct {
	AuthService auth_contracts.AuthServiceContract
}

func NewAuthHandler(authService auth_contracts.AuthServiceContract) AuthHandler {
	return AuthHandler{AuthService: authService}
}

func (authHandler AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(auth_types.LoginReqBody)
	if decodeErr := json.NewDecoder(r.Body).Decode(reqBody); decodeErr != nil {
		return pkg_types.NewServerError("unable to parse request body", "AuthHandler.Login", decodeErr.Error())
	}
	if validationErr := pkg_utils.ValidateHttpReqBody(reqBody); validationErr != nil {
		return pkg_types.NewClientValidationError(
			validationErr,
		)
	}
	res, err := authHandler.AuthService.Login(*reqBody)
	if err != nil {
		return err
	}
	pkg_utils.JsonResponse(w, http.StatusOK, res)
	return nil
}

func (authHandler AuthHandler) Signup(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(auth_types.SignupReqBody)
	if decodeErr := json.NewDecoder(r.Body).Decode(reqBody); decodeErr != nil {
		return pkg_types.NewServerError("unable to parse request body", "AuthHandler.Signup", decodeErr.Error())
	}
	if validationErr := pkg_utils.ValidateHttpReqBody(reqBody); validationErr != nil {
		return pkg_types.NewClientValidationError(
			validationErr,
		)
	}
	res, err := authHandler.AuthService.Signup(*reqBody)
	if err != nil {
		return err
	}
	pkg_utils.JsonResponse(w, http.StatusCreated, res)
	return nil
}

func (authHandler AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	authHandler.AuthService.RefreshToken()
	return nil
}

func (authHandler AuthHandler) ForgetPassword(w http.ResponseWriter, r *http.Request) error {
	authHandler.AuthService.ForgetPassword()
	return nil
}

func (authHandler AuthHandler) Profile(w http.ResponseWriter, r *http.Request) error {
	authHandler.AuthService.Profile()
	return nil
}
