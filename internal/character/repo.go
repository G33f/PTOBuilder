package character

import (
	"PTOBuilder/internal/character/model"
	"context"
)

type Repo interface {
	CreateRole(ctx context.Context, role *model.Role)
	CreateCharacter(ctx context.Context, character *model.Character)
	GetCharacters(ctx context.Context) *model.Character
}
