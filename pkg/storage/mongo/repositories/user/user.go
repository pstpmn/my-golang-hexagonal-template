package userRepository

import (
	"context"

	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/domain"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	r struct {
		db    *mongo.Client
		cache port.ICache
	}
)

func (r *r) FindAll(pctx context.Context) ([]domain.User, error) {
	users := []domain.User{
		{
			ID:       "8e6c6952-fcfe-4ffd-86b4-ce23c37538f0",
			Name:     "Mr.Math Matic",
			Email:    "math@fake.com",
			IsActive: true,
		},
		{
			ID:       "294bd4ab-4e9c-482f-ba73-e2ab3d20efe7",
			Name:     "Mr.Eloy Musk",
			Email:    "eloy@fake.com",
			IsActive: true,
		},
	}
	return users, nil
}

func (r *r) FindOneById(pctx context.Context, userId string) (*domain.User, error) {
	return nil, domain.ErrUserNotFound
}

func NewUserRepo(db *mongo.Client, cache port.ICache) port.IUserRepo {
	return &r{
		db:    db,
		cache: cache,
	}
}
