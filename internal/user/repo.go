package user

import (
	"PTOBuilder/internal/user/model"
	"context"
)

type Repo interface {
	UserRegistration(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, user *model.User) error
}
