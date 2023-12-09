package cliente

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type clienteUseCase struct{}

func NewClienteUseCase() IClienteUseCase {
	return &clienteUseCase{}
}

func (interactor *clienteUseCase) GetAll(logContext context.Context) ([]domain.Cliente, error) {
	return []domain.Cliente{
		{Nome: "Cliente 1", Email: "cliente@mail.com", CPF: "394.671.960-00"},
		{Nome: "Cliente 2", Email: "cliente2@mail.com", CPF: "963.953.450-10"},
	}, nil
}
