package services

import (
	"errors"

	"github.com/EENmalikhanDev/be-golang/models"
	"github.com/EENmalikhanDev/be-golang/storage"
)

func GetAllProducts() []models.Product {
	products := []models.Product{}
	for _, product := range storage.Products {
		products = append(products, product)
	}
	return products
}

func GetProductByID(id string) (models.Product, error) {
	product, exists := storage.Products[id]
	if !exists {
		return models.Product{}, errors.New("product not found")
	}
	return product, nil
}

func CreateProduct(product models.Product) {
	storage.Products[product.ID] = product
}

func UpdateProduct(id string, product models.Product) error {
	_, exists := storage.Products[id]
	if !exists {
		return errors.New("product not found")
	}
	storage.Products[id] = product
	return nil
}

func DeleteProduct(id string) error {
	_, exists := storage.Products[id]
	if !exists {
		return errors.New("product not found")
	}
	delete(storage.Products, id)
	return nil
}
