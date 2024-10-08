package entities

import "strings"

type Floor struct {
	y     int
	width int
}

func NewFloor(positionY int, width int) *Floor {
	return &Floor{positionY, width}
}

func (f *Floor) Draw() (int, int, string) {
	return 1, f.y, strings.Repeat("F", f.width)
}
