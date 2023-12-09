package cliente

import (
	"context"

	"github.com/hcsouza/fiap-tech-fast-food/internal/core/domain"
)

type IClienteUseCase interface {
	GetAll(context.Context) ([]domain.Cliente, error)
}
