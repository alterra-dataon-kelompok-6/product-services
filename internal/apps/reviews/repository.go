package reviews

import (
	"errors"

	model "product-services/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(review model.Review) (*model.Review, error)
	GetAll() (*[]model.Review, error)
	GetById(id uint) (*model.Review, error)
	Update(id uint, review map[string]interface{}) (*model.Review, error)
	Delete(id uint) (bool, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) Repository {
	return &repository{DB: DB}
}

func (r *repository) Create(review model.Review) (*model.Review, error) {
	if err := r.DB.Save(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *repository) GetAll() (*[]model.Review, error) {
	var categories []model.Review
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *repository) GetById(id uint) (*model.Review, error) {
	var review model.Review
	if err := r.DB.Where("id = ?", id).Find(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *repository) Update(id uint, updatedData map[string]interface{}) (*model.Review, error) {
	if review, _ := r.GetById(id); review.ID <= 0 {
		return nil, errors.New("data not found")
	}
	var newReview model.Review
	if err := r.DB.Model(model.Review{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Find(&newReview).Error; err != nil {
		return nil, err
	}

	return &newReview, nil
}

func (r *repository) Delete(id uint) (bool, error) {
	if review, _ := r.GetById(id); review.ID <= 0 {
		return false, errors.New("data not found")
	}
	if err := r.DB.Debug().Where("id = ?", id).Delete(&model.Review{}).Error; err != nil {
		return false, errors.New("failed to delete data")
	}

	return true, nil
}
