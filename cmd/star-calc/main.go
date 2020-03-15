package main

import (
	"encoding/json"
	"flag"
	"github.com/Aoana/go-star-sim/internal/pkg/stars"
	"io/ioutil"
	"log"
	"time"
)

var timestep func()

func main() {

	// Set radius of star cluster
	numbStars := flag.Int("numbStars", 12, "Number of stars in cluster")
	numbSteps := flag.Int("numbSteps", 100, "Number of time steps")
	calcModel := flag.String("calcModel", "Exact", "\"Exact\" or \"BarnesHut\"")
	gridWidth := flag.Int("gridWidth", 1920, "Grid width size")
	gridHeight := flag.Int("gridHeight", 1080, "Grid height size")
	outputFile := flag.String("outputFile", "/tmp/output", "Path to output file")
	flag.Parse()

	// Store simulation data
	stars.Data.Width = *gridWidth
	stars.Data.Height = *gridHeight
	stars.Data.Steps = *numbSteps
	stars.Data.Model = *calcModel

	switch *calcModel {
	case "Exact":
		timestep = stars.TimestepExact
	case "BarnesHut":
		timestep = stars.TimestepBarnesHut
	default:
		log.Fatal("Unknown gravity model")
	}

	// Spawn all stars
	stars.StartValues(*numbStars)

	// Create coordinate slice for storing positions
	pixels := make([]stars.Pixel, len(stars.StarList))
	for i := range pixels {
		pixels[i].Px = make([]uint16, stars.Data.Steps)
		pixels[i].Py = make([]uint16, stars.Data.Steps)
	}

	// Run simulation
	log.Println("Simulation starting")
	log.Printf("Stars=%d, Model=%s, Grid=%dx%d, Timesteps=%d", len(stars.StarList), *calcModel, stars.Data.Width, stars.Data.Height, stars.Data.Steps)

	start := time.Now()
	for steps := 0; steps < stars.Data.Steps; steps++ {
		// Physical calculation (based on method)
		timestep()
		// Store position for post processing
		for s := range stars.StarList {
			pixels[s].Px[steps] = uint16(stars.StarList[s].X)
			pixels[s].Px[steps] = uint16(stars.StarList[s].Y)
		}
	}
	stars.Data.Time = time.Since(start)
	log.Println("Simulation complete, storing to file")

	stars.Data.Stars = pixels

	log.Printf("%+v", stars.Data)

	f, _ := json.MarshalIndent(stars.Data, "", " ")
	_ = ioutil.WriteFile(*outputFile, f, 0644)

	// Debug check to read data
	var check stars.SimData
	tmpdata, err := ioutil.ReadFile(*outputFile)
	if err != nil {
		log.Fatal("Unable to create file")
	}
	err = json.Unmarshal(tmpdata, &check)
	if err != nil {
		log.Println("error:", err)
	}
	log.Printf("%+v", check)

}
