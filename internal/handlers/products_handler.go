package handlers

import (
	"net/http"

	"github.com/Mukam21/Online_Market_Production/internal/models"
	"github.com/Mukam21/Online_Market_Production/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service services.ProductServiceInterface
}

func NewProductHandler(service services.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data: " + err.Error()})
		return
	}

	// Дополнительная валидация
	if product.Name == "" || product.Price <= 0 || product.Quantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, Price (>0), and Quantity (>=0) are required"})
		return
	}

	if err := h.Service.AddProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product added", "id": product.ID})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.Service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
