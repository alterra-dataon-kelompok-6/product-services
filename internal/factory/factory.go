package factory

import (
	"belajar-go-echo/database"

	"gorm.io/gorm"
)

type Factory struct {
	DB *gorm.DB
}

func NewFactory() *Factory {
	return &Factory{
		DB: database.GetConnection(),
	}
}
