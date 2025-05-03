package main // package space-invaders

import (
	"fmt"

	"github.com/TobiAdeniyi/space-invaders/internal/game"
	"github.com/TobiAdeniyi/space-invaders/internal/render"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	g := game.InitGame()
	rl.InitWindow(game.SCREEN_WIDTH, game.SCREEN_HEIGHT, "Space Invaders")

	for !rl.WindowShouldClose() {
		// Update the game state
		g.Update()

		// Draw the game
		render.Draw(g)
	}

	fmt.Println("Game Over!")
}
