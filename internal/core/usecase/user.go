package usecase

import (
	"context"

	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/domain"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
)

type (
	user struct {
		userRepo port.IUserRepo
	}
)

func (u *user) GetAll(pctx context.Context) ([]domain.User, error) {
	result, err := u.userRepo.FindAll(pctx)
	return result, err
}

func (u *user) GetUser(pctx context.Context, userId string) (*domain.User, error) {
	result, err := u.userRepo.FindOneById(pctx, userId)
	return result, err
}

func NewUserUseCase(userRepo port.IUserRepo) port.IUserUseCase {
	return &user{
		userRepo: userRepo,
	}
}
