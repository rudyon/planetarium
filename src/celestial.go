package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Celestial struct {
	Name          string
	Type          string
	Radius        float32
	Color         rl.Color
	Angle         float64
	RotationSpeed float64
}

func NewCelestial(name string, celest_type string, radius float32, color rl.Color) Celestial {
	return Celestial{Name: name, Type: celest_type, Radius: radius, Color: color}
}
