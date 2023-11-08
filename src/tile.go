package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	ResourceType   string
	ResourceAmount int
	TerrainType    string
	Structure      string
}

func createTiles(celestialType string) []Tile {
	var tiles []Tile

	var terrainProbabilities map[string]int
	var resourceProbabilities map[string]int

	switch celestialType {
	case "Terrestrial":
		terrainProbabilities = map[string]int{
			"Plains": 1,
		}
		resourceProbabilities = map[string]int{
			"Silica": 2,
			"Iron":   1,
		}
	case "Gas Giant":
		terrainProbabilities = map[string]int{
			"Gas": 1,
		}
		resourceProbabilities = map[string]int{
			"Helium": 1,
		}
	case "Star":
		terrainProbabilities = map[string]int{
			"Plasma": 1,
		}
		resourceProbabilities = map[string]int{
			"Helium": 1,
		}
	}

	for i := 0; i < 25; i++ {
		terrainType := getRandomType(terrainProbabilities)
		resourceType := getRandomType(resourceProbabilities)
		resourceAmount := i * 10

		tile := Tile{
			ResourceType:   resourceType,
			ResourceAmount: resourceAmount,
			TerrainType:    terrainType,
			Structure:      "None",
		}

		tiles = append(tiles, tile)
	}

	return tiles
}

// getRandomType selects a random type based on probabilities.
func getRandomType(probabilities map[string]int) string {
	totalProbability := 0
	for _, prob := range probabilities {
		totalProbability += prob
	}

	randNum := rl.GetRandomValue(0, int32(totalProbability))

	for t, prob := range probabilities {
		if randNum < int32(prob) {
			return t
		}
		randNum -= int32(prob)
	}

	// This should not happen, but if it does, return the first type
	for t := range probabilities {
		return t
	}

	return "Unknown"
}

func getResourceIcon(resourceType string) string {
	switch resourceType {
	case "Silica":
		return "Si"
	case "Iron":
		return "Fe"
	case "Helium":
		return "H"
	default:
		return "?"
	}
}
