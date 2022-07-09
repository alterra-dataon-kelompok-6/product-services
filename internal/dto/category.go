package dto

import model "product-services/internal/models"

type CategoryRequestParams struct {
	ID uint `json:"id" param:"id" query:"id" form:"id" xml:"id"`
}

type CategoryRequestBody struct {
	Category string `json:"category" validate:"required"`
}

type CategoryResponseGetById struct {
	model.Category
	Products []model.Product `json:"products"`
}
