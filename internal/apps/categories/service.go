package categories

import (
	"errors"
	"log"

	"product-services/internal/dto"
	"product-services/internal/factory"
	model "product-services/internal/models"
	"product-services/internal/repository"
)

type Service interface {
	Create(payload dto.CategoryRequestBody) (*model.Category, error)
	GetAll() (*[]model.Category, error)
	GetById(payload dto.CategoryRequestParams) (*dto.CategoryResponseGetById, error)
	Update(id uint, payload dto.CategoryRequestBody) (*model.Category, error)
	Delete(payload dto.CategoryRequestParams) (interface{}, error)
}

type service struct {
	CategoryRepository repository.CategoryRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		CategoryRepository: repository.NewCategoryRepo(f.DB),
	}
}

var ProductRepo = repository.NewProductRepo(factory.NewFactory().DB)

func (s service) Create(payload dto.CategoryRequestBody) (*model.Category, error) {
	var newCategory = model.Category{
		Category: payload.Category,
	}

	category, err := s.CategoryRepository.Create(newCategory)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s service) GetAll() (*[]model.Category, error) {
	categories, err := s.CategoryRepository.GetAll()
	if err != nil || len(*categories) <= 0 {
		return nil, errors.New("data not found")
	}
	return categories, nil
}

func (s service) GetById(payload dto.CategoryRequestParams) (*dto.CategoryResponseGetById, error) {
	category, err := s.CategoryRepository.GetById(payload.ID)
	if err != nil || category.ID == 0 {
		log.Println(err, category, payload)
		return nil, errors.New("data not found")
	}
	// result
	result := new(dto.CategoryResponseGetById)
	result.Category = *category

	products, err := ProductRepo.GetAll()
	if err != nil {
		return nil, err
	}

	result.Products = *products

	return result, nil
}

func (s service) Update(id uint, payload dto.CategoryRequestBody) (*model.Category, error) {
	var updatedData map[string]interface{} = make(map[string]interface{})

	if payload.Category != "" {
		updatedData["category"] = payload.Category
	}

	category, err := s.CategoryRepository.Update(id, updatedData)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s service) Delete(payload dto.CategoryRequestParams) (interface{}, error) {
	deleted, err := s.CategoryRepository.Delete(payload.ID)
	if err != nil {
		return nil, err
	}
	log.Println(deleted, "deleted")
	return deleted, err
}
