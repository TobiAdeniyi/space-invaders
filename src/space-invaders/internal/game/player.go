package game

const (
	PLAYER_SPEED       = .1
	PLAYER_WIDTH       = 80
	PLAYER_HEIGHT      = 20
	PLAYER_POSSITION_X = 400
	PLAYER_POSSITION_Y = SCREEN_HEIGHT - PLAYER_HEIGHT
)

type Player struct {
	Body    Body
	bullets []*Bullet
}

func NewPlayer() Player {
	return Player{
		Body: Body{
			hitBox: Box{
				width:  PLAYER_WIDTH,
				height: PLAYER_HEIGHT,
			},
			possition: Possition{
				x: PLAYER_POSSITION_X,
				y: PLAYER_POSSITION_Y,
			},
		},
		bullets: []*Bullet{},
	}
}

func (p *Player) MoveLeft() {
	if p.Body.possition.x > 0 {
		p.Body.possition.Update(-PLAYER_SPEED, 0)
	}
}

func (p *Player) MoveRight() {
	if p.Body.possition.x+p.Body.hitBox.width < SCREEN_WIDTH {
		p.Body.possition.Update(PLAYER_SPEED, 0)
	}
}

func (p *Player) Shoot() {
	bullet := NewBullet(
		p.Body.possition.x,
		p.Body.possition.y,
	)
	p.bullets = append(p.bullets, bullet)
}
