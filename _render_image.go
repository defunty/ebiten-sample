package main

import (
	_ "image/jpeg"
	_ "image/gif"
	"log"
	"fmt" // for debug

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth = 320
	screenHeight = 240
)

var (
	img *ebiten.Image
)

// for debug
func OutputLog(s string) {
	fmt.Println(s)
	// log.Fatal(err)
}

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("i_robot.gif")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	count int
}

// error型を返す
func (g *Game) Update() error {
	//g.count++
	return nil
}

// error型を返す
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Render Image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}