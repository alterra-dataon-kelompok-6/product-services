package dto

import model "product-services/internal/models"

type ReviewRequestParams struct {
	ID uint
}

type ReviewRequestBodyCreate struct {
	CustomerID uint   `json:"customer_id" validate:"required"`
	ProductID  uint   `json:"product_id" validate:"required"`
	Rating     uint   `json:"rating" validate:"required"`
	Review     string `json:"review"`
}

// for update date use review id so customer_id and product_id removed
type ReviewRequestBodyUpdate struct {
	Rating uint   `json:"rating" validate:"required"`
	Review string `json:"review"`
}

type ReviewResponseBodyCreate struct {
	Status bool
	Data   model.Review
}
