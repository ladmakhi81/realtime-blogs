package pkg_types

type ApiResponse struct {
	IsSuccess   bool   `json:"isSuccess"`
	Data        any    `json:"data"`
	RequestedAt string `json:"requestedAt"`
}
