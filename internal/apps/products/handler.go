package products

import (
	"belajar-go-echo/internal/factory"
	model "belajar-go-echo/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	repository Repository
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		repository: NewRepo(f.DB),
	}
}

func (h handler) GetAll(e echo.Context) error {
	products, err := h.repository.GetAll()

	if err != nil || len(*products) <= 0 {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   products,
	})

}

func (h handler) GetById(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	product, err := h.repository.GetById(uint(id))
	if err != nil || product.ID == 0 {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   product,
	})
}

func (h handler) Create(e echo.Context) error {
	var newProduct model.Product
	if err := e.Bind(&newProduct); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}
	product, err := h.repository.Create(newProduct)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to create data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   product,
	})
}

func (h handler) Update(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	var updatedData map[string]interface{} = make(map[string]interface{})

	if err := e.Bind(&updatedData); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}

	product, err := h.repository.Update(uint(id), updatedData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to update data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   product,
	})

}

func (h handler) Delete(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	_, err := h.repository.Delete(uint(id))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to delete data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "data has been deleted",
	})
}
