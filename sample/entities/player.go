package entities

type Player struct {
	X int
	Y int
}

func NewPlayer(x, y int) *Player {
	return &Player{x, y}
}

func (p *Player) Draw() (int, int, string) {
	return p.X, p.Y, "P"
}
