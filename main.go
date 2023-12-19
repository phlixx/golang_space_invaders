package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Space Invaders Go")

	// init the game:
	var XPos int16 = screenWidth / 2
	var YPos int16 = screenHeight
	var ship *Ship = CreateShip(XPos, YPos)
	var bullet *Bullet = CreateBullet(XPos, YPos+bulletOffset)

	g := &Game{ship, bullet}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
