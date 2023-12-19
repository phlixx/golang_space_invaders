package main
import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ShipImagePath string = "img/Ship.png"
	BulletImagePath string = "img/Bullet.png"
)

type Ship struct{
	Img *ebiten.Image
	XPos int16
	YPos int16
}
func CreateShip(XPos, YPos int16) *Ship {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(ShipImagePath)
	if err != nil { 
		log.Fatal(err)
	}
	return &Ship{
		Img: img, 
		XPos: XPos, 
		YPos: YPos,
	}
}

type ShootingBullet struct {
	Img *ebiten.Image
	XPos int16
	YPos int16
	visible bool
}

func CreateBullet(XPos, YPos int16) *ShootingBullet {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(BulletImagePath)
	if err != nil { 
		log.Fatal(err)
	}
	return &ShootingBullet{
		Img: img, 
		XPos: XPos, 
		YPos: YPos,
		visible: false,
	}
}

// --> way to increment bullet movement
