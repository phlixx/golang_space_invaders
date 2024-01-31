package assets

import (
	"log"
	"time"

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
	invaderDeathImage    string  = "img/RedInvader.png"
	invaderScale         float64 = 0.34
)

const (
	StateAlive     int = 0
	StateExploding int = 1
	StateDead      int = 2
)

const (
	MoveRight int = 0
	MoveDown  int = 1
	MoveLeft  int = 2
)

var invaderMovingPattern [13]int = [13]int{
	MoveRight,
	MoveRight,
	MoveRight,
	MoveRight,
	MoveRight,
	MoveRight,
	MoveDown,
	MoveLeft,
	MoveLeft,
	MoveLeft,
	MoveLeft,
	MoveLeft,
	MoveLeft,
}

type Invader struct {
	Img               *ebiten.Image
	XPos              int
	YPos              int
	Scale             float64
	State             int
	MoveCounter       int
	MoveCounterOffset int
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
		Img:               getInvaderImage(invaderA1ImagePath),
		XPos:              XPos,
		YPos:              YPos,
		Scale:             invaderScale * 0.7,
		State:             StateAlive,
		MoveCounter:       0,
		MoveCounterOffset: 10,
	}
}
func NewInvaderA2(XPos, YPos int) *Invader {
	return &Invader{
		Img:               getInvaderImage(invaderA2ImagePath),
		XPos:              XPos,
		YPos:              YPos,
		Scale:             invaderScale * 0.7,
		State:             StateAlive,
		MoveCounter:       0,
		MoveCounterOffset: 20,
	}
}

func NewInvaderB1(XPos, YPos int) *Invader {
	return &Invader{
		Img:               getInvaderImage(invaderB1ImagePath),
		XPos:              XPos,
		YPos:              YPos,
		Scale:             invaderScale * 0.9,
		State:             StateAlive,
		MoveCounter:       0,
		MoveCounterOffset: 30,
	}
}

func NewInvaderB2(XPos, YPos int) *Invader {
	return &Invader{
		Img:               getInvaderImage(invaderB2ImagePath),
		XPos:              XPos,
		YPos:              YPos,
		Scale:             invaderScale * 0.9,
		State:             StateAlive,
		MoveCounter:       0,
		MoveCounterOffset: 40,
	}
}

func NewInvaderC1(XPos, YPos int) *Invader {
	return &Invader{
		Img:               getInvaderImage(invaderC1ImagePath),
		XPos:              XPos,
		YPos:              YPos,
		Scale:             invaderScale,
		State:             StateAlive,
		MoveCounter:       0,
		MoveCounterOffset: 50,
	}
}

func NewInvaderC2(XPos, YPos int) *Invader {
	return &Invader{
		Img:               getInvaderImage(invaderC2ImagePath),
		XPos:              XPos,
		YPos:              YPos,
		Scale:             invaderScale,
		State:             StateAlive,
		MoveCounter:       0,
		MoveCounterOffset: 60,
	}
}

func (inv *Invader) DrawUpdateInvader(drawingFunc func(*ebiten.Image, *ebiten.DrawImageOptions)) {
	invDrawOptions := &ebiten.DrawImageOptions{}
	invDrawOptions.GeoM.Scale(inv.Scale, inv.Scale)
	invDrawOptions.GeoM.Translate(float64(inv.XPos), float64(inv.YPos))
	drawingFunc(inv.Img, invDrawOptions)
}

func (inv *Invader) KillAnimation(drawingFunc func(*ebiten.Image, *ebiten.DrawImageOptions)) {
	img, _, err := ebitenutil.NewImageFromFile(invaderDeathImage)
	if err != nil {
		log.Fatal(err)
	}

	inv.Img = img

	// sleep for 0.5 sec
	time.Sleep(500 * time.Millisecond)
	inv.State = StateDead

	// TODO: does not work
}

func (inv *Invader) InsideCollisionBox(xPos, yPos int) bool {
	// checks wheter a given x,y combination is in the collision box of an invader:

	width, height := float64(inv.Img.Bounds().Dx()), float64(inv.Img.Bounds().Dy())
	hitboxWidth := width * inv.Scale
	hitboxHeight := height * inv.Scale

	hitboxDiffX := float64(xPos) - float64(inv.XPos)
	hitboxDiffY := float64(yPos) - float64(inv.YPos)
	if hitboxDiffX > 0 && hitboxDiffY > 0 && (hitboxDiffX <= hitboxWidth) && (hitboxDiffY <= hitboxHeight) {
		return true
	}
	return false
}

func (inv *Invader) MoveInvader(increment int) {
	oldCount := inv.MoveCounter + inv.MoveCounterOffset
	// update move counter
	inv.MoveCounter = inv.MoveCounter + 1
	checkMove := oldCount%30 == 0
	if !checkMove {
		return
	}
	moveIndex := oldCount / 30 % 13
	currentMove := invaderMovingPattern[moveIndex]

	switch move := currentMove; move {
	case MoveLeft:
		inv.XPos = inv.XPos - increment
	case MoveDown:
		inv.YPos = inv.YPos + increment*3
	case MoveRight:
		inv.XPos = inv.XPos + increment
	}

}
