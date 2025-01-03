package categories_handlers

import (
	"encoding/json"
	"net/http"

	categories_contracts "github.com/ladmakhi81/realtime-blogs/internal/categories/contracts"
	categories_types "github.com/ladmakhi81/realtime-blogs/internal/categories/types"
	users_contracts "github.com/ladmakhi81/realtime-blogs/internal/users/contracts"
	pkg_types "github.com/ladmakhi81/realtime-blogs/pkg/types"
	pkg_utils "github.com/ladmakhi81/realtime-blogs/pkg/utils"
)

type CategoryHandler struct {
	CategoryService categories_contracts.CategoryServiceContract
	UserService     users_contracts.UserServiceContract
}

func NewCategoryHandler(
	categoryService categories_contracts.CategoryServiceContract,
	userService users_contracts.UserServiceContract,
) CategoryHandler {
	return CategoryHandler{
		CategoryService: categoryService,
		UserService:     userService,
	}
}

func (categoryHandler CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	authUserId := r.Context().Value("AuthUser").(*pkg_types.UserAuthClaim).ID
	authUser, authUserErr := categoryHandler.UserService.FindUserById(authUserId)
	if authUserErr != nil {
		return authUserErr
	}
	reqBody := new(categories_types.ModifyCategoryReqBody)
	if decodeErr := json.NewDecoder(r.Body).Decode(reqBody); decodeErr != nil {
		return pkg_types.NewServerError(
			"error in parsing json request body",
			"CategoryHandler.CreateCategory.NewDecoder",
			decodeErr.Error(),
		)
	}
	if validationErr := pkg_utils.ValidateHttpReqBody(reqBody); validationErr != nil {
		return pkg_types.NewClientValidationError(validationErr)
	}
	category, createCategoryErr := categoryHandler.CategoryService.CreateCategory(*reqBody, authUser)
	if createCategoryErr != nil {
		return createCategoryErr
	}
	pkg_utils.JsonResponse(w, http.StatusCreated, category)
	return nil
}

func (categoryHandler CategoryHandler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) error {
	authUserId := r.Context().Value("AuthUser").(*pkg_types.UserAuthClaim).ID
	var categoryId, parseErr = pkg_utils.ExtractNumericRouteParam(r, "id")
	if parseErr != nil {
		return parseErr
	}
	deleteCategoryErr := categoryHandler.CategoryService.DeleteCategoryById(categoryId, authUserId)
	if deleteCategoryErr != nil {
		return deleteCategoryErr
	}
	pkg_utils.JsonResponse(w, http.StatusOK, nil)
	return nil
}

func (categoryHandler CategoryHandler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) error {
	authUserId := r.Context().Value("AuthUser").(*pkg_types.UserAuthClaim).ID
	var categoryId, parseErr = pkg_utils.ExtractNumericRouteParam(r, "id")
	if parseErr != nil {
		return parseErr
	}
	reqBody := new(categories_types.ModifyCategoryReqBody)
	if decodeErr := json.NewDecoder(r.Body).Decode(reqBody); decodeErr != nil {
		return pkg_types.NewServerError(
			"error in parsing input values",
			"CategoryHandler.UpdateCategoryById",
			decodeErr.Error(),
		)
	}
	defer r.Body.Close()
	if validateErr := pkg_utils.ValidateHttpReqBody(reqBody); validateErr != nil {
		return pkg_types.NewClientValidationError(validateErr)
	}
	if updateErr := categoryHandler.CategoryService.UpdateCategoryById(categoryId, authUserId, *reqBody); updateErr != nil {
		return updateErr
	}
	return nil
}

func (categoryHandler CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) error {
	page, limit, paginationParseErr := pkg_utils.ExtractPaginationQuery(r.URL.Query())
	if paginationParseErr != nil {
		return paginationParseErr
	}
	categories, categoriesCount, err := categoryHandler.CategoryService.GetCategories(page, limit)
	if err != nil {
		return err
	}
	pkg_utils.JsonResponse(
		w,
		http.StatusOK,
		pkg_types.NewDatasourcePagination(
			*categories,
			pkg_utils.CalcTotalPaginationPage(limit, categoriesCount),
			categoriesCount,
			page,
		),
	)
	return nil
}
