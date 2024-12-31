package pkg_decorators

import (
	"context"
	"net/http"
	"strings"

	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

func ApiAuthDecorator(fn pkg_types.ApiHttpHandler) pkg_types.ApiHttpHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		authHeaderStr := r.Header.Get("Authorization")
		authHeader := strings.Split(authHeaderStr, " ")
		bearer := strings.ToLower(authHeader[0])
		token := authHeader[1]
		if bearer != "bearer" || token == "" {
			return pkg_types.NewClientError(http.StatusUnauthorized, "Unauthorized")
		}
		verifyClaim, verifyErr := pkg_utils.VerifyAccessToken(token)
		if verifyErr != nil {
			return verifyErr
		}
		requestWithAuthContext := r.WithContext(context.WithValue(r.Context(), "AuthUser", verifyClaim))
		return fn(w, requestWithAuthContext)
	}
}
