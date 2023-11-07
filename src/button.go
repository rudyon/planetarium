package main

type ButtonAction func()

type Button struct {
	X, Y   int32
	Width  int32
	Height int32
	Label  string
	Action ButtonAction
}

func CreateButton(x, y, width, height int32, label string, action ButtonAction) Button {
	return Button{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Label:  label,
		Action: action,
	}
}
