package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	width           int
	height          int
	loi             *Loi
	backgroundImage *ebiten.Image
}

func NewGame(width, height int) (*Game, error) {
	loi := NewLoi(*NewVector2f(0, float64(height-200)), *NewVector2f(10, 0))
	game := &Game{width, height, loi, nil}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Super Loi v" + VERSION)

	f, err := os.Open("./data/sprites/character/walk1.png")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	walk1 := ebiten.NewImageFromImage(img)
	loi.animation.AddImage(walk1)

	f2, err := os.Open("./data/sprites/character/walk2.png")
	if err != nil {
		return nil, err
	}
	defer f2.Close()
	img2, _, err := image.Decode(f2)
	if err != nil {
		return nil, err
	}
	walk2 := ebiten.NewImageFromImage(img2)
	loi.animation.AddImage(walk2)

	f3, err := os.Open("./data/sprites/background/level1.png")
	if err != nil {
		return nil, err
	}
	defer f3.Close()
	img, _, err = image.Decode(f3)
	if err != nil {
		return nil, err
	}
	game.backgroundImage = ebiten.NewImageFromImage(img)

	return game, nil
}

func (g *Game) Width() int {
	return g.width
}

func (g *Game) Height() int {
	return g.height
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(g.backgroundImage, op)
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(g.loi.Position.X, g.loi.Position.Y)
	screen.DrawImage(g.loi.animation.GetImage(), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

const SPEED = 1

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.loi.Velocity.X = -4 * SPEED
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.loi.Velocity.X = 4 * SPEED
	}
	g.loi.Update()
	return nil
}

func (g *Game) Run() error {
	err := ebiten.RunGame(g)
	return err
}
