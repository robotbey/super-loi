package main

type Loi struct {
	Position  Vector2f
	animation *Animation
	Velocity  Vector2f
}

func (loi *Loi) Update() {
	loi.Position.X += loi.Velocity.X
	if loi.Velocity.X > 0 {
		loi.Velocity.X--
	}
	loi.Position.Y += loi.Velocity.Y
	if loi.Velocity.Y > 0 {
		loi.Velocity.Y--
	}
	if loi.Velocity.X > 0 || loi.Velocity.Y > 0 {
		loi.animation.Advance()
	}
}

func NewLoi(pos Vector2f, v Vector2f) *Loi {
	return &Loi{
		Position:  pos,
		animation: NewAnimation(),
		Velocity:  v,
	}
}
