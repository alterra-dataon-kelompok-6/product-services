package main

import (
	"belajar-go-echo/internal/factory"
	"belajar-go-echo/internal/http"
	m "belajar-go-echo/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	f := factory.NewFactory()
	http.NewHttp(app, f)

	m.LogMiddleware(app)

	app.Start(":8080")
}
