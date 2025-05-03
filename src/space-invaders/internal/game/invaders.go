package game

import (
	"fmt"
	"time"
)

const (
	// ENEMY_ROW_COUNT is the number of rows of enemies
	ENEMY_ROW_COUNT = 5
	// ENEMY_COL_COUNT is the number of columns of enemies
	ENEMY_COL_COUNT = 11
	// ENEMY_WIDTH is the width of an enemy
	ENEMY_WIDTH = 40
	// ENEMY_HEIGHT is the height of an enemy
	ENEMY_HEIGHT = 40
	// ENEMY_SPEED is the speed of an enemy
	ENEMY_SPEED = 20
	// ENEMY_DIRECTION is the direction of an enemy
	ENEMY_DIRECTION = 1 // 1 = right, 0 = down, -1 = left
	// ENEMY_PREVIOUS_DIRECTION is the previous direction of an enemy
	ENEMY_PREVIOUS_DIRECTION = 0 // 1 = right, 0 = down, -1 = left
	// ENEMY_START_X is the starting x position of the center of mass of the invasion army
	ENEMY_START_X = 0
	// ENEMY_START_Y is the starting y position of the center of mass of the invasion army
	ENEMY_START_Y = 0
	// ENEMY_X_BUFFER is the horizontal buffer between enemies on the same row
	ENEMY_X_BUFFER = 50
	// ENEMY_Y_BUFFER is the vertical buffer between enemies on the same column
	ENEMY_Y_BUFFER = 50
	// INVADERS_MOVEMENT_FREQUENCY is the frequency of invaders movement
	INVADERS_MOVEMENT_FREQUENCY = 1000 // in milliseconds
)

type Enemy struct {
	Body  Body
	alive bool
}

// TODO: Remove enemies as they are hit by bullets
type Invaders struct {
	prevTime time.Time

	army              [][]*Enemy
	armyHeight        float32
	armyWidth         float32
	possition         Possition
	speed             float32
	direction         int // 1 = right, 0 = down, -1 -= left
	previousDirection int
}

func NewInvaders() Invaders {
	army := make([][]*Enemy, ENEMY_ROW_COUNT)
	armyHeight := float32(ENEMY_ROW_COUNT * ENEMY_Y_BUFFER)
	armyWidth := float32(ENEMY_COL_COUNT * ENEMY_X_BUFFER)
	for row := range army {
		army[row] = make([]*Enemy, ENEMY_COL_COUNT)
		for col := range army[row] {
			army[row][col] = &Enemy{
				Body: Body{
					hitBox: Box{
						width:  ENEMY_WIDTH,
						height: ENEMY_HEIGHT,
					},
					possition: Possition{
						x: float32(ENEMY_START_X + col*ENEMY_X_BUFFER),
						y: float32(ENEMY_START_Y + row*ENEMY_Y_BUFFER),
					},
				},
				alive: true,
			}
		}
	}

	t := time.Now()
	return Invaders{
		prevTime:          t,
		army:              army,
		armyHeight:        armyHeight,
		armyWidth:         armyWidth,
		speed:             ENEMY_SPEED,
		direction:         ENEMY_DIRECTION,
		previousDirection: ENEMY_PREVIOUS_DIRECTION,
		possition:         Possition{x: ENEMY_START_X, y: ENEMY_START_Y},
	}
}

func (i *Invaders) Update() {
	t := time.Now()
	// Move every 1000 milliseconds
	if t.Sub(i.prevTime).Milliseconds() >= INVADERS_MOVEMENT_FREQUENCY {
		i.CheckDirectionChange()
		fmt.Printf("Direction: %d, Previous Direction: %d, Possition: (%.2f, %.2f)\n", i.direction, i.previousDirection, i.possition.x, i.possition.y)
		i.Move()
		i.prevTime = t
	}
}

func (i *Invaders) CheckDirectionChange() {
	if i.direction == 1 && i.possition.x+i.armyWidth >= SCREEN_WIDTH { // we're moving right
		i.direction, i.previousDirection = 0, 1
	} else if i.direction == -1 && i.possition.x <= 0 { // we're moving left
		i.direction, i.previousDirection = 0, -1
	} else if i.direction == 0 { // we're moving down
		if i.previousDirection == 1 { // we were moving right, so next move left
			i.direction, i.previousDirection = -1, 0
		} else if i.previousDirection == -1 { // we were moving left, so next move right
			i.direction, i.previousDirection = 1, 0
		}
	}
}

func (i *Invaders) Move() {
	// Move the enemies in the current direction
	switch i.direction {
	case 1:
		i.moveRight()
	case 0:
		i.moveDown()
	case -1:
		i.moveLeft()
	}
}

func (i *Invaders) moveRight() {
	i.possition.Update(i.speed, 0)
	for row := range i.army {
		for col := range i.army[row] {
			if enemy := i.army[row][col]; enemy.alive {
				enemy.Body.possition.Update(i.speed, 0)
			}
		}
	}
}

func (i *Invaders) moveLeft() {
	i.possition.Update(-i.speed, 0)
	for row := range i.army {
		for col := range i.army[row] {
			if enemy := i.army[row][col]; enemy.alive {
				enemy.Body.possition.Update(-i.speed, 0)
			}
		}
	}
}

func (i *Invaders) moveDown() {
	i.possition.Update(0, i.speed)
	for row := range i.army {
		for col := range i.army[row] {
			if enemy := i.army[row][col]; enemy.alive {
				enemy.Body.possition.Update(0, i.speed)
			}
		}
	}
}

func (i *Invaders) CheckCollision(bullet *Bullet) {
	if !bullet.active {
		return
	}

OuterLoop:
	for row, enemies := range i.army {
		y, h := enemies[0].Body.possition.y, enemies[0].Body.hitBox.height
		// Note: enemy rows are in accending order of y possition going form
		// the top of the screen to the bottom. So, if the bullet has passed the
		// next row (e.g., the top row of enemies), then bullet will also have
		// passed all following enemies.
		bulletY, bulletH := bullet.Body.possition.y, bullet.Body.hitBox.height
		if bulletY+bulletH < y {
			break // bullet has passed the row but did not hit any previous enemies
		}

		if bulletY <= y+h { // equivalent to: topRow <= tobBullet <= bottomRow
			for col, enemy := range enemies {
				x, w := enemy.Body.possition.x, enemy.Body.hitBox.width
				// Note: similar to the above logic of the y possition
				// if the bullet is to the left of the enemy, we do not
				// need to consider it anymore.
				bulletX, bulletW := bullet.Body.possition.x, bullet.Body.hitBox.width
				if bulletX+bulletW < x {
					break OuterLoop // bullet is on th left of enemy and we're moving right
				}

				if enemy.alive && bulletX <= x+w {
					// bullet hit the enemy
					i.army[row][col].alive = false
					bullet.active = false
					break OuterLoop
				}
			}
		}
	}
}
