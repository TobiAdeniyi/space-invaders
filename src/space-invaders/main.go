package main // package space-invaders

import (
	"fmt"

	"github.com/TobiAdeniyi/space-invaders/internal/game"
	"github.com/TobiAdeniyi/space-invaders/internal/render"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
)

func main() {
	g := game.InitGame(SCREEN_WIDTH, SCREEN_HEIGHT)
	rl.SetTargetFPS(60)
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Space Invaders")
	assets := render.LoadAssets()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyEscape) {
			break
		}

		if g.IsGameOver() {
			render.DrawGameOver(g)
		} else {
			// Update the game state
			g.Update()
			// Draw the game
			render.Draw(g, assets)
		}
	}

	fmt.Println("Closing Game!")
}


// TODO:
// 1. Change color of the enemies texture to white
// 2. Draw disapering enemies texture when they are killed
// 3. Make enemy body size varies based on game
// 4. Increase the speed of the enemies over time
// 5. Add bulders as shown in game
// 6. Update game over logic
