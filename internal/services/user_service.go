package services

import (
	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/Mukam21/Online_Market_Production/internal/repositories"
)

// Интерфейс UserService
type UserService interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

// Реализация UserService
type userService struct {
	userRepo repositories.UserRepository
}

// Конструктор
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Создание пользователя
func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

// Поиск пользователя по Email
func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}
