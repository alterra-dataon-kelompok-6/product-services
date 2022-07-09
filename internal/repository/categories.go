package repository

import (
	"errors"

	model "product-services/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category model.Category) (*model.Category, error)
	GetAll() (*[]model.Category, error)
	GetById(id uint) (*model.Category, error)
	Update(id uint, category map[string]interface{}) (*model.Category, error)
	Delete(id uint) (bool, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepo(DB *gorm.DB) CategoryRepository {
	return &categoryRepository{DB: DB}
}

func (r *categoryRepository) Create(category model.Category) (*model.Category, error) {
	if err := r.DB.Save(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetAll() (*[]model.Category, error) {
	var categories []model.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *categoryRepository) GetById(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(id uint, updatedData map[string]interface{}) (*model.Category, error) {
	if category, _ := r.GetById(id); category.ID <= 0 {
		return nil, errors.New("data not found")
	}
	var newCategory model.Category
	if err := r.DB.Model(model.Category{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Find(&newCategory).Error; err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func (r *categoryRepository) Delete(id uint) (bool, error) {
	if category, _ := r.GetById(id); category.ID <= 0 {
		return false, errors.New("data not found")
	}
	if err := r.DB.Debug().Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return false, errors.New("failed to delete data")
	}

	return true, nil
}
