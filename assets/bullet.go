package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	bulletImagePath string  = "img/Bullet.png"
	bulletScale     float64 = 0.2
)

type Bullet struct {
	Img     *ebiten.Image
	XPos    int
	YPos    int
	Scale   float64
	Visible bool
}

func NewBullet(XPos, YPos int) *Bullet {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(bulletImagePath)
	if err != nil {
		log.Fatal(err)
	}
	return &Bullet{
		Img:     img,
		XPos:    XPos,
		YPos:    YPos,
		Scale:   bulletScale,
		Visible: false,
	}
}

func (bullet *Bullet) SetInitialPosition(XPos, YPos int) {
	bullet.YPos = YPos
	bullet.XPos = XPos
}

func (bullet *Bullet) MoveBullet(increment int) {
	// first trivial idea: increment ypos
	bullet.YPos = bullet.YPos - increment
}

func (bullet *Bullet) DrawUpdateBullet(drawingFunc func(*ebiten.Image, *ebiten.DrawImageOptions)) {
	bulletDrawOptions := &ebiten.DrawImageOptions{}
	bulletDrawOptions.GeoM.Scale(bullet.Scale, bullet.Scale)
	bulletDrawOptions.GeoM.Translate(float64(bullet.XPos), float64(bullet.YPos))
	drawingFunc(bullet.Img, bulletDrawOptions)
}
