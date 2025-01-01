package users_types

type UploadFileResponse struct {
	Filename string `json:"filename"`
}

func NewUploadedFileResponse(filename string) UploadFileResponse {
	return UploadFileResponse{Filename: filename}
}
