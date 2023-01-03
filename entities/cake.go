package entities

import validation "github.com/go-ozzo/ozzo-validation"

type Cake struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int64  `json:"rating"`
	Image       string `json:"image"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CakePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int64  `json:"rating"`
	Image       string `json:"image"`
}

func (l Cake) Validate() error {
	return validation.ValidateStruct(
		&l,
		validation.Field(&l.Title, validation.Required),
		validation.Field(&l.Description, validation.Required),
		validation.Field(&l.Rating, validation.Required),
		validation.Field(&l.Image, validation.Required),
	)
}
