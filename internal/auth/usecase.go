package auth

import (
	"PTOBuilder/internal/auth/model"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, user *model.User) error
	SignIn(ctx context.Context, user *model.User) (string, error)
	ValidToken(tokenString string) error
}
