package reviews

import (
	"net/http"
	"strconv"

	"product-services/internal/factory"
	model "product-services/internal/models"

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
	review, err := h.repository.GetById(uint(id))
	if err != nil || review.ID == 0 {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   review,
	})
}

func (h handler) Create(e echo.Context) error {
	var newReview model.Review
	if err := e.Bind(&newReview); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}
	review, err := h.repository.Create(newReview)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to create data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   review,
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

	review, err := h.repository.Update(uint(id), updatedData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "failed to update data",
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   review,
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
