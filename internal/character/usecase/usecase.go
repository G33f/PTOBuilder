package usecase

import (
	"PTOBuilder/internal/character"
	"PTOBuilder/internal/character/model"
	"PTOBuilder/pkg/logging"
)

type useCase struct {
	log  *logging.Logger
	repo *character.Repo
}

func NewUseCase(log *logging.Logger, repo *character.Repo) character.UseCase {
	return &useCase{
		log:  log,
		repo: repo,
	}
}

func (u useCase) CreateRole(role *model.Role) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) CreateCharacter(character *model.Character) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) GetCharacters() *model.Character {
	//TODO implement me
	panic("implement me")
}
