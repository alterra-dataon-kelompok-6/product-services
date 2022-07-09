package reviews

import (
	"log"
	"net/http"

	"product-services/internal/dto"
	"product-services/internal/factory"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h handler) GetAll(e echo.Context) error {
	categories, err := h.service.GetAll()

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
	// id, _ := strconv.Atoi(e.Param("id"))
	var payload dto.ReviewRequestParams

	if err := (&echo.DefaultBinder{}).BindPathParams(e, &payload); err != nil {
		log.Println(err)
	}

	review, err := h.service.GetById(payload)
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
	var payload dto.ReviewRequestBodyCreate
	if err := e.Bind(&payload); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}
	log.Println(payload, "payload create review")
	review, err := h.service.Create(payload)
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
	// id, _ := strconv.Atoi(e.Param("id"))
	var id dto.ReviewRequestParams

	if err := (&echo.DefaultBinder{}).BindPathParams(e, &id); err != nil {
		log.Println(err)
	}

	var payload dto.ReviewRequestBodyUpdate

	if err := e.Bind(&payload); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}

	review, err := h.service.Update(id.ID, payload)
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
	var payload dto.ReviewRequestParams

	if err := (&echo.DefaultBinder{}).BindPathParams(e, &payload); err != nil {
		log.Println(err)
	}
	_, err := h.service.Delete(payload)

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
