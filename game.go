package main

import (

	//"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/phlixx/golang_space_invaders/assets"
)

const (
	screenWidth      int = 320
	screenHeight     int = 240
	shipSingleMove   int = 3
	bulletSingleMove int = 3
	bulletOffset     int = 5
)

type Game struct {
	ship   *assets.Ship
	bullet *assets.Bullet
}

func (g *Game) Update() error {
	// movement of ship
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && !g.bullet.Visible {
		// start shooting
		g.activateBullet()
	}
	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 {
		// move right
		g.ship.MoveShip(shipSingleMove, screenWidth)
	}
	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 {
		// move left
		g.ship.MoveShip(-shipSingleMove, screenWidth)
	}

	g.moveAssets()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.RGBA{0x29, 0x62, 0xff, 0xff})
	screenUpdateFunc := screen.DrawImage
	// draw the bullet:
	if g.bullet.Visible {
		g.bullet.DrawUpdateBullet(screenUpdateFunc)
	}
	g.ship.DrawUpdateShip(screenUpdateFunc)

	ebitenutil.DebugPrint(screen, "Space Invaders")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) moveAssets() {
	// if bullet is visible: move it
	if g.bullet.Visible {
		g.bullet.MoveBullet(bulletSingleMove)
		// check if bullet is outside drawing area and then reset
		if g.bullet.YPos <= 0 {
			g.bullet.Visible = false
		}
	}
}

func (g *Game) activateBullet() {
	// set initial position to current position of ship - 1/2 of the bullet size
	shipXPos := g.ship.XPos
	shipYPos := g.ship.YPos - int(float64(g.ship.Img.Bounds().Dy())*g.ship.Scale) + bulletOffset
	bulletSizeOffset := int(float64(g.bullet.Img.Bounds().Dx()/2) * g.bullet.Scale)

	g.bullet.SetInitialPosition(shipXPos-bulletSizeOffset, shipYPos)
	g.bullet.Visible = true
}
