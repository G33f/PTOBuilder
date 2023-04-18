package character

import (
	"PTOBuilder/internal/character/model"
	"context"
)

type UseCase interface {
	CreateRole(ctx context.Context, role *model.Role) error
	CreateCharacter(ctx context.Context, character *model.Character) error
	GetCharacter(ctx context.Context, characterName string) (*model.Character, error)
}
