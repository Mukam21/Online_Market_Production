package repositories

import (
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"gorm.io/gorm"
)

// Интерфейс для репозитория корзины
type CartRepository interface {
	AddToCart(cart *models.Cart) error
	RemoveFromCart(userID uint, productID uint) error
	Checkout(userID uint) error
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{DB: db}
}

type cartRepository struct {
	DB *gorm.DB
}

func (r *cartRepository) AddToCart(cart *models.Cart) error {
	return r.DB.Create(cart).Error
}

func (r *cartRepository) RemoveFromCart(userID uint, productID uint) error {
	return r.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Cart{}).Error
}

func (r *cartRepository) Checkout(userID uint) error {
	return r.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}
