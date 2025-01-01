package blogs_types

type CreateBlogReqBody struct {
	Title      string   `json:"title" validate:"required,min=3"`
	Content    string   `json:"content" validate:"required,min=3"`
	CategoryId uint     `json:"categoryId" validate:"required,gte=1,numeric"`
	Tags       []string `json:"tags" validate:"required,dive"`
}
