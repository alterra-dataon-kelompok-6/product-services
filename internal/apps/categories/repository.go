package categories

import (
	model "belajar-go-echo/internal/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(category model.Category) (*model.Category, error)
	GetAll() (*[]model.Category, error)
	GetById(id uint) (*model.Category, error)
	Update(id uint, category map[string]interface{}) (*model.Category, error)
	Delete(id uint) (bool, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) Repository {
	return &repository{DB: DB}
}

func (r *repository) Create(category model.Category) (*model.Category, error) {
	if err := r.DB.Save(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repository) GetAll() (*[]model.Category, error) {
	var categories []model.Category
	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *repository) GetById(id uint) (*model.Category, error) {
	var category model.Category
	if err := r.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repository) Update(id uint, updatedData map[string]interface{}) (*model.Category, error) {
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

func (r *repository) Delete(id uint) (bool, error) {
	if category, _ := r.GetById(id); category.ID <= 0 {
		return false, errors.New("data not found")
	}
	if err := r.DB.Debug().Where("id = ?", id).Delete(&model.Category{}).Error; err != nil {
		return false, errors.New("failed to delete data")
	}

	return true, nil
}
