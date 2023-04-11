package repo

import (
	"PTOBuilder/internal/character"
	"PTOBuilder/internal/character/model"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"context"
)

type repo struct {
	log    *logging.Logger
	client storage.Client
}

func NewRepo(log *logging.Logger, client storage.Client) character.Repo {
	return &repo{
		log:    log,
		client: client,
	}
}

func (r *repo) CreateRole(ctx context.Context, role *model.Role) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) CreateCharacter(ctx context.Context, character *model.Character) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) GetCharacters(ctx context.Context) *model.Character {
	//TODO implement me
	panic("implement me")
}
