package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/marcusolsson/tui-go"
)

var buffer []string

/* Function to run the groutine to run for stdin read */
func read(r io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			s := scan.Text()
			lines <- s
		}
	}()
	return lines
}

func main() {
	ch := read(os.Stdin) //Reading from Stdin
	go func() {
		for m := range ch {
			buffer = append(buffer, m)
		}
	}()

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
