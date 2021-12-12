package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	frames        []*ebiten.Image
	index         int
	speed         float64
	lastFrameTime float64
}

func (a *Animation) GetImage() *ebiten.Image {
	return a.frames[a.index]
}

func (a *Animation) AddImage(img *ebiten.Image) {
	a.frames = append(a.frames, img)
}

func (a *Animation) Advance() {
	currentTime := float64(time.Now().UnixMilli())
	if currentTime-a.lastFrameTime > 1000*a.speed {
		a.index++
		if a.index >= len(a.frames) {
			a.index = 0
		}
		a.lastFrameTime = currentTime
	}
}

func NewAnimation() *Animation {
	return &Animation{
		frames:        make([]*ebiten.Image, 0),
		index:         0,
		speed:         0.2,
		lastFrameTime: float64(time.Now().UnixMilli()),
	}
}
