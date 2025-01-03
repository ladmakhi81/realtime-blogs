package pkg_decorators

import (
	"fmt"
	"net/http"

	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

func ApiErrorDecorator(fn pkg_types.ApiHttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if serverErr, isServerErr := err.(*pkg_types.ServerError); isServerErr {
			pkg_utils.JsonResponse(
				w,
				http.StatusInternalServerError,
				map[string]string{"message": "internal server error"},
			)
			fmt.Println(serverErr)
			// store error in files
			return
		}
		if clientErr, isClientErr := err.(*pkg_types.ClientError); isClientErr {
			pkg_utils.JsonResponse(
				w,
				clientErr.StatusCode,
				map[string]string{"message": clientErr.Message},
			)
			return
		}
		if validationErr, isValidationErr := err.(*pkg_types.ClientValidationError); isValidationErr {
			pkg_utils.JsonResponse(
				w,
				validationErr.StatusCode,
				map[string]any{"message": validationErr.Message, "errors": validationErr.Detail},
			)
			return
		}

	}
}
