package witch

import "werewolves/characters"

type Witch struct {
	status characters.CharacterStatus
	typ    characters.CharacterType
}

func New() characters.ICharacter {
	return &Witch{status: characters.Alive, typ: characters.Good}
}

func (w *Witch) GetStatus() characters.CharacterStatus {
	return w.status
}

func (w *Witch) GetType() characters.CharacterType {
	return w.typ
}
