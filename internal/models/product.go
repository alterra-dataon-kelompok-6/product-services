package model

type Product struct {
	Common

	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type Category struct {
	Common
	Name string `json:"name"`
}
