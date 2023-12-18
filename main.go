package main

import (
	"log"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)
const (
	screenWidth  = 320
	screenHeight = 240
)
var ship *Ship
// init the image
func init() {
	var XPos int16 = screenWidth/2
	var YPos int16 = screenHeight

	ship = CreateShip(XPos, YPos)


}
type Game struct{}

func (g *Game) Update() error {
	// movement of ship 
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		// shoot
		return nil
	}
	if inpututil.KeyPressDuration(ebiten.KeyRight) >0 {
		// move left
		g.MoveShip(3)
	}
	if inpututil.KeyPressDuration(ebiten.KeyLeft)>0 {
		// move right
		g.MoveShip(-3)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x29, 0x62, 0xff, 0xff})
	// draw the ship:
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(ship.XPos), float64(ship.YPos))
	op.GeoM.Translate(-float64(ship.Img.Bounds().Dx())/4, -float64(ship.Img.Bounds().Dy())/2)
	screen.DrawImage(ship.Img, op)

	ebitenutil.DebugPrint(screen, "Space Invaders")
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) MoveShip(steps int16) error{
	newPos := ship.XPos + steps
	if 0 <= float64(newPos) - float64(ship.Img.Bounds().Dx())/4 && float64(newPos) +float64(ship.Img.Bounds().Dx())/4 <= screenWidth {
		ship.XPos = newPos
	}
	return nil
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
