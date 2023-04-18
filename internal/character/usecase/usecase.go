package usecase

import (
	"PTOBuilder/internal/character"
	"PTOBuilder/internal/character/model"
	"PTOBuilder/pkg/logging"
	"context"
)

type useCase struct {
	log  *logging.Logger
	repo character.Repo
}

func NewUseCase(log *logging.Logger, repo character.Repo) character.UseCase {
	return &useCase{
		log:  log,
		repo: repo,
	}
}

func (u *useCase) CreateRole(ctx context.Context, role *model.Role) error {
	err := u.repo.CreateRole(ctx, role)
	if err != nil {
		u.log.Info(err)
	}
	return err
}

func (u *useCase) CreateCharacter(ctx context.Context, character *model.Character) error {
	err := u.repo.CreateCharacter(ctx, character)
	if err != nil {
		u.log.Info(err)
	}
	return err
}

func (u *useCase) GetCharacter(ctx context.Context, characterName string) (*model.Character, error) {
	hero, err := u.repo.GetCharacter(ctx, characterName)
	if err != nil {
		u.log.Info(err)
		return nil, err
	}
	return hero, nil
}
