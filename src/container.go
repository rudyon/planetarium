package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Content func()

type ContainerTab struct {
	Label      string
	TabContent Content
}

type Container struct {
	X, Y        int32
	Width       int32
	Height      int32
	SelectedTab int
	Tabs        []ContainerTab
}

func CreateContainer(x, y, width, height int32, tabs []ContainerTab) Container {
	return Container{
		X:           x,
		Y:           y,
		Width:       width,
		Height:      height,
		SelectedTab: 0,
		Tabs:        tabs,
	}
}

func DrawTabContainer(container *Container) {
	tabHeight := int32(32)
	tabY := container.Y

	// Calculate the tab width based on the number of tabs
	tabWidth := container.Width / int32(len(container.Tabs))

	// Draw tabs
	for i, tab := range container.Tabs {
		tabX := container.X + int32(i)*tabWidth
		tabRect := rl.NewRectangle(float32(tabX), float32(tabY), float32(tabWidth), float32(tabHeight))

		// Check if the current tab is selected
		if i == container.SelectedTab {
			rl.DrawRectangleRec(tabRect, rl.Black)
			rl.DrawRectangleLinesEx(tabRect, 1, rl.White)
			rl.DrawText(tab.Label, tabX+8, tabY+8, 20, rl.White)
		} else {
			rl.DrawRectangleRec(tabRect, rl.Black)
			rl.DrawRectangleLinesEx(tabRect, 1, rl.Gray)
			rl.DrawText(tab.Label, tabX+8, tabY+8, 20, rl.Gray)
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), tabRect) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			container.SelectedTab = i
		}
	}

	// Draw content of the selected tab
	selectedTab := container.Tabs[container.SelectedTab]

	// Adjust the position for the content based on the container's location
	rl.PushMatrix()                                              // Push the current transformation matrix onto the stack
	rl.Translatef(float32(container.X), float32(container.Y), 0) // Translate to the container's position

	selectedTab.TabContent()

	rl.PopMatrix() // Restore the previous transformation matrix
}
