package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/phlixx/golang_space_invaders/assets"
)

func initGame() *Game {
	var XPos int = screenWidth / 2
	var YPos int = screenHeight
	var ship *assets.Ship = assets.NewShip(XPos, YPos)
	var bullet *assets.Bullet = assets.NewBullet(XPos, YPos+bulletOffset)

	g := &Game{ship, bullet, nil}
	g.createInvaders()
	return g
}

func main() {
	// set window options
	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Space Invaders Go")
	// init game
	g := initGame()
	// run game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
