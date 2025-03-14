package repositories

import (
	"errors"

	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

// userRepository is the concrete implementation of UserRepository
type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// CreateUser inserts a new user into the database
func (r *userRepository) CreateUser(user *models.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		// Check for unique constraint violation
		if err, ok := result.Error.(*pq.Error); ok && err.Code == "23505" {
			return errors.New("user with this email already exists")
		}
		return result.Error
	}
	return nil
}

// FindByEmail retrieves a user by email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
