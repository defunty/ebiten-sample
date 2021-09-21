package main

import (
	"bytes"
	"image"
	_"image/jpeg"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth = 320
	screenHeight = 240
)

var (
	gophersImage *ebiten.Image
)

type Game struct {
	count int
}

// error型を返す
func (g *Game) Update() error {
	g.count++
	return nil
}

// error型を返す
func (g *Game) Draw(screen *ebiten.Image) {
	w, h := gophersImage.Size()
	op := &ebiten.DrawImageOptions{}

	// GeoM https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#GeoM
	// A GeoM represents a matrix to transform geometry when rendering an image.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(g.count%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	screen.DrawImage(gophersImage, op) // https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rotate (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}