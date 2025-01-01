package users_handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	users_types "github.com/ladmakhi81/realtime-blogs/internal/users/types"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type UserHandler struct {
	UserService users_contracts.UserServiceContract
}

func NewUserHandler(userService users_contracts.UserServiceContract) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}

func (userHandler UserHandler) EditUser(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(users_types.EditUserReqBody)
	if parseJsonErr := json.NewDecoder(r.Body).Decode(reqBody); parseJsonErr != nil {
		return pkg_types.NewServerError(
			"error in parsing request body",
			"UserHandler.EditUser",
			parseJsonErr.Error(),
		)
	}
	defer r.Body.Close()
	authUserId := r.Context().Value("AuthUser").(*pkg_types.UserAuthClaim).ID

	if updateErr := userHandler.UserService.UpdateUser(authUserId, *reqBody); updateErr != nil {
		return updateErr
	}
	return nil
}

func (userHandler UserHandler) UploadProfile(w http.ResponseWriter, r *http.Request) error {
	file, fileHandler, parsedFileErr := r.FormFile("image")
	if parsedFileErr != nil {
		return pkg_types.NewClientError(
			http.StatusBadRequest,
			"invalid file uploaded",
		)
	}
	defer file.Close()
	if isImage := strings.HasPrefix(fileHandler.Header.Get("Content-Type"), "image/"); !isImage {
		return pkg_types.NewClientError(
			http.StatusBadRequest,
			"invalid file uploaded",
		)
	}
	uploadedFile, uploadedFileErr := pkg_utils.FileUploader(file, fileHandler)
	if uploadedFileErr != nil {
		return uploadedFileErr
	}
	pkg_utils.JsonResponse(
		w,
		http.StatusCreated,
		users_types.NewUploadedFileResponse(uploadedFile),
	)
	return nil
}
