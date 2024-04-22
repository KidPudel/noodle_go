package main

import (
	"log"

	"github.com/KidPudel/noodle_go/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(util.ScreenWidth, util.ScreenHeight)
	ebiten.SetWindowTitle("noodle")
	game := InitGame()
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
