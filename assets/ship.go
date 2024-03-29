package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	shipImagePath string  = "img/Ship.png"
	shipScale     float64 = 0.2
)

type Ship struct {
	Img   *ebiten.Image
	XPos  int
	YPos  int
	Scale float64
}

func NewShip(XPos, YPos int) *Ship {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(shipImagePath)
	if err != nil {
		log.Fatal(err)
	}
	return &Ship{
		Img:   img,
		XPos:  XPos,
		YPos:  YPos,
		Scale: shipScale,
	}
}
func (ship *Ship) MoveShip(increment int, screenWidth int) {
	// first trivial idea: increment ypos
	newShipXPos := ship.XPos + increment
	if 0 <= float64(newShipXPos)-float64(ship.Img.Bounds().Dx())/2*float64(ship.Scale) && float64(newShipXPos)+float64(ship.Img.Bounds().Dx())/2*float64(ship.Scale) <= float64(screenWidth) {
		ship.XPos = newShipXPos
	}
}
func (ship *Ship) DrawUpdateShip(drawingFunc func(*ebiten.Image, *ebiten.DrawImageOptions)) {
	// draw the ship:
	shipDrawOptions := &ebiten.DrawImageOptions{}
	shipDrawOptions.GeoM.Scale(ship.Scale, ship.Scale)
	shipDrawOptions.GeoM.Translate(float64(ship.XPos), float64(ship.YPos))
	shipDrawOptions.GeoM.Translate(-float64(ship.Img.Bounds().Dx())/2*ship.Scale, -float64(ship.Img.Bounds().Dy())*ship.Scale)
	drawingFunc(ship.Img, shipDrawOptions)
}
