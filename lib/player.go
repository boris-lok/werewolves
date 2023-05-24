package lib

type UserId string

type PlayerStatus int

const (
	Alive PlayerStatus = iota
	Die
)

type Role int

const (
	Wolf Role = iota
	Witch
	Seer
	Hunter
	Civilian
)

type Player struct {
	Id     UserId
	Name   string
	Role   Role
	Status PlayerStatus
}

// New create a player
func New(id UserId, name string, role Role) *Player {
	return &Player{Id: id, Name: name, Status: Alive}
}
