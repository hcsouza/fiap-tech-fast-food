package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/product"
)

type productHandler struct {
	interactor product.IProductUseCase
}

func NewProductHandler(gRouter *gin.RouterGroup, interactor product.IProductUseCase) {
	handler := &productHandler{
		interactor: interactor,
	}

	gRouter.GET("/product/:category", handler.GetProductByCategoryHandler)
	gRouter.POST("/product", handler.CreateProductHandler)
	gRouter.PUT("/product/:id", handler.UpdateProductHandler)
	gRouter.DELETE("/product/:id", handler.DeleteProductHandler)
}

func (handler *productHandler) GetProductByCategoryHandler(c *gin.Context) {
	category, exists := c.Params.Get("category")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category is required"})
		return
	}

	actions, err := handler.interactor.GetByCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actions)
}

func (handler *productHandler) CreateProductHandler(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.interactor.Create(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (handler *productHandler) UpdateProductHandler(c *gin.Context) {
	productId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is required"})
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.interactor.Update(productId, product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (handler *productHandler) DeleteProductHandler(c *gin.Context) {
	productId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is required"})
		return
	}

	err := handler.interactor.Delete(productId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
