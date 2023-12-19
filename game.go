package main

import (

	//"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth    = 320
	screenHeight   = 240
	shipSingleMove = 3
	bulletOffset   = 5
)

type Game struct {
	ship   *Ship
	bullet *Bullet
}

func (g *Game) Update() error {
	// movement of ship
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// shoot
		g.Shoot()
	}
	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 {
		// move left
		g.MoveShip(shipSingleMove)
	}
	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 {
		// move right
		g.MoveShip(-shipSingleMove)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{0x29, 0x62, 0xff, 0xff})
	// draw the ship:
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(g.ship.XPos), float64(g.ship.YPos))
	op.GeoM.Translate(-float64(g.ship.Img.Bounds().Dx())/4, -float64(g.ship.Img.Bounds().Dy())/2)
	screen.DrawImage(g.ship.Img, op)

	// draw the bullet:
	if g.bullet.visible {
		// TODO: adjust the initial position
		screen.DrawImage(g.bullet.Img, op)
	}

	ebitenutil.DebugPrint(screen, "Space Invaders")
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) MoveShip(steps int16) {
	newPos := g.ship.XPos + steps
	if 0 <= float64(newPos)-float64(g.ship.Img.Bounds().Dx())/4 && float64(newPos)+float64(g.ship.Img.Bounds().Dx())/4 <= screenWidth {
		g.ship.XPos = newPos
	}
}

func (g *Game) Shoot() {
	// start shooting process
	g.bullet.visible = true
}
