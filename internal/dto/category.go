package dto

type CategoryRequestParams struct {
	ID uint `json:"id" param:"id" query:"id" form:"id" xml:"id"`
}

type CategoryRequestBody struct {
	Category string `json:"category" validate:"required"`
}
