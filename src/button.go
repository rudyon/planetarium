package main

import rl "github.com/gen2brain/raylib-go/raylib"

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
	rl.DrawText(button.Tooltip, 400, 600, 20, rl.White)
}
