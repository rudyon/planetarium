package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StarSystem struct {
	Celestials []Celestial
}

func NewStarSystem(celestials []Celestial) StarSystem {
	// Set random initial angles and rotation speeds for celestial objects
	for i := range celestials {
		celestials[i].Angle = rand.Float64() * 2 * math.Pi
		minSpeed := 6.0
		maxSpeed := 12.0
		celestials[i].RotationSpeed = (minSpeed + rand.Float64()*(maxSpeed-minSpeed)) * 6
	}

	return StarSystem{Celestials: celestials}
}

func (s *StarSystem) Draw(screenSplitHorizontal, screenSplitVertical float32, selectedCelelestial Celestial) {
	numCelestials := len(s.Celestials)
	totalDistance := float32(0)

	// Determine the bottom left corner coordinates using screenSplit variables
	bottomLeftX := screenSplitHorizontal / 2
	bottomLeftY := screenSplitVertical * 1.5

	for i := 0; i < numCelestials; i++ {
		celestial := &s.Celestials[i]

		celestial.Angle += 0.005 // Reduced rotation speed

		rl.DrawCircleLines(int32(bottomLeftX), int32(bottomLeftY), totalDistance, celestial.Color)

		// Calculate the distance dynamically to make the celestial objects touch
		distance := celestial.Radius
		if i < numCelestials-1 {
			// Add the radius of the next celestial to the distance
			distance += s.Celestials[i+1].Radius
		}

		// Calculate the coordinates for the bottom left quarter
		x := float32(bottomLeftX) + totalDistance*float32(math.Cos(celestial.Angle))
		y := bottomLeftY + totalDistance*float32(math.Sin(celestial.Angle))

		if celestial.Name == selectedCelelestial.Name {
			rl.DrawCircle(int32(x), int32(y), float32(celestial.Radius+2), rl.White)
		}

		rl.DrawCircle(int32(x), int32(y), float32(celestial.Radius), celestial.Color)

		// Accumulate the distance for the next celestial
		totalDistance += distance
	}
}

func createSolSystem() StarSystem {
	// Define the celestial objects for the Sol System
	sun := NewCelestial("Sol", "Star", 20, rl.Yellow)
	mercury := NewCelestial("Mercury", "Terrestrial", 4, rl.LightGray)
	venus := NewCelestial("Venus", "Terrestrial", 5.5, rl.Orange)
	earth := NewCelestial("Terra", "Terrestrial", 6, rl.Blue)
	mars := NewCelestial("Mars", "Terrestrial", 5, rl.Red)
	jupiter := NewCelestial("Jupiter", "Gas Giant", 18, rl.Brown)
	saturn := NewCelestial("Saturn", "Gas Giant", 16, rl.Beige)
	uranus := NewCelestial("Uranus", "Gas Giant", 10, rl.SkyBlue)
	neptune := NewCelestial("Neptune", "Gas Giant", 9, rl.DarkBlue)

	// Create the StarSystem with the Celestial objects
	celestials := []Celestial{sun, mercury, venus, earth, mars, jupiter, saturn, uranus, neptune}
	return NewStarSystem(celestials)
}
