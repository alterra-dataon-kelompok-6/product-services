package reviews

import (
	"log"

	"product-services/internal/dto"
	"product-services/internal/factory"
	model "product-services/internal/models"
	"product-services/internal/repository"
)

type Service interface {
	Create(payload dto.ReviewRequestBodyCreate) (*model.Review, error)
	GetAll() (*[]model.Review, error)
	GetById(payload dto.ReviewRequestParams) (*model.Review, error)
	Update(id uint, payload dto.ReviewRequestBodyUpdate) (*model.Review, error)
	Delete(payload dto.ReviewRequestParams) (interface{}, error)
}

type service struct {
	ReviewRepository repository.ReviewRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		ReviewRepository: repository.NewReveiewRepo(f.DB),
	}
}

func (s service) Create(payload dto.ReviewRequestBodyCreate) (*model.Review, error) {
	var newReview = model.Review{
		CustomerID: payload.CustomerID,
		ProductID:  payload.ProductID,
		Rating:     payload.Rating,
		Review:     payload.Review,
	}

	review, err := s.ReviewRepository.Create(newReview)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (s service) GetAll() (*[]model.Review, error) {
	categories, err := s.ReviewRepository.GetAll()
	if err != nil || len(*categories) <= 0 {
		return nil, err
	}
	return categories, nil
}

func (s service) GetById(payload dto.ReviewRequestParams) (*model.Review, error) {
	review, err := s.ReviewRepository.GetById(payload.ID)
	if err != nil || review.ID == 0 {
		return nil, err
	}
	return review, nil
}

func (s service) Update(id uint, payload dto.ReviewRequestBodyUpdate) (*model.Review, error) {
	var updatedData map[string]interface{} = make(map[string]interface{})

	if payload.Rating != 0 {
		updatedData["rating"] = payload.Rating
	}
	if payload.Review != "" {
		updatedData["review"] = payload.Review
	}

	review, err := s.ReviewRepository.Update(id, updatedData)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (s service) Delete(payload dto.ReviewRequestParams) (interface{}, error) {
	deleted, err := s.ReviewRepository.Delete(payload.ID)
	if err != nil {
		return nil, err
	}
	log.Println(deleted, "deleted")
	return deleted, err
}
