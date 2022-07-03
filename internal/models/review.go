package model

type Review struct {
	Common
	ProductID  uint   `json:"product_id"`
	CustomerID uint   `json:"customer_id"`
	Rating     uint   `json:"rating"`
	Review     string `json:"review"`
}
