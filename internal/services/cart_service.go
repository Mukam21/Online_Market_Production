package services

import (
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/Mukam21/Online_Market_Production/internal/repositories"
)

type CartService interface {
	AddToCart(userID uint, cart *models.Cart) error // Изменяем email на userID
	RemoveFromCart(userID uint, productID uint) error
	Checkout(userID uint) error
}

type cartService struct {
	repo repositories.CartRepository
}

func NewCartService(repo repositories.CartRepository) CartService {
	return &cartService{repo: repo}
}

func (s *cartService) AddToCart(userID uint, cart *models.Cart) error {
	cart.UserID = userID
	return s.repo.AddToCart(cart)
}

func (s *cartService) RemoveFromCart(userID uint, productID uint) error {
	return s.repo.RemoveFromCart(userID, productID)
}

func (s *cartService) Checkout(userID uint) error {
	return s.repo.Checkout(userID)
}
