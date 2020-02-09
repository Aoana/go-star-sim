package stars

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/atedja/go-vector"
	"image/color"
)

// Simulation variables
var (
	dt = 1.0
	// Width
	W int
	// Height
	H int
	// The color white
	White = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
)

// StarList is a global slice of objects
var StarList []*objects.Object

func init() {
}

// StartValues set starting position and velocity
// Fixed starting position and velocity is random
func StartValues(r int) error {

	T := vector.NewWithValues([]float64{float64(W / 2), float64(H / 2)})

	for i := -r; i <= r; i++ {
		for j := -r; j <= r; j++ {
			if i*i+j*j <= r*r {

				// Logical starting position
				X := vector.NewWithValues([]float64{float64(i), float64(j)})

				// Velocity vector
				var V vector.Vector

				// Velocity perpendicular to circle
				if X[0] == 0 && X[1] == 0 {
					continue
				} else if X[0] == 0 {
					V = vector.NewWithValues([]float64{X[1], 0.0})
				} else if X[0] > 0 {
					V = vector.NewWithValues([]float64{X[1] / X[0], -X[0]})
				} else {
					V = vector.NewWithValues([]float64{-X[1] / X[0], -X[0]})
				}

				// Velocity vector with fixed length
				VS := vector.Unit(V)
				VS.Scale(0.5)

				// Translate position to middle of screen
				XT := vector.Add(X, T)

				// Construct objects
				s, err := objects.New(XT[0], XT[1], VS[0], VS[1])
				if err != nil {
					return err
				}
				StarList = append(StarList, s)
			}
		}
	}
	fmt.Println("Number of stars: ", len(StarList))
	return nil
}

// TimestepStars updates position and velocity of all stars
func TimestepStars() error {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		err := StarList[i].Position(dt)
		if err != nil {
			return err
		}
	}

	// Update velocities of all stars based on gravity calculation
	for i := range StarList {
		err := StarList[i].Velocity(0, 0, dt)
		if err != nil {
			return err
		}
	}
	return nil

}
