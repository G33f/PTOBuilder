package user

import (
	"PTOBuilder/internal/user/model"
	"context"
)

type UseCase interface {
	UserRegistration(ctx context.Context, user *model.User) error
	CheckUser(ctx context.Context, user *model.User) error
}
