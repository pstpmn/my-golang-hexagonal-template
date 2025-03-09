package port

import (
	"context"

	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/domain"
)

type IUserUseCase interface {
	GetAll(pctx context.Context) ([]domain.User, error)
	GetUser(pctx context.Context, userId string) (*domain.User, error)
}
