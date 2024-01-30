package main

import (

	//"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/phlixx/golang_space_invaders/assets"
)

const (
	screenWidth      int = 224
	screenHeight     int = 256
	shipSingleMove   int = 1
	bulletSingleMove int = 3
	bulletOffset     int = 5
)

type Game struct {
	ship     *assets.Ship
	bullet   *assets.Bullet
	invaders *([]*assets.Invader)
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
	g.checkBulletCollision()

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

	// draw the invaders
	for _, invader := range *g.invaders {
		switch state := invader.State; state {
		case assets.StateDead:
			break
		case assets.StateAlive:
			invader.DrawUpdateInvader(screenUpdateFunc)
		case assets.StateExploding:
			go invader.KillAnimation(screenUpdateFunc)
		}
	}

	//ebitenutil.DebugPrint(screen, "Space Invaders")
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

func (g *Game) checkBulletCollision() {
	for _, invader := range *g.invaders {
		// ignore dead or already exploding assets
		if invader.State != assets.StateAlive {
			continue
		}
		if invader.InsideCollisionBox(g.bullet.XPos, g.bullet.YPos) {
			invader.State = assets.StateExploding
			g.bullet.XPos = g.ship.XPos
			g.bullet.YPos = g.ship.YPos
			g.bullet.Visible = false
			break
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

func (g *Game) createInvaders() {
	// spawn all invaders 6 x 11
	invaderFactories := []func(int, int) *assets.Invader{
		assets.NewInvaderA1,
		assets.NewInvaderA2,
		assets.NewInvaderB1,
		assets.NewInvaderB2,
		assets.NewInvaderC1,
		assets.NewInvaderC2,
	}

	// for each invader spawn 11
	invaders := make([]*assets.Invader, 6*11)
	const x_offset = 18
	xpos := (screenWidth - x_offset*11) / 2
	ypos := 40
	for row, invaderFactory := range invaderFactories {
		for col := 1; col <= 11; col++ {
			invaders[row*11+col-1] = invaderFactory(xpos, ypos)
			xpos = xpos + x_offset
		}
		xpos = xpos - 11*x_offset
		ypos = ypos + 20

	}

	// set invader game slice:

	g.invaders = &invaders
}
