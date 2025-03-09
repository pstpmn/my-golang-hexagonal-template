package port

import (
	"context"

	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/domain"
)

type IUserRepo interface {
	FindAll(pctx context.Context) ([]domain.User, error)
	FindOneById(pctx context.Context, userId string) (*domain.User, error)
}
