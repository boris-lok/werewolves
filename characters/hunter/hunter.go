package hunter

import (
	"werewolves/characters"
)

type Hunter struct {
	status characters.CharacterStatus
	typ    characters.CharacterType
}

func New() characters.ICharacter {
	return &Hunter{status: characters.Alive, typ: characters.Good}
}

func (h *Hunter) GetStatus() characters.CharacterStatus {
	return h.status
}

func (h *Hunter) GetType() characters.CharacterType {
	return h.typ
}
