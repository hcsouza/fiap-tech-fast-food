package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/order"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/orderStatus"
	"net/http"
	"strings"
)

type orderHandler struct {
	interactor order.IOrderUseCase
}

func NewOrderHandler(gRouter *gin.RouterGroup, interactor order.IOrderUseCase) {
	handler := &orderHandler{
		interactor: interactor,
	}

	gRouter.GET("/order/:id", handler.FindById)
	gRouter.GET("/order/status/:status", handler.GetAllByStatus)
	gRouter.POST("/order", handler.CreateOrderHandler)
	gRouter.PUT("/order/:id", handler.UpdateOrderHandler)
}

func (handler *orderHandler) FindById(c *gin.Context) {
	orderId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	order, err := handler.interactor.FindById(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (handler *orderHandler) GetAllByStatus(c *gin.Context) {
	status, exists := c.Params.Get("status")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	orderSts, err := orderStatus.ParseOrderStatus(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	orders, err := handler.interactor.GetAllByStatus(orderSts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (handler *orderHandler) CreateOrderHandler(c *gin.Context) {
	var order order.OrderCreateDTO

	if err := c.ShouldBindJSON(&order); err != nil {
		var verr validator.ValidationErrors
		var msgFieldError string
		if errors.As(err, &verr) {
			msgFieldError = strings.Split(verr[0].Namespace(), ".")[1]
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(fieldErrMsg, msgFieldError)})
			return
		}
	}

	if len(order.OrderItemsDTO) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "at least one product must be selected"})
		return
	}

	orderId, err := handler.interactor.CreateOrder(order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OrderId": orderId})
}

func (handler *orderHandler) UpdateOrderHandler(c *gin.Context) {
	orderId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	var order order.OrderUpdateDTO
	if err := c.ShouldBindJSON(&order); err != nil {
		var verr validator.ValidationErrors
		var msgFieldError string
		if errors.As(err, &verr) {
			msgFieldError = strings.Split(verr[0].Namespace(), ".")[1]
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(fieldErrMsg, msgFieldError)})
			return
		}
	}

	_, err := orderStatus.ParseOrderStatus(order.OrderStatus.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	if len(order.OrderItemsDTO) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "at least one product must be selected"})
		return
	}

	err = handler.interactor.UpdateOrder(orderId, order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
