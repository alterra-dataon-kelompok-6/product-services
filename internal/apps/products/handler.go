package products

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
	products, err := h.service.GetAll()

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
	// id, _ := strconv.Atoi(e.Param("id"))
	var payload dto.ProductRequestParams

	if err := (&echo.DefaultBinder{}).BindPathParams(e, &payload); err != nil {
		log.Println(err)
	}
	product, err := h.service.GetById(payload)
	if err != nil || product.ID == 0 {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  false,
			"message": "data not found",
		})
	}
	// categoryRepo := categories.NewRepo(factory.NewFactory().DB)
	// category, _ := categoryRepo.GetById(uint(product.CategoryID))

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   product,
	})
}

func (h handler) Create(e echo.Context) error {
	var payload dto.ProductRequestBodyCreate
	if err := e.Bind(&payload); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}
	product, err := h.service.Create(payload)
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
	// id, _ := strconv.Atoi(e.Param("id"))
	var id dto.ProductRequestParams
	var payload dto.ProductRequestBodyUpdate

	if err := (&echo.DefaultBinder{}).BindPathParams(e, &id); err != nil {
		log.Println(err)
	}

	if err := e.Bind(&payload); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  false,
			"message": "invalid data",
		})
	}

	// var updatedData map[string]interface{} = make(map[string]interface{})

	// if err := e.Bind(&updatedData); err != nil {
	// 	return e.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"status":  false,
	// 		"message": "invalid data",
	// 	})
	// }

	product, err := h.service.Update(id.ID, payload)
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
	// id, _ := strconv.Atoi(e.Param("id"))
	var payload dto.ProductRequestParams

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
