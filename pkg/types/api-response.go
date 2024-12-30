package pkg_types

import "time"

type ApiResponse struct {
	IsSuccess   bool      `json:"isSuccess"`
	Data        any       `json:"data"`
	RequestedAt time.Time `json:"requestedAt"`
}
