package game

import (
	"fmt"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player       Player
	Invaders     Invaders
	ScreanWidth  int
	ScreanHeight int
}

func InitGame(screenWidth, screenHeight int) Game {
	return Game{
		Player: NewPlayer(
			float32(screenWidth),
			float32(screenHeight),
		),
		Invaders:     NewInvaders(float32(screenWidth)),
		ScreanWidth:  screenWidth,
		ScreanHeight: screenHeight,
	}
}

func (g *Game) Update() {
	if g.Invaders.fleetCount == 0 {
		// Player has killed all invaders
		fmt.Println("Player has killed all invaders")
		g.GameOver()
		return
	}

	// Update player and bullet position
	if rl.IsKeyDown(rl.KeyRight) {
		g.Player.MoveRight()
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		g.Player.MoveLeft()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		g.Player.Shoot()
	}

	// Remove bullets that are inactive
	for i := len(g.Player.bullets) - 1; i >= 0; i-- {
		bullet := g.Player.bullets[i]
		bullet.Update()
		if !bullet.active {
			g.Player.bullets = slices.Delete(g.Player.bullets, i, i+1)
		}
	}

	// Update invaders possition
	g.Invaders.Update()

	// Check for bullet collisions with invaders
	for i := len(g.Player.bullets) - 1; i >= 0; i-- {
		bullet := g.Player.bullets[i]
		g.Invaders.CheckCollision(bullet)
	}

	// Check player collision once invaders are low enough to colid with player
	if g.Invaders.possition.y+ENEMY_HEIGHT > g.Player.Body.possition.y {
		if lowestInvader := g.Invaders.GetLowestInvader(); lowestInvader == nil {
			// Player has killed all invaders
			fmt.Println("Player has killed all invaders")
			g.GameOver()
		} else if lowestInvader.Body.possition.x > g.Player.Body.possition.y {
			// Player has been hit by an invader
			fmt.Printf(
				"Player has been hit by an invader: (%.2f, %.2f)\n",
				lowestInvader.Body.possition.x,
				lowestInvader.Body.possition.y,
			)
			g.GameOver()
		}
	}
}

func (g *Game) GetBullets() []*Bullet {
	return g.Player.bullets
}

func (g *Game) GetInvadersToDraw() []*Enemy {
	invaders := []*Enemy{}
	for _, fleet := range g.Invaders.army {
		for _, enemy := range fleet.members {
			if enemy.alive || enemy.DrawDeath {
				invaders = append(invaders, enemy)
			}
		}
	}
	return invaders
}

// TODO: Add game over functionality
func (g *Game) GameOver() {
	fmt.Println("Game Over")

	if g.Player.alive {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
}

func (g *Game) IsGameOver() bool {
	return !g.Player.alive || g.Invaders.fleetCount == 0
}
