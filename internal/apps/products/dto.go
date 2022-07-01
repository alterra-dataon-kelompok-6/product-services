package products

type ProductRequestParams struct {
	ID int
}

type ProductRequestBody struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
