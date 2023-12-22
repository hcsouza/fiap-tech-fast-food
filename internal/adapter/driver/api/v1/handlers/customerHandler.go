package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/useCases/customer"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/cpf"
)

type customerHandler struct {
	interactor customer.ICustomerUseCase
}

const (
	fieldErrMsg = "Invalid value for field: '%s'"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("IsCpfValid", CpfValidator)
	}
}

func NewCustomerHandler(gRouter *gin.RouterGroup, interactor customer.ICustomerUseCase) {
	handler := &customerHandler{
		interactor: interactor,
	}

	gRouter.GET("/customers", handler.GetCustomersHandler)
	gRouter.POST("/customers", handler.CreateCustomerHandler)

}

func (handler *customerHandler) CreateCustomerHandler(ctx *gin.Context) {

	var createRequest customer.CustomerCreateDTO
	err := ctx.ShouldBindJSON(&createRequest)

	if err != nil {
		var verr validator.ValidationErrors
		var msgFieldError string
		if errors.As(err, &verr) {
			msgFieldError = strings.Split(verr[0].Namespace(), ".")[1]
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(fieldErrMsg, msgFieldError)})
			return
		}
	}

	customer, err := handler.interactor.CreateCustomer(ctx.Request.Context(), createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, customer)
}

func (handler *customerHandler) GetCustomersHandler(ctx *gin.Context) {
	cpf := ctx.Query("cpf")
	params := map[string]string{"cpf": cpf}

	actions, err := handler.interactor.GetCustomer(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, actions)
}

func CpfValidator(fl validator.FieldLevel) bool {
	rawCpf, ok := fl.Field().Interface().(string)
	cpfToValidate := cpf.CPF(rawCpf)

	if ok {
		return cpfToValidate.IsValid()
	}
	return false
}
