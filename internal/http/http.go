package http

import (
	"product-services/internal/apps/categories"
	"product-services/internal/apps/products"
	"product-services/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	products.NewHandler(f).Route(e.Group("/products"))
	categories.NewHandler(f).Route(e.Group("/categories"))
}
