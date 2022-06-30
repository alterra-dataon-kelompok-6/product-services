package http

import (
	"belajar-go-echo/internal/apps/categories"
	"belajar-go-echo/internal/apps/products"
	"belajar-go-echo/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	products.NewHandler(f).Route(e.Group("/products"))
	categories.NewHandler(f).Route(e.Group("/categories"))
}
