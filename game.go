package main

import (
	"image/color"

	"github.com/KidPudel/noodle_go/entities"
	"github.com/KidPudel/noodle_go/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Noodle             *entities.Noodle
	Flour              *entities.Flour
	VerticalGridLine   *ebiten.Image
	HorizontalGridLine *ebiten.Image

	Score int
}

func InitGame() *Game {
	noodle := entities.CreateNoodle()
	flour := entities.SpawnFlour()
	vertical := ebiten.NewImage(util.NoodleHeadSize/10, util.ScreenHeight)
	horizontal := ebiten.NewImage(util.ScreenWidth, util.NoodleHeadSize/10)
	game := &Game{Noodle: noodle, Flour: flour, VerticalGridLine: vertical, HorizontalGridLine: horizontal}
	return game
}

// game's loop basic functionality

func (g *Game) Update() error {
	g.Noodle.Update(g.Score)
	g.Flour.Update(g.Noodle.Pos, &g.Score)

	return nil
}

func (g *Game) DrawGrid(screen *ebiten.Image) {
	g.VerticalGridLine.Fill(color.RGBA{211, 193, 191, 255})
	g.HorizontalGridLine.Fill(color.RGBA{211, 193, 191, 255})
	for i := range util.ScreenWidth / util.NoodleHeadSize {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(util.NoodleHeadSize*i), 0)
		screen.DrawImage(g.VerticalGridLine, op)
	}
	for i := range util.ScreenHeight / util.NoodleHeadSize {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(util.NoodleHeadSize*i))
		screen.DrawImage(g.HorizontalGridLine, op)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{254, 230, 228, 255})
	g.DrawGrid(screen)
	g.Flour.Draw(screen)
	g.Noodle.Draw(screen)
}

func (g *Game) Layout(width, height int) (logicalWidth, logicalHeihgt int) {
	return width, height
}
