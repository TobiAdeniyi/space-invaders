package game

const (
	// BULLET_SPPED is the fixed speed of the bullet
	BULLET_SPPED = 1
	// BULLET_WIDTH is the width of the bullet
	BULLET_WIDTH = 5
	// BULLET_HEIGHT is the height of the bullet
	BULLET_HEIGHT = 10
)

type Bullet struct {
	Body   Body
	active bool
}

func NewBullet(x, y float32) *Bullet {
	return &Bullet{
		Body: Body{
			possition: Possition{
				x: x,
				y: y,
			},
			hitBox: Box{
				height: BULLET_HEIGHT,
				width:  BULLET_WIDTH,
			},
		},
		active: true,
	}
}

func (b *Bullet) updatePossition() {
	b.Body.possition.Update(0, -BULLET_SPPED)
}

func (b *Bullet) isOutOfBounds() bool {
	return b.Body.possition.y < 0
}

func (b *Bullet) Update() {
	b.updatePossition()
	if b.isOutOfBounds() {
		b.active = false
	}
}
