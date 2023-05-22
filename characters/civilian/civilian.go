package civilian

import "werewolves/characters"

type Civilian struct {
	status characters.CharacterStatus
	typ    characters.CharacterType
}

func New() characters.ICharacter {
	return &Civilian{status: characters.Alive, typ: characters.Good}
}

func (c *Civilian) GetStatus() characters.CharacterStatus {
	return c.status
}

func (c *Civilian) GetType() characters.CharacterType {
	return c.typ
}
