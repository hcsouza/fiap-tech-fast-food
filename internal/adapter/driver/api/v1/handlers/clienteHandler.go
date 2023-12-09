package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/cliente"
)

type clienteHandler struct {
	interactor cliente.IClienteUseCase
}

func NewClienteHandler(gRouter *gin.RouterGroup, interactor cliente.IClienteUseCase) {
	handler := &clienteHandler{
		interactor: interactor,
	}

	gRouter.GET("/clientes", handler.GetClientesHandler)

}

func (handler *clienteHandler) GetClientesHandler(c *gin.Context) {

	actions, err := handler.interactor.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actions)
}
