package services

import (
	"errors"

	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/Mukam21/Online_Market_Production/internal/repositories"
)

// Интерфейс для ProductService
type ProductServiceInterface interface {
	AddProduct(product *models.Product) error
	GetProducts() ([]models.Product, error)
}

// Реализация ProductService
type ProductService struct {
	Repo repositories.ProductRepository
}

// Конструктор
func NewProductService(repo repositories.ProductRepository) ProductServiceInterface {
	return &ProductService{Repo: repo}
}

// Добавление продукта
func (s *ProductService) AddProduct(product *models.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}
	return s.Repo.CreateProduct(product)
}

// Получение списка продуктов
func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repo.GetProducts()
}
