package main
import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const ImagePath string = "img/Ship.png"

type Ship struct{
	Img *ebiten.Image
	XPos int16
	YPos int16
}
func CreateShip(XPos, YPos int16) *Ship{
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(ImagePath)
	if err != nil { 
		log.Fatal(err)
	}
	return &Ship{Img: img, XPos: XPos, YPos: YPos}
}