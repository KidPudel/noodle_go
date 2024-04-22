package entities

import (
	"slices"
	"time"

	"github.com/KidPudel/noodle_go/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Noodle struct {
	Pos       util.Vector2D
	Direction util.Vector2D
	Tail      []util.Vector2D

	noodleImage *ebiten.Image

	timeOfLastTick time.Time
	timeOfLastJump time.Time
}

func CreateNoodle() *Noodle {
	noodleImage := ebiten.NewImage(util.NoodleHeadSize, util.NoodleHeadSize)
	noodle := &Noodle{Pos: util.Vector2D{X: 20, Y: 20}, noodleImage: noodleImage, timeOfLastTick: time.Now(), timeOfLastJump: time.Now()}
	return noodle
}

func (n *Noodle) Update(score *int) {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		n.Direction = util.NewVector(util.NoodleHeadSize, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		n.Direction = util.NewVector(-util.NoodleHeadSize, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		n.Direction = util.NewVector(0, -util.NoodleHeadSize)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		n.Direction = util.NewVector(0, util.NoodleHeadSize)
	}

	if time.Since(n.timeOfLastJump).Seconds() >= time.Since(n.timeOfLastTick).Seconds()*5 {
		// append new tail, only when needed
		if *score > len(n.Tail) {
			n.Tail = append(n.Tail, util.NewVector(0, 0))
		}
		// translate head positions to the previous ones
		for i := range *score {
			reverseI := *score - i - 1
			if reverseI == 0 {
				n.Tail[reverseI] = n.Pos
			} else {
				n.Tail[reverseI] = n.Tail[reverseI-1]
			}
		}

		// push in direction
		n.Pos = n.Pos.Add(n.Direction)
		n.timeOfLastJump = time.Now()

		if n.Pos.X >= util.ScreenWidth {
			n.Pos.X = 0
		} else if n.Pos.X < 0 {
			n.Pos.X = util.ScreenWidth
		}
		if n.Pos.Y >= util.ScreenHeight {
			n.Pos.Y = 0
		} else if n.Pos.Y < 0 {
			n.Pos.Y = util.ScreenHeight
		}

		if slices.Contains(n.Tail, n.Pos) {
			*score = 0
			n.Tail = nil
		}

	}

	n.timeOfLastTick = time.Now()
}

func (n *Noodle) Draw(screen *ebiten.Image) {
	n.noodleImage.Fill(util.NoodleColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.Pos.X, n.Pos.Y)
	screen.DrawImage(n.noodleImage, op)

	for _, tail := range n.Tail {
		tailOp := &ebiten.DrawImageOptions{}
		tailOp.GeoM.Translate(tail.X, tail.Y)
		screen.DrawImage(n.noodleImage, tailOp)
	}

}
