package main
import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ShipImagePath string = "img/Ship.png"

type Ship struct{
	Img *ebiten.Image
	XPos int16
	YPos int16
}
func CreateShip(XPos, YPos int16) *Ship{
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(ShipImagePath)
	if err != nil { 
		log.Fatal(err)
	}
	return &Ship{Img: img, XPos: XPos, YPos: YPos}
}

type ShootingParticle struct{
	XPos int16
	YPos int16
}

func SpawnParticle(ShipXPos, ShipYPos int16){
	// bla
}