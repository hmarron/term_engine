package main

import (
	"time"

	engine "github.com/hmarron/term_engine"
	"github.com/hmarron/term_engine/sample/entities"
)

func main() {
	game := engine.NewEngine(time.Millisecond * 20)
	maxX, maxY := game.GetMaxPos()

	floor := entities.NewFloor(maxY, maxX)
	fallingBall := entities.NewFallingBall(maxX/2, 1, 1)
	player := entities.NewPlayer(1, 1)

	game.AddDrawable(floor)
	game.AddDrawable(fallingBall)
	game.AddDrawable(player)

	game.SetKeyFunctions(map[rune]func(){
		'A': func() {
			player.Y--
		},
		'B': func() {
			player.Y++
		},
		'C': func() {
			player.X++
		},
		'D': func() {
			player.X--
		},
	})

	game.Start()
}
