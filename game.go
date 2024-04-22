package main

import (
	"bytes"
	"log"

	"github.com/KidPudel/noodle_go/entities"
	raudio "github.com/KidPudel/noodle_go/resources"
	"github.com/KidPudel/noodle_go/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Game struct {
	Noodle             *entities.Noodle
	Flour              *entities.Flour
	VerticalGridLine   *ebiten.Image
	HorizontalGridLine *ebiten.Image

	Score     int
	GameTheme *audio.Player
}

func InitGame() *Game {
	noodle := entities.CreateNoodle()
	flour := entities.SpawnFlour()
	vertical := ebiten.NewImage(util.NoodleHeadSize/15, util.ScreenHeight)
	horizontal := ebiten.NewImage(util.ScreenWidth, util.NoodleHeadSize/15)
	audioContext := audio.NewContext(44100)
	stream, err := wav.DecodeWithoutResampling(bytes.NewReader(raudio.GameTheme))
	if err != nil {
		log.Fatal(err)
	}
	player, err := audioContext.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}
	player.Play()
	game := &Game{Noodle: noodle, Flour: flour, VerticalGridLine: vertical, HorizontalGridLine: horizontal, GameTheme: player}
	return game
}

// game's loop basic functionality

func (g *Game) Update() error {
	if !g.GameTheme.IsPlaying() {
		g.GameTheme.Rewind()
		g.GameTheme.Play()
	}
	g.Noodle.Update(&g.Score)
	g.Flour.Update(g.Noodle.Pos, &g.Score)

	return nil
}

func (g *Game) DrawGrid(screen *ebiten.Image) {
	g.VerticalGridLine.Fill(util.GridColor)
	g.HorizontalGridLine.Fill(util.GridColor)
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
	screen.Fill(util.BackgroundColor)
	g.DrawGrid(screen)
	g.Flour.Draw(screen)
	g.Noodle.Draw(screen)
}

func (g *Game) Layout(width, height int) (logicalWidth, logicalHeihgt int) {
	return width, height
}
