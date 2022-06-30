package categories

import (
	model "belajar-go-echo/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(category model.Category) (*model.Category, error)
	GetAll() (*[]model.Category, error)
	GetById(id int) (*model.Category, error)
	Update(id int, category map[string]interface{}) (*model.Category, error)
	Delete(id int) (bool, error)
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

func (r *repository) GetById(id int) (*model.Category, error) {
	var category model.Category
	if err := r.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repository) Update(id int, updatedData map[string]interface{}) (*model.Category, error) {
	var category model.Category
	if err := r.DB.Model(model.Category{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Find(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repository) Delete(id int) (bool, error) {
	if err := r.DB.Delete(model.Category{}, id).Error; err != nil {
		return false, err
	}

	return true, nil
}
