package dto

type ProductRequestParams struct {
	ID uint
}

type ProductRequestBodyCreate struct {
	CategoryID  uint   `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description"`
}

type ProductRequestBodyUpdate struct {
	CategoryID  *uint   `json:"category_id"`
	Name        *string `json:"name"`
	Stock       *int    `json:"stock"`
	Price       *int    `json:"price"`
	Description *string `json:"description"`
}
