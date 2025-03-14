package repositories

import (
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"gorm.io/gorm"
)

// Интерфейс для репозитория продуктов
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProducts() ([]models.Product, error)
}

// Реализация ProductRepository
type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{DB: db}
}

// Создание нового продукта
func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.DB.Create(product).Error
}

// Получение списка всех продуктов
func (r *productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
