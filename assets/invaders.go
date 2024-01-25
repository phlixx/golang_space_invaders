package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	invaderImageBasePath string  = "img/Invader"
	invaderA1ImagePath   string  = "A1.png"
	invaderA2ImagePath   string  = "A2.png"
	invaderB1ImagePath   string  = "B1.png"
	invaderB2ImagePath   string  = "B2.png"
	invaderC1ImagePath   string  = "C1.png"
	invaderC2ImagePath   string  = "C2.png"
	invaderScale         float64 = 0.34
)

type Invader struct {
	Img   *ebiten.Image
	XPos  int
	YPos  int
	Scale float64
}

func getInvaderImage(invaderImagePath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(invaderImageBasePath + invaderImagePath)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func NewInvaderA1(XPos, YPos int) *Invader {
	return &Invader{
		Img:   getInvaderImage(invaderA1ImagePath),
		XPos:  XPos,
		YPos:  YPos,
		Scale: invaderScale * 0.7,
	}
}
func NewInvaderA2(XPos, YPos int) *Invader {
	return &Invader{
		Img:   getInvaderImage(invaderA2ImagePath),
		XPos:  XPos,
		YPos:  YPos,
		Scale: invaderScale * 0.7,
	}
}

func NewInvaderB1(XPos, YPos int) *Invader {
	return &Invader{
		Img:   getInvaderImage(invaderB1ImagePath),
		XPos:  XPos,
		YPos:  YPos,
		Scale: invaderScale * 0.9,
	}
}

func NewInvaderB2(XPos, YPos int) *Invader {
	return &Invader{
		Img:   getInvaderImage(invaderB2ImagePath),
		XPos:  XPos,
		YPos:  YPos,
		Scale: invaderScale * 0.9,
	}
}

func NewInvaderC1(XPos, YPos int) *Invader {
	return &Invader{
		Img:   getInvaderImage(invaderC1ImagePath),
		XPos:  XPos,
		YPos:  YPos,
		Scale: invaderScale,
	}
}

func NewInvaderC2(XPos, YPos int) *Invader {
	return &Invader{
		Img:   getInvaderImage(invaderC2ImagePath),
		XPos:  XPos,
		YPos:  YPos,
		Scale: invaderScale,
	}
}

func (inv *Invader) DrawUpdateInvader(drawingFunc func(*ebiten.Image, *ebiten.DrawImageOptions)) {
	invDrawOptions := &ebiten.DrawImageOptions{}
	invDrawOptions.GeoM.Scale(inv.Scale, inv.Scale)
	invDrawOptions.GeoM.Translate(float64(inv.XPos), float64(inv.YPos))
	drawingFunc(inv.Img, invDrawOptions)
}
