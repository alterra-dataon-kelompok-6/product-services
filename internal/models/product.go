package model

type Product struct {
	Common

	CategoryID  uint   `json:"category_id"`
	Name        string `json:"name"`
	Stock       uint   `json:"stock"`
	Price       uint   `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
