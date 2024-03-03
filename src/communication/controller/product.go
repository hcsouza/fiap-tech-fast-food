package controller

import (
	"github.com/hcsouza/fiap-tech-fast-food/src/common/interfaces"
	"github.com/hcsouza/fiap-tech-fast-food/src/communication/gateway"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/entity"
	"github.com/hcsouza/fiap-tech-fast-food/src/core/usecase"
	valueobject "github.com/hcsouza/fiap-tech-fast-food/src/core/valueObject"
)

type productController struct {
	useCase interfaces.ProductUseCase
}

func NewProductController(datasource interfaces.DatabaseSource) *productController {
	gateway := gateway.NewProductGateway(datasource)
	return &productController{
		useCase: usecase.NewProductUseCase(gateway),
	}
}

func (pc *productController) GetAll() ([]entity.Product, error) {
	return pc.useCase.GetAll()
}

func (pc *productController) GetByCategory(category valueobject.Category) ([]entity.Product, error) {
	return pc.useCase.GetByCategory(category)
}

func (pc *productController) Create(product *entity.Product) error {
	return pc.useCase.Create(product)
}

func (pc *productController) Update(productId string, product *entity.Product) error {
	return pc.useCase.Update(productId, product)
}

func (pc *productController) Delete(productId string) error {
	return pc.useCase.Delete(productId)
}

func (pc *productController) FindById(id string) (*entity.Product, error) {
	return pc.useCase.FindById(id)
}
