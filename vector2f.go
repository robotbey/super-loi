package main

type Vector2f struct {
	X float64
	Y float64
}

func NewVector2f(x float64, y float64) *Vector2f {
	return &Vector2f{
		X: x,
		Y: y,
	}
}
