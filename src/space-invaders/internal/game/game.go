package game

import (
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
)

type Game struct {
	Player   Player
	Invaders Invaders
}

func InitGame() Game {
	return Game{
		Player:   NewPlayer(),
		Invaders: NewInvaders(),
	}
}

func (game *Game) Update() {
	// Update player and bullet position
	if rl.IsKeyDown(rl.KeyRight) {
		game.Player.MoveRight()
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		game.Player.MoveLeft()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		game.Player.Shoot()
	}

	// Remove bullets that are inactive
	for i := len(game.Player.bullets) - 1; i >= 0; i-- {
		bullet := game.Player.bullets[i]
		bullet.Update()
		if !bullet.active {
			game.Player.bullets = slices.Delete(game.Player.bullets, i, i+1)
		}
	}

	// Update invaders possition
	game.Invaders.Update()

	// Check for bullet collisions with invaders
	for i := len(game.Player.bullets) - 1; i >= 0; i-- {
		bullet := game.Player.bullets[i]
		game.Invaders.CheckCollision(bullet)
	}

	// TODO: Check for invader collisions with player
}

func (game *Game) GetBullets() []*Bullet {
	return game.Player.bullets
}

func (game *Game) GetLivingInvaders() []*Enemy {
	invaders := []*Enemy{}
	for _, row := range game.Invaders.army {
		for _, enemy := range row {
			if enemy.alive {
				invaders = append(invaders, enemy)
			}
		}
	}
	return invaders
}

// TODO: Add game over functionality
