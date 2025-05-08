package game

const (
	PLAYER_SPEED       = 10
	PLAYER_WIDTH       = 80
	PLAYER_HEIGHT      = 20
	PLAYER_POSSITION_X = 400
	// PLAYER_POSSITION_Y = SCREEN_HEIGHT - PLAYER_HEIGHT
)

type Player struct {
	alive   bool
	Body    Body
	bullets []*Bullet
	playerXBoundary float32
}

func NewPlayer(screenWidth, screenHeight float32) Player {
	return Player{
		alive: true,
		Body: Body{
			hitBox: Box{
				width:  PLAYER_WIDTH,
				height: PLAYER_HEIGHT,
			},
			possition: Possition{
				x: PLAYER_POSSITION_X,
				y: screenHeight - PLAYER_HEIGHT,
			},
		},
		bullets: []*Bullet{},
		playerXBoundary: screenWidth,
	}
}

func (p *Player) MoveLeft() {
	if p.Body.possition.x > 0 {
		p.Body.possition.Update(-PLAYER_SPEED, 0)
	}
}

func (p *Player) MoveRight() {
	if p.Body.possition.x+p.Body.hitBox.width < p.playerXBoundary {
		p.Body.possition.Update(PLAYER_SPEED, 0)
	}
}

func (p *Player) Shoot() {
	bullet := NewBullet(
		p.Body.possition.x + float32(PLAYER_WIDTH/2) - float32(BULLET_WIDTH/2),
		p.Body.possition.y,
	)
	p.bullets = append(p.bullets, bullet)
}
