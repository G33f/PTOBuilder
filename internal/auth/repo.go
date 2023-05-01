package auth

import (
	"PTOBuilder/internal/auth/model"
	"context"
)

type Repo interface {
	AddUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, user *model.User) error
}
