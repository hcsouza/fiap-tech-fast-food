package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (handler *customerHandler) GetCustomersHandler(c *gin.Context) {

	actions, err := handler.interactor.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actions)
}
