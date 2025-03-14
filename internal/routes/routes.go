package routes

import (
	"net/http"
	"strings"

	"github.com/Mukam21/Online_Market_Production/internal/handlers"
	"github.com/Mukam21/Online_Market_Production/internal/jwt"
	"github.com/Mukam21/Online_Market_Production/internal/repositories"
	"github.com/Mukam21/Online_Market_Production/internal/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Middleware для проверки JWT
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Сохраняем UserID в контексте
		c.Set("userID", claims.UserID) // Сохраняем UserID, если это необходимо для дальнейшего использования
		c.Next()
	}
}

// Основной роутер
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Репозитории
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Сервисы
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(orderRepo)

	// Хендлеры
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Группы маршрутов
	api := r.Group("/api")
	{
		// Auth
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}

		// Products (публичный просмотр товаров)
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
		}
	}

	// Защищённые маршруты
	protected := api.Group("/", authMiddleware())
	{
		// Products (создание товаров)
		protected.POST("/products", productHandler.CreateProduct)

		// Cart
		protected.POST("/cart", cartHandler.AddToCart)
		protected.DELETE("/cart/remove", cartHandler.RemoveFromCart)
		protected.POST("/cart/checkout", cartHandler.Checkout)

		// Orders
		protected.POST("/orders", orderHandler.CreateOrder)
	}

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
