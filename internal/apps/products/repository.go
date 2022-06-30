package products

import (
	model "belajar-go-echo/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product model.Product) (*model.Product, error)
	GetAll() (*[]model.Product, error)
	GetById(id int) (*model.Product, error)
	Update(id int, product map[string]interface{}) (*model.Product, error)
	Delete(id int) (bool, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) Repository {
	return &repository{DB: DB}
}

func (r *repository) Create(product model.Product) (*model.Product, error) {
	if err := r.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) GetAll() (*[]model.Product, error) {
	var products []model.Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (r *repository) GetById(id int) (*model.Product, error) {
	var product model.Product
	if err := r.DB.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Update(id int, updatedData map[string]interface{}) (*model.Product, error) {
	var product model.Product
	if err := r.DB.Model(model.Product{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *repository) Delete(id int) (bool, error) {
	if err := r.DB.Delete(model.Product{}, id).Error; err != nil {
		return false, err
	}

	return true, nil
}
