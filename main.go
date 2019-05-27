package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

func main() {
	output := tui.NewVBox()
	outputScroll := tui.NewScrollArea(output)
	outputScroll.SetAutoscrollToBottom(true)
	outputBox := tui.NewVBox(outputScroll)
	outputBox.SetBorder(true)
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)
	input.OnChanged(func(e *tui.Entry) {
		output.Append(tui.NewLabel(e.Text()))
	})
	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)
	root := tui.NewVBox(outputBox, inputBox)
	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}
	ui.SetKeybinding("Esc", func() {
		ui.Quit()
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
