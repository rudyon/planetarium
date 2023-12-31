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
	const maxVisibleResources = 17
	var tiles []Tile
	var resources []Resource
	var structureRecipes []StructureRecipe
	var scrollPosition = 0
	var devMode = false
	var selectedTileIndex int = -1
	var maxScrollPosition int
	var container Container
	resources = initializeResources()
	structureRecipes = initializeStructureRecipes()
	solSystem := createSolSystem()
	selectedCelestialIndex := 3
	selectedCelestial := solSystem.Celestials[selectedCelestialIndex]
	selectedCelestialLabel := selectedCelestial.Name
	tiles = selectedCelestial.Tiles
	recpieButtons := []Button{}
	buttons := []Button{}

	rl.InitWindow(windowWidth, windowHeight, "Planetarium")
	rl.SetTargetFPS(60)

	// TODO tick based time passing
	// TODO lots of code is being reused (i am sorry)

	tabs := []ContainerTab{
		{
			"Resources",
			func() {
				// Draw the resources table here
				for i := 0; i < maxVisibleResources; i++ {
					index := i + scrollPosition
					if index < len(resources) {
						resource := resources[index]
						textX := int32(16)
						textY := int32(48 + i*20)
						resourceText := fmt.Sprintf("%s: %d", resource.Name, resource.Amount)
						rl.DrawText(resourceText, textX, textY, 20, rl.White)
					}
				}
				rl.DrawRectangleLines(8, 40, 220, 20*maxVisibleResources+16, rl.White)
			},
		},
		{
			"Recpies",
			func() {
				// Draw the buttons
				for i, button := range recpieButtons {
					rl.DrawRectangleLines(button.X, button.Y, button.Width, button.Height, rl.White)
					rl.DrawText(button.Label, button.X+8, button.Y+8, 20, rl.White)

					// Detect button clicks using the button's position
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(button.X+container.X), float32(button.Y+container.Y), float32(button.Width), float32(button.Height))) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							recpieButtons[i].Action()
						}
					}
				}

				// Draw tooltips for the hovered buttons
				for _, button := range recpieButtons {
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(button.X+container.X), float32(button.Y+container.Y), float32(button.Width), float32(button.Height))) {
						DrawButtonTooltip(button, container)
					}
				}
			},
		},
	}

	container = CreateContainer(380, 380, 400, 400, tabs)

	buttons = append(buttons, CreateButton(32, screenSplitHorizontal-16, 200, 32, selectedCelestialLabel, nil, ""))
	buttons = append(buttons, CreateButton(200+48, screenSplitHorizontal-16, 32, 32, "<", func() {
		if selectedCelestialIndex > 0 {
			selectedCelestialIndex--
			selectedCelestial = solSystem.Celestials[selectedCelestialIndex]
			tiles = selectedCelestial.Tiles
			selectedCelestialLabel = selectedCelestial.Name
			buttons[0] = CreateButton(32, screenSplitHorizontal-16, 200, 32, selectedCelestialLabel, nil, "")
		}
	}, ""))
	buttons = append(buttons, CreateButton(200+48*2, screenSplitHorizontal-16, 32, 32, ">", func() {
		if selectedCelestialIndex < len(solSystem.Celestials)-1 {
			selectedCelestialIndex++
			selectedCelestial = solSystem.Celestials[selectedCelestialIndex]
			tiles = selectedCelestial.Tiles
			selectedCelestialLabel = selectedCelestial.Name
			buttons[0] = CreateButton(32, screenSplitHorizontal-16, 200, 32, selectedCelestialLabel, nil, "")
		}
	}, ""))

	for i, recipe := range structureRecipes {
		structureName := recipe.Structure
		buttonLabel := "Construct " + structureName
		tooltip := "Required Resources: "

		// Generate a string with required resources for the tooltip
		for _, req := range recipe.RequiredResources {
			tooltip += fmt.Sprintf("\n%s: %d, ", req.ResourceName, req.Amount)
		}
		tooltip = strings.TrimSuffix(tooltip, ", ") // Remove the trailing comma and space

		recpieButtons = append(recpieButtons, CreateButton(8, int32(40+32*i), 220, 32, buttonLabel,
			func(structureName string, requiredResources []ResourceRequirement) func() {
				return func() {
					if selectedTileIndex != -1 {
						enoughResources := true

						// Check if required resources are available
						for _, req := range requiredResources {
							foundResource := false
							for _, resource := range resources {
								if resource.Name == req.ResourceName && resource.Amount >= req.Amount {
									foundResource = true
									break
								}
							}
							if !foundResource {
								enoughResources = false
								break
							}
						}

						if enoughResources {
							// Deduct the required resources
							for _, req := range requiredResources {
								for i, resource := range resources {
									if resource.Name == req.ResourceName {
										resources[i].Amount -= req.Amount
									}
								}
							}

							// Add the structure to the selected tile
							tiles[selectedTileIndex].Structure = structureName
						} else {
							// Provide feedback to the player that resources are insufficient
							fmt.Println("Insufficient resources to build", structureName)
						}
					}
				}
			}(structureName, recipe.RequiredResources), tooltip))
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		mousePosition := rl.GetMousePosition()

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

			iconText := getResourceIcon(tile.ResourceType)
			iconSize := int32(40)
			iconX := x + tileWidth/2 - int32(rl.MeasureText(iconText, iconSize)/2)
			iconY := y + tileHeight/2 - iconSize/2
			rl.DrawText(iconText, iconX, iconY, int32(iconSize), rl.White)

			// TODO this needs to be dynamic & tick system
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
		solSystem.Draw(screenSplitHorizontal, screenSplitVertical, selectedCelestial)

		for i, button := range buttons {
			rl.DrawRectangleLines(button.X, button.Y, button.Width, button.Height, rl.White)
			rl.DrawText(button.Label, button.X+8, button.Y+8, 20, rl.White)

			// Detect button clicks using the button's position
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(float32(button.X), float32(button.Y), float32(button.Width), float32(button.Height))) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					buttons[i].Action()
				}
			}
		}

		DrawTabContainer(&container)

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
