package main

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ButtonAction func()

type Button struct {
	X, Y    int32
	Width   int32
	Height  int32
	Label   string
	Action  ButtonAction
	Tooltip string
}

func CreateButton(x, y, width, height int32, label string, action ButtonAction, tooltip string) Button {
	return Button{
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		Label:   label,
		Action:  action,
		Tooltip: tooltip,
	}
}

func DrawButtonTooltip(button Button) {
	// Calculate the multiline text dimensions
	text := strings.Split(button.Tooltip, "\n")
	textWidth := int32(0)
	textHeight := int32(20 * len(text)) // Assuming a fixed text height

	for _, line := range text {
		lineWidth := rl.MeasureText(line, 20)
		if lineWidth > textWidth {
			textWidth = lineWidth
		}
	}

	// Calculate the tooltip position
	tooltipX := rl.GetMouseX() + 8
	tooltipY := rl.GetMouseY() + 8

	// Adjust tooltip position if it goes beyond the right or bottom screen edges
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	if tooltipX+textWidth > screenWidth {
		tooltipX = screenWidth - textWidth
	}
	if tooltipY+textHeight > screenHeight {
		tooltipY = screenHeight - textHeight
	}

	// Draw the black background rectangle to fit the multiline text
	rl.DrawRectangle(tooltipX-8, tooltipY-8, textWidth+8, textHeight+16, rl.Black)
	rl.DrawRectangleLines(tooltipX-8, tooltipY-8, textWidth+8, textHeight+16, rl.White)

	// Draw the multiline tooltip text on top of the background
	currentY := tooltipY
	for _, line := range text {
		rl.DrawText(line, tooltipX, currentY, 20, rl.White)
		currentY += 20
	}
}
