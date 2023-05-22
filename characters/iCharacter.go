package characters

type CharacterStatus int

const (
	Alive CharacterStatus = iota
	Die
)

type CharacterType int

const (
	Good CharacterType = iota
	Bad
)

type ICharacter interface {
	GetStatus() CharacterStatus

	GetType() CharacterType
}
