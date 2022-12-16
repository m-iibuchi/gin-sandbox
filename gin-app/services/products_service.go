package services

import (
	"gin-sandbox/domains/products"
	"gin-sandbox/utils/errors"

	"gorm.io/gorm"
)

func CreateProduct(product products.Product) (*products.Product, *errors.ApiErr) {
	if err := product.Validate(); err != nil {
		return nil, err
	}

	if err := product.Save(); err != nil {
		return nil, err
	}

	return &product, nil
}

func GetProduct(productID uint) (*products.Product, *errors.ApiErr) {
	p := &products.Product{Model: gorm.Model{ID: productID}}
	if err := p.Get(); err != nil {
		return nil, err
	}
	return p, nil
}
