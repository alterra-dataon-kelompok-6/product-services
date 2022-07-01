package categories

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
	categories, err := h.repository.GetAll()

	if err != nil || len(*categories) <= 0 {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   categories,
	})

}

func (h handler) GetById(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	category, err := h.repository.GetById(id)
	if err != nil || category.ID == 0 {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   category,
	})
}

func (h handler) Create(e echo.Context) error {
	var newCategory model.Category
	if err := e.Bind(&newCategory); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}
	category, err := h.repository.Create(newCategory)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to create data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   category,
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

	category, err := h.repository.Update(id, updatedData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to update data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   category,
	})

}

func (h handler) Delete(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	_, err := h.repository.Delete(id)
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
