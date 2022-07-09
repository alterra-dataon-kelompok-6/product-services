package dto

import model "product-services/internal/models"

type ProductRequestParams struct {
	ID uint `json:"id" param:"id" query:"id" form:"id" xml:"id"`
}

type ProductRequestBodyCreate struct {
	CategoryID  uint   `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Stock       uint   `json:"stock" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	Description string `json:"description"`
}

type ProductRequestBodyUpdate struct {
	CategoryID  *uint   `json:"category_id"`
	Name        *string `json:"name"`
	Stock       *uint   `json:"stock"`
	Price       *uint   `json:"price"`
	Description *string `json:"description"`
}

type ProductCategory struct {
	CategoryID uint   `json:"category_id"`
	Category   string `json:"category"`
}

type ProductResponseGetById struct {
	model.Product
	Category ProductCategory `json:"category"`
}
