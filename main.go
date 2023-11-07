package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Resource struct {
	Name   string
	Amount int
}

type Tile struct {
	ResourceType   string
	ResourceAmount int
	TerrainType    string
	Structure      string
}

type Button struct {
	X, Y   int32
	Width  int32
	Height int32
	Label  string
	Action func() // Function to be executed when the button is clicked
}

func NewButton(x, y, width, height int32, label string) Button {
	return Button{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Label:  label,
	}
}

func isMouseHover(x, y, width, height int32) bool {
	return rl.GetMouseX() >= x && rl.GetMouseX() <= x+width && rl.GetMouseY() >= y && rl.GetMouseY() <= y+height
}

func main() {
	const windowWidth = 800
	const windowHeight = windowWidth
	const screenSplitVertical = windowWidth / 2
	const screenSplitHorizontal = windowWidth / 2
	const tileWidth = 64
	const tileHeight = 64
	const tilesPerRow = 5
	const maxVisibleResources = 3
	var tiles []Tile
	var resources []Resource
	var scrollPosition = 0
	var devMode = false

	rl.InitWindow(windowWidth, windowHeight, "Planetarium")
	rl.SetTargetFPS(60)

	resources = []Resource{
		{"Silica", 100},
		{"Metal", 50},
		{"Energy", 200},
	}

	maxScrollPosition := len(resources) - maxVisibleResources
	selectedTileIndex := -1 // Initialize with an invalid index

	for i := 0; i < 25; i++ {
		resourceType := "Silica"
		resourceAmount := i * 10
		terrainType := "Plains"

		tile := Tile{
			ResourceType:   resourceType,
			ResourceAmount: resourceAmount,
			TerrainType:    terrainType,
			Structure:      "None",
		}

		tiles = append(tiles, tile)
	}

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
	solSystem := NewStarSystem(celestials)

	buttons := []Button{
		{
			X:      424,
			Y:      278,
			Width:  120,
			Height: 32,
			Label:  "Construct",
			Action: func() {
				if selectedTileIndex != -1 {
					tiles[selectedTileIndex].Structure = "Miner"
				}
			},
		},
	}

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyDown) && scrollPosition < maxScrollPosition {
			scrollPosition++
		}
		if rl.IsKeyPressed(rl.KeyUp) && scrollPosition > 0 {
			scrollPosition--
		}

		// Check for mouse wheel input
		scrollPosition -= int(rl.GetMouseWheelMove())
		if scrollPosition < 0 {
			scrollPosition = 0
		} else if scrollPosition > maxScrollPosition {
			scrollPosition = maxScrollPosition
		}

		if rl.IsKeyPressed(rl.KeyD) {
			if devMode {
				devMode = false
			} else {
				devMode = true
			}
		}

		for _, button := range buttons {
			if isMouseHover(button.X, button.Y, button.Width, button.Height) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				button.Action()
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Tile selection and drawing
		for i, tile := range tiles {
			// Calculate tile position
			x := int32(i%tilesPerRow*tileWidth + 32)
			y := int32(i/tilesPerRow*tileHeight + 32)

			// Check if the mouse is over the tile
			if isMouseHover(x, y, tileWidth, tileHeight) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				// Mouse has clicked this tile, store the index
				selectedTileIndex = i
			}

			if selectedTileIndex == i {
				rl.DrawRectangle(x, y, tileWidth, tileHeight, rl.Red)
				rl.DrawRectangleLines(x, y, tileWidth, tileHeight, rl.White)
			} else {
				rl.DrawRectangleLines(x, y, tileWidth, tileHeight, rl.White)
			}

			// Determine the icon text based on the resource type
			iconText := ""
			switch tiles[i].ResourceType {
			case "Silica":
				iconText = "Si"
			default:
				iconText = "?" // Default icon for unknown resource
			}

			iconSize := int32(40)
			iconX := x + tileWidth/2 - int32(rl.MeasureText(iconText, iconSize)/2)
			iconY := y + tileHeight/2 - iconSize/2
			rl.DrawText(iconText, iconX, iconY, int32(iconSize), rl.White)

			if tile.Structure == "Miner" && tile.ResourceType == "Silica" {
				// Check if the tile has resources
				if tile.ResourceAmount > 0 {
					// Decrease the resource count on the tile
					tiles[i].ResourceAmount-- // Update the tile within the slice

					// Find the resource with the matching type in the resources slice
					for j, resource := range resources {
						if resource.Name == "Silica" {
							// Increase the resource count in the resources slice
							resources[j].Amount++
							// Check if the tile is now depleted
							if tile.ResourceAmount == 0 {
								// Stop the miner by changing its structure
								tiles[i].Structure = "None" // Update the tile within the slice
							}
							break // Exit the loop once the resource is found
						}
					}
				}
			}

		}

		// Display information for the selected tile
		if selectedTileIndex >= 0 {
			tile := tiles[selectedTileIndex]
			textX := int32(screenSplitVertical + 20)
			textY := int32(40)
			infoText := fmt.Sprintf("Tile %d\nResourceType: %s\nResourceAmount: %d\nTerrainType: %s\nStructure: %s", selectedTileIndex+1, tile.ResourceType, tile.ResourceAmount, tile.TerrainType, tile.Structure)
			rl.DrawText(infoText, textX, textY, 20, rl.White)
		}

		// Draw the Sol System
		solSystem.Draw(screenSplitHorizontal, screenSplitVertical)

		// Draw the resources table
		for i := 0; i < maxVisibleResources; i++ {
			index := i + scrollPosition
			if index < len(resources) {
				resource := resources[index]
				textX := int32(387 + 8)
				textY := int32(387 + 8 + i*20)
				resourceText := fmt.Sprintf("%s: %d", resource.Name, resource.Amount)
				rl.DrawText(resourceText, textX, textY, 20, rl.White)
			}
		}
		rl.DrawRectangleLines(387, 387, 140, 20*maxVisibleResources+16, rl.White)

		for _, button := range buttons {
			rl.DrawRectangleLines(button.X, button.Y, button.Width, button.Height, rl.White)
			rl.DrawText(button.Label, button.X+8, button.Y+8, 20, rl.White)
		}

		if devMode {
			rl.DrawLine(rl.GetMouseX()-9999, rl.GetMouseY(), rl.GetMouseX()+9999, rl.GetMouseY(), rl.Pink)
			rl.DrawLine(rl.GetMouseX(), rl.GetMouseY()-9999, rl.GetMouseX(), rl.GetMouseY()+9999, rl.Pink)
			rl.DrawText(fmt.Sprintf("X: %d\nY: %d", rl.GetMouseX(), rl.GetMouseY()), rl.GetMouseX()+8, rl.GetMouseY()+8, 20, rl.Pink)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
