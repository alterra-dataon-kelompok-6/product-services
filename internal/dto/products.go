package dto

type ProductRequestParams struct {
	ID uint
}

type ProductRequestBody struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
