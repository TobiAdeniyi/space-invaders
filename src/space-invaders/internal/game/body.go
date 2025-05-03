package game

type Box struct {
	height float32
	width  float32
}

// For a given element in the game, the possition represents
// the uper left corner of the element.
type Possition struct {
	x float32
	y float32
}

func (p *Possition) Update(x, y float32) {
	p.x += x
	p.y += y
}

type Body struct {
	possition Possition
	hitBox  Box
}

func (b *Body) GetWidth() float32 {
	return b.hitBox.width
}

func (b *Body) GetHeight() float32 {
	return b.hitBox.height
}

func (b *Body) GetPossitionX() float32 {
	return b.possition.x
}

func (b *Body) GetPossitionY() float32 {
	return b.possition.y
}
