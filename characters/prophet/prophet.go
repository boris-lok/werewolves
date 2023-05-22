package prophet

import "werewolves/characters"

type Prophet struct {
	status characters.CharacterStatus
	typ    characters.CharacterType
}

func New() characters.ICharacter {
	return &Prophet{status: characters.Alive, typ: characters.Good}
}

func (p *Prophet) GetStatus() characters.CharacterStatus {
	return p.status
}

func (p *Prophet) GetType() characters.CharacterType {
	return p.typ
}
