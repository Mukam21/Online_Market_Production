package services

import (
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/Mukam21/Online_Market_Production/internal/repositories"
)

type OrderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(userID uint) (*models.Order, error) {
	return s.repo.CreateOrder(userID)
}
