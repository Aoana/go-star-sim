package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Aoana/go-star-sim/internal/pkg/stars"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

// Game is part of ebiten and defines the game
type Game struct{}

var step int
var calcTime float64

// Layout is part of ebiten Game interface
// Defines the screen and is set to always run in full screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return stars.Data.Width, stars.Data.Height
}

// Update is part of ebiten Game interface
// Is called for every frame and executes one timestep
func (g *Game) Update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw all stars
	for i := range stars.Data.Stars {
		screen.Set(int(stars.Data.Stars[i].Px[step]), int(stars.Data.Stars[i].Py[step]), stars.White)
	}
	step++

	if step >= stars.Data.Steps {
		return errors.New("End of data")
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%s\n\nCurrent Step: %d\nTPS: %0.2f", stars.Data.Summary, step, ebiten.CurrentTPS()))
	return nil
}

func main() {

	game := &Game{}

	// Set radius of star cluster
	inputFile := flag.String("file", "/tmp/output", "File to read")
	flag.Parse()

	err := stars.Read(*inputFile, &stars.Data)
	if err != nil {
		log.Fatal("Unable to read file")
	}

	log.Println("Viewing simulation")
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Star System")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	log.Println("Simulation stopped", err)
}
