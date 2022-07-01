package products

import (
	model "belajar-go-echo/internal/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product model.Product) (*model.Product, error)
	GetAll() (*[]model.Product, error)
	GetById(id uint) (*model.Product, error)
	Update(id uint, product map[string]interface{}) (*model.Product, error)
	Delete(id uint) (bool, error)
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

func (r *repository) GetById(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.DB.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Update(id uint, updatedData map[string]interface{}) (*model.Product, error) {
	if product, _ := r.GetById(id); product.ID <= 0 {
		return nil, errors.New("data not found")
	}
	var newProduct model.Product
	if err := r.DB.Model(model.Product{}).Where("id = ?", id).Updates(updatedData).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Where("id = ?", id).Find(&newProduct).Error; err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func (r *repository) Delete(id uint) (bool, error) {
	if product, _ := r.GetById(id); product.ID <= 0 {
		return false, errors.New("data not found")
	}
	if err := r.DB.Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
