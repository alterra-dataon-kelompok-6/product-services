package dto

type CategoryRequestParams struct {
	ID uint `json:"id"`
}

type CategoryRequestBody struct {
	Category string `json:"category" validate:"required"`
}
