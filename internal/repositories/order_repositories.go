package repositories

import (
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"gorm.io/gorm"
)

// Интерфейс для репозитория заказов
type OrderRepository interface {
	CreateOrder(userID uint) (*models.Order, error)
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{DB: db}
}

type orderRepository struct {
	DB *gorm.DB
}

func (r *orderRepository) CreateOrder(userID uint) (*models.Order, error) {
	order := &models.Order{UserID: userID, Total: 0, Status: "pending"}
	if err := r.DB.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
