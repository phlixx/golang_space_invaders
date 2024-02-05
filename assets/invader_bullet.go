// TODO: animate the invader bullets like in the real game
package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	invaderBulletImagePath string  = "img/Bullet.png"
	invaderBulletScale     float64 = 0.2
)

type InvaderBullet struct {
	Img     *ebiten.Image
	XPos    int
	YPos    int
	Scale   float64
	Visible bool
}

func NewInvaderBullet(XPos, YPos int) *InvaderBullet {
	var err error
	var img *ebiten.Image
	img, _, err = ebitenutil.NewImageFromFile(bulletImagePath)
	if err != nil {
		log.Fatal(err)
	}
	return &InvaderBullet{
		Img:     img,
		XPos:    XPos,
		YPos:    YPos,
		Scale:   bulletScale,
		Visible: false,
	}
}
func (bullet *InvaderBullet) DrawUpdateBullet(drawingFunc func(*ebiten.Image, *ebiten.DrawImageOptions)) {
	bulletDrawOptions := &ebiten.DrawImageOptions{}
	bulletDrawOptions.GeoM.Scale(bullet.Scale, bullet.Scale)
	bulletDrawOptions.GeoM.Translate(float64(bullet.XPos), float64(bullet.YPos))
	drawingFunc(bullet.Img, bulletDrawOptions)
}

func (bullet *InvaderBullet) MoveInvaderBullet(increment int) {
	// first trivial idea: increment ypos in positive direction
	bullet.YPos = bullet.YPos + increment
}
