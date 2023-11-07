package main

import (
	"fmt"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	var structureRecipes []StructureRecipe
	var scrollPosition = 0
	var devMode = false
	var selectedTileIndex int = -1
	var maxScrollPosition int

	rl.InitWindow(windowWidth, windowHeight, "Planetarium")
	rl.SetTargetFPS(60)

	resources = initializeResources()
	structureRecipes = initializeStructureRecipes()
	tiles = createTiles()
	solSystem := createSolSystem()

	buttons := []Button{}

	for i, recipe := range structureRecipes {
		structureName := recipe.Structure
		buttonLabel := "Construct " + structureName
		tooltip := "Required Resources: "

		// Generate a string with required resources for the tooltip
		for _, req := range recipe.RequiredResources {
			tooltip += fmt.Sprintf("\n%s:%d, ", req.ResourceName, req.Amount)
		}
		tooltip = strings.TrimSuffix(tooltip, ", ") // Remove the trailing comma and space

		buttons = append(buttons, CreateButton(532, int32(390+32*i), 220, 32, buttonLabel,
			func(structureName string) func() {
				return func() {
					if selectedTileIndex != -1 {
						enoughResources := true

						for _, req := range recipe.RequiredResources {
							for _, resource := range resources {
								if resource.Name == req.ResourceName && resource.Amount < req.Amount {
									enoughResources = false
									break
								}
							}
							if !enoughResources {
								break
							}
						}

						if enoughResources {
							// Deduct the required resources
							for _, req := range recipe.RequiredResources {
								for i, resource := range resources {
									if resource.Name == req.ResourceName {
										resources[i].Amount -= req.Amount
									}
								}
							}

							// Add the structure to the selected tile
							tiles[selectedTileIndex].Structure = structureName
						}
					}
				}
			}(structureName), tooltip))
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeyDown) && scrollPosition < maxScrollPosition {
			scrollPosition++
		}
		if rl.IsKeyPressed(rl.KeyUp) && scrollPosition > 0 {
			scrollPosition--
		}

		// Detect button clicks
		mousePosition := rl.GetMousePosition()
		for i, button := range buttons {
			if rl.CheckCollisionPointRec(mousePosition, rl.NewRectangle(float32(button.X), float32(button.Y), float32(button.Width), float32(button.Height))) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					buttons[i].Action()
				} else if button.Tooltip != "" {
					DrawButtonTooltip(button)
				}
			}
		}

		// Check for mouse wheel input
		scrollPosition -= int(rl.GetMouseWheelMove())
		if scrollPosition < 0 {
			scrollPosition = 0
		} else if scrollPosition > maxScrollPosition {
			scrollPosition = maxScrollPosition
		}

		if rl.IsKeyPressed(rl.KeyD) {
			devMode = !devMode
		}

		for i, tile := range tiles {
			// Calculate tile position
			x := int32(i%tilesPerRow*tileWidth + 32)
			y := int32(i/tilesPerRow*tileHeight + 32)

			tileRec := rl.NewRectangle(float32(x), float32(y), float32(tileWidth), float32(tileHeight))

			if rl.CheckCollisionPointRec(mousePosition, tileRec) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
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
				iconText = "?" // Default icon for an unknown resource
			}

			iconSize := int32(40)
			iconX := x + tileWidth/2 - int32(rl.MeasureText(iconText, iconSize)/2)
			iconY := y + tileHeight/2 - iconSize/2
			rl.DrawText(iconText, iconX, iconY, int32(iconSize), rl.White)

			if tile.Structure == "Miner" && tile.ResourceType == "Silica" {
				// Check if the tile has resources
				if tile.ResourceAmount > 0 {
					// Decrease the resource count on the tile
					tiles[i].ResourceAmount--

					// Find the resource with the matching type in the resources slice
					for j, resource := range resources {
						if resource.Name == "Silica" {
							// Increase the resource count in the resources slice
							resources[j].Amount++

							// Check if the tile is now depleted
							if tile.ResourceAmount == 0 {
								// Stop the miner by changing its structure
								tiles[i].Structure = "None"
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
			const lineLength = 1000 // Adjust the line length as needed
			rl.DrawLine(rl.GetMouseX()-lineLength, rl.GetMouseY(), rl.GetMouseX()+lineLength, rl.GetMouseY(), rl.Pink)
			rl.DrawLine(rl.GetMouseX(), rl.GetMouseY()-lineLength, rl.GetMouseX(), rl.GetMouseY()+lineLength, rl.Pink)
			rl.DrawText(fmt.Sprintf("X: %d\nY: %d", rl.GetMouseX(), rl.GetMouseY()), rl.GetMouseX()+8, rl.GetMouseY()+8, 20, rl.Pink)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
