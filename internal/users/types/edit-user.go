package users_types

type EditUserReqBody struct {
	ProfileURL string `json:"profileUrl" validate:"min=3"`
	FirstName  string `json:"firstName" validate:"min=3"`
	LastName   string `json:"lastName" validate:"min=3"`
}
