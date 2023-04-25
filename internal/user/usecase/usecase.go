package usecase

import (
	"PTOBuilder/internal/user"
	"PTOBuilder/internal/user/model"
	"PTOBuilder/pkg/logging"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type useCase struct {
	log  *logging.Logger
	repo user.Repo
}

func NewUseCase(log *logging.Logger, repo user.Repo) user.UseCase {
	return &useCase{
		log:  log,
		repo: repo,
	}
}

func (u *useCase) UserRegistration(ctx context.Context, user *model.User) error {
	if check := u.isUserExist(ctx, user); check == true {
		return fmt.Errorf("user already exist")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Info(err)
		return err
	}
	user.Password = string(hash)
	err = u.repo.UserRegistration(ctx, user)
	if err != nil {
		u.log.Info(err)
		return err
	}
	return nil
}

func (u *useCase) isUserExist(ctx context.Context, user *model.User) bool {
	err := u.repo.GetUser(ctx, user)
	if err == pgx.ErrNoRows {
		return false
	}
	return true
}

func (u *useCase) CheckUser(ctx context.Context, user *model.User) error {
	panic("some")
}
