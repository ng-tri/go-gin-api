package service

import (
	"errors"

	"go-gin-api/internal/database"
	"go-gin-api/internal/model"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	var products []model.Product
	err := database.DB.Find(&products).Error
	return products, err
}

func (s *ProductService) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) Create(p *model.Product) (*model.Product, error) {
	if p.Name == "" {
		return nil, errors.New("name is required")
	}
	err := database.DB.Create(p).Error
	return p, err
}

func (s *ProductService) Update(id uint, data *model.Product) (*model.Product, error) {
	var product model.Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	product.Name = data.Name
	product.Price = data.Price

	err = database.DB.Save(&product).Error
	return &product, err
}

func (s *ProductService) Delete(id uint) error {
	return database.DB.Delete(&model.Product{}, id).Error
}
