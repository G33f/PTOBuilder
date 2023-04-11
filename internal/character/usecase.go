package character

import "PTOBuilder/internal/character/model"

type UseCase interface {
	CreateRole(role *model.Role)
	CreateCharacter(character *model.Character)
	GetCharacters() *model.Character
}
