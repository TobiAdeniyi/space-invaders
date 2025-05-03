package render

import (
	"github.com/TobiAdeniyi/space-invaders/internal/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(game game.Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Draw plauer
	rl.DrawRectangle(
		int32(game.Player.Body.GetPossitionX()),
		int32(game.Player.Body.GetPossitionY()),
		int32(game.Player.Body.GetWidth()),
		int32(game.Player.Body.GetHeight()),
		rl.Green,
	)

	// Draw bullets
	for _, bullet := range game.GetBullets() {
		rl.DrawRectangle(
			int32(bullet.Body.GetPossitionX()),
			int32(bullet.Body.GetPossitionY()),
			int32(bullet.Body.GetWidth()),
			int32(bullet.Body.GetHeight()),
			rl.White,
		)
	}

	// Draw invaders
	for _, invaders := range game.GetLivingInvaders() {
		rl.DrawRectangle(
			int32(invaders.Body.GetPossitionX()),
			int32(invaders.Body.GetPossitionY()),
			int32(invaders.Body.GetWidth()),
			int32(invaders.Body.GetHeight()),
			rl.White,
		)
	}

	rl.EndDrawing()
}
