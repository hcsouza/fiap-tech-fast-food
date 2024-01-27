package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"
	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
	coreErrors "github.com/hcsouza/fiap-tech-fast-food/internal/core/errors"
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

	gRouter.GET("/customer", handler.GetCustomerHandler)
	gRouter.POST("/customer", handler.CreateCustomerHandler)

}

// Create Customer godoc
// @Summary Create a new customer
// @Description Create a new customer
// @Tags Customer Routes
// @Param        data    body     customer.CustomerCreateDTO  true  "Customer information"
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Customer{}
// @Router /api/v1/customer [post]
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

	if errors.Is(err, coreErrors.ErrDuplicatedKey) {
		ctx.JSON(http.StatusConflict, gin.H{"error": "customer already exists"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, customer)
}

// Get Customer godoc
// @Summary Get customer by CPF
// @Description Get customer by CPF
// @Tags Customer Routes
// @Param        cpf    query     string  true  "19119119100"
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.Customer{}
// @Router /api/v1/customer [get]
func (handler *customerHandler) GetCustomerHandler(ctx *gin.Context) {
	cpf := ctx.Query("cpf")
	params := map[string]string{"cpf": cpf}
	var customer *domain.Customer // Only to swaggo doc

	actions, err := handler.interactor.GetCustomer(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer = actions

	ctx.JSON(http.StatusOK, customer)
}

func CpfValidator(fl validator.FieldLevel) bool {
	rawCpf, ok := fl.Field().Interface().(string)
	cpfToValidate := cpf.CPF(rawCpf)

	if ok {
		return cpfToValidate.IsValid()
	}
	return false
}
