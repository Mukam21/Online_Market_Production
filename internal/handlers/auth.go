package handlers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key") // Секретный ключ для подписания токенов

// Структура для запроса авторизации
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Структура для JWT
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Функция для генерации токена
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Время жизни токена (24 часа)

	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "marketplace",
		},
	}

	// Создаем новый JWT токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) // Подписываем токен
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Функция для проверки и парсинга токена
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
