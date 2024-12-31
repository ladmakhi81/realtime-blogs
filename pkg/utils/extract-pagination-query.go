package pkg_utils

import (
	"net/http"
	"net/url"
	"strconv"

	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

func ExtractPaginationQuery(queryParams url.Values) (page, limit uint, convertErr error) {
	page = 0
	limit = 10

	pageQuery := queryParams.Get("page")
	limitQuery := queryParams.Get("limit")

	if len(pageQuery) > 0 {
		if data, err := strconv.Atoi(pageQuery); err != nil {
			convertErr = pkg_types.NewClientError(http.StatusBadRequest, "invalid page number")
			return
		} else {
			page = uint(data)
		}
	}

	if len(limitQuery) > 0 {
		if data, err := strconv.Atoi(limitQuery); err != nil {
			convertErr = pkg_types.NewClientError(http.StatusBadRequest, "invalid limit number")
			return
		} else {
			limit = uint(data)
		}
	}

	return
}
