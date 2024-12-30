package pkg_utils

import (
	"encoding/json"
	"net/http"
	"time"

	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

func JsonResponse(w http.ResponseWriter, status int, value any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	res := pkg_types.ApiResponse{IsSuccess: status < 400, Data: value, RequestedAt: time.Now()}
	json.NewEncoder(w).Encode(res)
}
