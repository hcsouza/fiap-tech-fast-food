package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
)

type customerHandler struct {
	interactor customer.ICustomerUseCase
}

func NewCustomerHandler(gRouter *gin.RouterGroup, interactor customer.ICustomerUseCase) {
	handler := &customerHandler{
		interactor: interactor,
	}

	gRouter.GET("/customers", handler.GetCustomersHandler)

}

// Get All Customers godoc
// @Summary Get all customers
// @Description Get all customers
// @Tags Customer Routes
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Customer{}
// @Router /api/v1/customers [get]
func (handler *customerHandler) GetCustomersHandler(c *gin.Context) {
	var customers []domain.Customer

	customers, err := handler.interactor.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}
