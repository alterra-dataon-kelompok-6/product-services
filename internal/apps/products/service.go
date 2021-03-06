package products

import (
	"errors"
	"log"

	"product-services/internal/dto"
	"product-services/internal/factory"
	model "product-services/internal/models"
	"product-services/internal/repository"
)

type Service interface {
	Create(payload dto.ProductRequestBodyCreate) (*model.Product, error)
	GetAll() (*[]model.Product, error)
	GetById(payload dto.ProductRequestParams) (*dto.ProductResponseGetById, error)
	Update(id uint, payload dto.ProductRequestBodyUpdate) (*model.Product, error)
	Delete(payload dto.ProductRequestParams) (interface{}, error)
}

type service struct {
	ProductRepository repository.ProductRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		ProductRepository: repository.NewProductRepo(f.DB),
	}
}

var CategoryRepo = repository.NewCategoryRepo(factory.NewFactory().DB)

func (s service) Create(payload dto.ProductRequestBodyCreate) (*model.Product, error) {
	var newProduct = model.Product{
		CategoryID:  payload.CategoryID,
		Name:        payload.Name,
		Stock:       payload.Stock,
		Price:       payload.Price,
		Image:       "",
		Description: payload.Description,
	}

	product, err := s.ProductRepository.Create(newProduct)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s service) GetAll() (*[]model.Product, error) {
	products, err := s.ProductRepository.GetAll()
	if err != nil || len(*products) <= 0 {
		return nil, errors.New("data not found")
	}
	return products, nil
}

func (s service) GetById(payload dto.ProductRequestParams) (*dto.ProductResponseGetById, error) {
	product, err := s.ProductRepository.GetById(payload.ID)
	if err != nil || product.ID == 0 {
		log.Println(err, product, payload)
		return nil, errors.New("data not found")
	}
	result := new(dto.ProductResponseGetById)
	result.Product = *product

	// get category by product.category_id
	category, err := CategoryRepo.GetById(product.CategoryID)
	if err != nil {
		return nil, err
	}

	result.Category.Category = category.Category
	result.Category.CategoryID = category.ID

	return result, nil
}

func (s service) Update(id uint, payload dto.ProductRequestBodyUpdate) (*model.Product, error) {
	var updatedData map[string]interface{} = make(map[string]interface{})

	if payload.CategoryID != nil {
		updatedData["category_id"] = payload.CategoryID
	}
	if payload.Name != nil {
		updatedData["name"] = payload.Name
	}
	if payload.Stock != nil {
		updatedData["stock"] = payload.Stock
	}
	if payload.Price != nil {
		updatedData["price"] = payload.Price
	}
	if payload.Description != nil {
		updatedData["description"] = payload.Description
	}

	product, err := s.ProductRepository.Update(id, updatedData)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s service) Delete(payload dto.ProductRequestParams) (interface{}, error) {
	deleted, err := s.ProductRepository.Delete(payload.ID)
	if err != nil {
		return nil, err
	}
	log.Println(deleted, "deleted")
	return deleted, err
}
