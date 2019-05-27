package main

import (
	"bufio"
	"io"
	"os"

	"github.com/rivo/tview"
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
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
