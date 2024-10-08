package entities

type FallingBall struct {
	x       int
	y       int
	gravity int
	speedY  int
}

func NewFallingBall(x, y, gravity int) *FallingBall {
	return &FallingBall{x, y, gravity, 0}
}

func (f *FallingBall) Draw() (int, int, string) {
	// update the speed
	f.speedY += f.gravity

	// update the position
	f.y += f.speedY

	return f.x, f.y, "O"
}
