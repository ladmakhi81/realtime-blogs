package pkg_utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateHttpReqBody(input any, w http.ResponseWriter) {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, err.Error())
	}
}
