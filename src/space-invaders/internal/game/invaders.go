package game

import (
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
	ENEMY_X_BUFFER = 70
	// ENEMY_Y_BUFFER is the vertical buffer between enemies on the same column
	ENEMY_Y_BUFFER = 50
	// INVADERS_MOVEMENT_FREQUENCY is the frequency of invaders movement
	INVADERS_MOVEMENT_FREQUENCY = 1000 // in milliseconds
)

type EnemyType int

const (
	SMALL_ENEMY EnemyType = iota
	MEDIUM_ENEMY
	LARGE_ENEMY
)

const ( // Pixel sixe * 5
	SMALL_ENEMY_WIDTH  = 40
	MEDIUM_ENEMY_WIDTH = 45
	LARGE_ENEMY_WIDTH  = 55
	MEDIUM_ENEMY_WIDTH_2 = 55
	LARRGE_ENEMY_HEIGHT_2 = 45
)

type Enemy struct {
	Body      Body
	alive     bool
	Type      EnemyType
	DrawDeath bool
}

type Fleet struct {
	count      int
	yPossition float32
	members    [ENEMY_COL_COUNT]*Enemy
}

// TODO: Remove enemies as they are hit by bullets
type Invaders struct {
	fleetCount    int
	army          [ENEMY_ROW_COUNT]Fleet
	armyHeight    float32
	armyWidth     float32
	armyXBoundary float32

	prevTime time.Time
	Frame    bool

	possition         Possition
	speed             float32
	direction         int // 1 = right, 0 = down, -1 -= left
	previousDirection int
}

func getEnemyType(row int) EnemyType {
	switch row {
	case 1, 2:
		return MEDIUM_ENEMY
	case 3, 4:
		return LARGE_ENEMY
	default:
		return SMALL_ENEMY

	}
}

func getEnemyWidth(row int) float32 {
	switch row {
	case 1, 2:
		return MEDIUM_ENEMY_WIDTH
	case 3, 4:
		return LARGE_ENEMY_WIDTH
	default:
		return SMALL_ENEMY_WIDTH
	}
}

func getEnemyStartX(row int) float32 {
	switch row {
	case 1, 2:
		return ENEMY_START_X + (LARGE_ENEMY_WIDTH-MEDIUM_ENEMY_WIDTH)/2
	case 3, 4:
		return ENEMY_START_X
	default:
		return ENEMY_START_X + (LARGE_ENEMY_WIDTH-SMALL_ENEMY_WIDTH)/2
	}
}

func NewInvaders(xBoundary float32) Invaders {
	army := [ENEMY_ROW_COUNT]Fleet{}
	armyHeight := float32(ENEMY_ROW_COUNT * ENEMY_Y_BUFFER)
	armyWidth := float32(ENEMY_COL_COUNT * ENEMY_X_BUFFER)
	for row := range army {
		members := [ENEMY_COL_COUNT]*Enemy{}
		enemy_start_x := getEnemyStartX(row)
		for col := range members {
			members[col] = &Enemy{
				Body: Body{
					hitBox: Box{
						width:  getEnemyWidth(row),
						height: ENEMY_HEIGHT,
					},
					possition: Possition{
						x: enemy_start_x + float32(col*ENEMY_X_BUFFER),
						y: float32(ENEMY_START_Y + row*ENEMY_Y_BUFFER),
					},
				},
				alive:     true,
				Type:      getEnemyType(row),
				DrawDeath: false,
			}
		}

		army[row] = Fleet{
			count:   ENEMY_COL_COUNT,
			members: members,
		}
	}

	t := time.Now()
	return Invaders{
		prevTime:          t,
		Frame:             false,
		army:              army,
		fleetCount:        ENEMY_ROW_COUNT,
		armyHeight:        armyHeight,
		armyWidth:         armyWidth,
		armyXBoundary:     xBoundary,
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
		i.Move()
		i.prevTime = t
		i.Frame = !i.Frame
	}
}

func (i *Invaders) CheckDirectionChange() {
	if i.direction == 1 && i.possition.x+i.armyWidth >= i.armyXBoundary { // we're moving right
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
	switch i.direction {
	case 1: // moving right
		i.moveDirection(i.speed, 0)
	case 0: // moving down
		i.moveDirection(0, i.speed)
		// move the fleet down
		for row := range i.army {
			fleet := &i.army[row]
			fleet.yPossition += i.speed
		}
	case -1: // moving left
		i.moveDirection(-i.speed, 0)
	}
}

func (i *Invaders) moveDirection(x, y float32) {
	i.possition.Update(x, y)
	for _, fleet := range i.army {
		for _, member := range fleet.members {
			if member != nil && member.alive {
				member.Body.possition.Update(x, y)
			}
		}
	}
}

func (i *Invaders) GetLowestInvader() *Enemy {
	for row := ENEMY_ROW_COUNT - 1; row >= 0; row-- {
		fleet := i.army[row]
		if fleet.count > 0 {
			for col := ENEMY_COL_COUNT - 1; col >= 0; col-- {
				enemy := fleet.members[col]
				if enemy.alive {
					return enemy
				}
			}
		}
	}
	return nil
}

func (i *Invaders) CheckCollision(bullet *Bullet) {
	if !bullet.active {
		return
	}

OuterLoop:
	for row, fleet := range i.army {
		if fleet.count == 0 {
			continue // no enemies in this fleet
		}

		y := i.possition.y + float32(row*ENEMY_Y_BUFFER)
		// Note: enemy rows are in accending order of y possition going form
		// the top of the screen to the bottom. So, if the bullet has passed the
		// next row (e.g., the top row/fleet), then bullet will also have
		// passed all following fleet.
		bulletY := bullet.Body.possition.y
		if bulletY+BULLET_HEIGHT < y {
			break // bullet has passed the row but did not hit any previous fleet
		}

		if bulletY <= y+ENEMY_HEIGHT { // equivalent to: topRow <= tobBullet <= bottomRow
			for _, enemy := range fleet.members {
				x := enemy.Body.possition.x //i.possition.x + float32(col*ENEMY_X_BUFFER)
				// Note: similar to the above logic of the y possition
				// if the bullet is to the left of the enemy, we do not
				// need to consider it anymore.
				bulletX := bullet.Body.possition.x
				if bulletX+BULLET_WIDTH < x {
					break OuterLoop // bullet is on th left of enemy and we're moving right
				}

				if enemy.alive && bulletX <= x+enemy.Body.hitBox.width {
					// bullet hit the enemy
					bullet.active = false
					enemy.alive = false
					enemy.DrawDeath = true

					// adjust fleet and enemy count
					i.army[row].count = fleet.count - 1
					if i.army[row].count == 0 {
						i.fleetCount = i.fleetCount - 1
					}

					break OuterLoop
				}
			}
		}
	}
}
