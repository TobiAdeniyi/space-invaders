package render

import (
	"github.com/TobiAdeniyi/space-invaders/internal/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(g game.Game, assets Assets) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Draw plauer
	rl.DrawRectangle(
		int32(g.Player.Body.GetPossitionX()),
		int32(g.Player.Body.GetPossitionY()),
		int32(g.Player.Body.GetWidth()),
		int32(g.Player.Body.GetHeight()),
		rl.Green,
	)

	// Draw bullets
	for _, bullet := range g.GetBullets() {
		rl.DrawRectangle(
			int32(bullet.Body.GetPossitionX()),
			int32(bullet.Body.GetPossitionY()),
			int32(bullet.Body.GetWidth()),
			int32(bullet.Body.GetHeight()),
			rl.White,
		)
	}

	// Draw invaders
	for _, enemy := range g.GetInvadersToDraw() {
		// rl.DrawRectangle(
		// 	int32(enemy.Body.GetPossitionX()),
		// 	int32(enemy.Body.GetPossitionY()),
		// 	int32(enemy.Body.GetWidth()),
		// 	int32(enemy.Body.GetHeight()),
		// 	rl.White,
		// )

		if enemy.DrawDeath {
			// TODO: draw death animation
			enemy.DrawDeath = false
			return
		}

		n := 0
		if g.Invaders.Frame {
			n = 1
		}

		texture, h, w := new(rl.Texture2D), enemy.Body.GetHeight(), enemy.Body.GetWidth()
		switch enemy.Type {
		case game.SMALL_ENEMY:
			texture = &assets.smallEnemyTexture[n]
		case game.MEDIUM_ENEMY:
			texture = &assets.mediumEnemyTexture[n]
			if g.Invaders.Frame {
				w = game.MEDIUM_ENEMY_WIDTH_2
			}
		case game.LARGE_ENEMY:
			texture = &assets.largeEnemyTexture[n]
			if g.Invaders.Frame {
				h = game.LARRGE_ENEMY_HEIGHT_2
			}
		}

		rl.DrawTextureRec(
			*texture,
			rl.Rectangle{
				X:      0, //enemy.Body.GetPossitionX(),
				Y:      0, //enemy.Body.GetPossitionY(),
				Width:  w,
				Height: h,
			},
			rl.Vector2{
				X: enemy.Body.GetPossitionX(),
				Y: enemy.Body.GetPossitionY(),
			},
			rl.White,
		)
	}

	rl.EndDrawing()
}

func DrawGameOver(g game.Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Draw plauer
	rl.DrawRectangle(
		int32(g.Player.Body.GetPossitionX()),
		int32(g.Player.Body.GetPossitionY()),
		int32(g.Player.Body.GetWidth()),
		int32(g.Player.Body.GetHeight()),
		rl.Green,
	)

	// Draw Game Over Screen
	rl.DrawText("Game Over", int32(g.ScreanWidth/2-50), int32(g.ScreanHeight/4), 20, rl.Red)

	rl.EndDrawing()
}
