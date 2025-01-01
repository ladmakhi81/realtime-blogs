package users_types

type EditUserReqBody struct {
	ProfileURL *string `json:"profileUrl,omitempty" validate:"min=3"`
	FirstName  *string `json:"firstName,omitempty" validate:"min=3"`
	LastName   *string `json:"lastName,omitempty" validate:"min=3"`
}
