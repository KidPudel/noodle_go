package entities

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/KidPudel/noodle_go/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Flour struct {
	Pos        util.Vector2D
	FlourImage *ebiten.Image
}

func SpawnFlour() *Flour {
	img := ebiten.NewImage(util.NoodleHeadSize, util.NoodleHeadSize)
	flourPos := util.NewVector(math.Ceil(float64(rand.Intn(util.ScreenWidth))/util.NoodleHeadSize)*util.NoodleHeadSize, math.Ceil(float64(rand.Intn(util.ScreenHeight))/util.NoodleHeadSize)*util.NoodleHeadSize)
	flour := &Flour{Pos: flourPos, FlourImage: img}
	return flour
}

func (f *Flour) Update(noodlePos util.Vector2D, score *int) {
	if noodlePos == f.Pos {
		*score++
		f.Pos = util.NewVector(math.Ceil(float64(rand.Intn(util.ScreenWidth))/util.NoodleHeadSize)*util.NoodleHeadSize, math.Ceil(float64(rand.Intn(util.ScreenHeight))/util.NoodleHeadSize)*util.NoodleHeadSize)
	}
}

func (f *Flour) Draw(screen *ebiten.Image) {
	f.FlourImage.Fill(color.RGBA{255, 255, 255, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.Pos.X, f.Pos.Y)
	screen.DrawImage(f.FlourImage, op)
}
