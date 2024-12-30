package categories_types

type ModifyCategoryReqBody struct {
	Title string `json:"title" validate:"required,min=3"`
}
