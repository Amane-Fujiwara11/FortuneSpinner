package repository

import (
	"context"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}