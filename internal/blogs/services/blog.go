package blogs_services

import (
	"net/http"

	blogs_contracts "github.com/ladmakhi81/realtime-blogs/internal/blogs/contracts"
	blogs_entities "github.com/ladmakhi81/realtime-blogs/internal/blogs/entities"
	blogs_types "github.com/ladmakhi81/realtime-blogs/internal/blogs/types"
	categories_contracts "github.com/ladmakhi81/realtime-blogs/internal/categories/contracts"
	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
)

type BlogService struct {
	BlogRepository  blogs_contracts.BlogRepositoryContract
	CategoryService categories_contracts.CategoryServiceContract
	UserService     users_contracts.UserServiceContract
}

func NewBlogService(
	blogRepo blogs_contracts.BlogRepositoryContract,
	categoryService categories_contracts.CategoryServiceContract,
	userService users_contracts.UserServiceContract,
) BlogService {
	return BlogService{
		BlogRepository:  blogRepo,
		CategoryService: categoryService,
		UserService:     userService,
	}
}

func (blogService BlogService) CreateBlog(reqBody blogs_types.CreateBlogReqBody, creatorId uint) (*blogs_entities.Blog, error) {
	category, findCategoryErr := blogService.CategoryService.GetCategoryById(reqBody.CategoryId)
	if findCategoryErr != nil {
		return nil, findCategoryErr
	}
	authUser, authUserErr := blogService.UserService.FindUserById(creatorId)
	if authUserErr != nil {
		return nil, authUserErr
	}
	blog := blogs_entities.NewBlog(
		reqBody.Title,
		reqBody.Content,
		authUser,
		category,
		reqBody.Tags,
	)
	if createBlogErr := blogService.BlogRepository.CreateBlog(blog); createBlogErr != nil {
		return nil, pkg_types.NewServerError(
			"error in creating blog",
			"BlogService.CreateBlog",
			createBlogErr.Error(),
		)
	}
	return blog, nil
}

func (blogService BlogService) GetBlogById(id uint) (*blogs_entities.Blog, error) {
	blog, findBlogErr := blogService.BlogRepository.GetBlogById(id)
	if findBlogErr != nil {
		return nil, pkg_types.NewServerError(
			"error in finding blog by id",
			"BlogService.GetBlogById",
			findBlogErr.Error(),
		)
	}
	if blog == nil {
		return nil, pkg_types.NewClientError(
			http.StatusNotFound,
			"blog not found",
		)
	}
	return blog, nil
}

func (blogService BlogService) DeleteBlogById(id uint, creatorId uint) error {
	blog, err := blogService.GetBlogById(id)
	if err != nil {
		return err
	}
	if blog.CreatedBy.ID != creatorId {
		return pkg_types.NewClientError(
			http.StatusForbidden,
			"only the creator of blog can delete this blog",
		)
	}
	deleteErr := blogService.BlogRepository.DeleteBlogById(blog.ID)
	if deleteErr != nil {
		return pkg_types.NewServerError(
			"error in deleting blog by id",
			"BlogService.DeleteBlogById",
			deleteErr.Error(),
		)
	}
	return nil
}
