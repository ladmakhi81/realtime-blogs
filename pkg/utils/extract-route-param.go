package pkg_utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

func ExtractNumericRouteParam(r *http.Request, paramName string) (uint, error) {
	params := mux.Vars(r)
	param := params[paramName]
	var data uint
	if parsedId, parsedErr := strconv.Atoi(param); parsedErr != nil {
		return 0, pkg_types.NewClientError(
			http.StatusBadRequest,
			fmt.Sprintf("invalid %s provided", paramName),
		)
	} else {
		data = uint(parsedId)
	}
	return data, nil
}
