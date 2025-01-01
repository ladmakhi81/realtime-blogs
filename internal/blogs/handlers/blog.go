package blogs_handlers

import (
	"encoding/json"
	"net/http"

	blogs_contracts "github.com/ladmakhi81/realtime-blogs/internal/blogs/contracts"
	blogs_types "github.com/ladmakhi81/realtime-blogs/internal/blogs/types"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type BlogHandler struct {
	BlogService blogs_contracts.BlogServiceContract
}

func NewBlogHandler(blogService blogs_contracts.BlogServiceContract) BlogHandler {
	return BlogHandler{BlogService: blogService}
}

func (blogHandler BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) error {
	reqBody := new(blogs_types.CreateBlogReqBody)
	if parseErr := json.NewDecoder(r.Body).Decode(reqBody); parseErr != nil {
		return pkg_types.NewServerError(
			"error in parsing request body",
			"BlogHandler.CreateBlog",
			parseErr.Error(),
		)
	}
	defer r.Body.Close()
	if validateErr := pkg_utils.ValidateHttpReqBody(reqBody); validateErr != nil {
		return pkg_types.NewClientValidationError(validateErr)
	}
	authUserId := r.Context().Value("AuthUser").(*pkg_types.UserAuthClaim).ID

	blog, createBlogErr := blogHandler.BlogService.CreateBlog(*reqBody, authUserId)
	if createBlogErr != nil {
		return createBlogErr
	}
	pkg_utils.JsonResponse(
		w,
		http.StatusCreated,
		blogs_types.NewCreateBlogResponse(
			blog,
		),
	)
	return nil
}

func (blogHandler BlogHandler) DeleteBlogById(w http.ResponseWriter, r *http.Request) {
}

func (blogHandler BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {
}

func (blogHandler BlogHandler) GetBlogById(w http.ResponseWriter, r *http.Request) {
}
