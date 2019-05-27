package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"time"

	c "github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

const (
	// Input box height.
	ih = 3
)

var buffer []string

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

// Set up the widgets and run the event loop.
func runGocui() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	// Create a new GUI.
	g, err := c.NewGui(c.OutputNormal)
	if err != nil {
		log.Println("Failed to create a GUI:", err)
		return
	}
	defer g.Close()

	// Activate the cursor for the current view.
	g.Cursor = true

	// The GUI object wants to know how to manage the layout.
	// Unlike `termui`, `gocui` does not use
	// a grid layout. Instead, it relies on a custom layout handler function
	// to manage the layout.

	// Here we set the layout manager to a function named `layout`
	// that is defined further down.
	g.SetManagerFunc(layout)

	// Bind the `quit` handler function (also defined further down) to Ctrl-C,
	// so that we can leave the application at any time.
	err = g.SetKeybinding("", c.KeyCtrlC, c.ModNone, quit)
	if err != nil {
		log.Println("Could not set key binding:", err)
		return
	}

	// Now let's define the views.

	// The terminal's width and height are needed for layout calculations.
	tw, th := g.Size()

	// Then the output view.
	ov, err := g.SetView("output", 0, 0, tw-1, th-ih-1)
	if err != nil && err != c.ErrUnknownView {
		log.Println("Failed to create output view:", err)
		return
	}
	ov.Title = "Output"
	ov.FgColor = c.ColorGreen
	// Let the view scroll if the output exceeds the visible area.
	ov.Autoscroll = true
	_, err = fmt.Fprintln(ov, "Press Ctrl-c to quit")
	if err != nil {
		log.Println("Failed to print into output view:", err)
	}

	// And finally the input view.
	iv, err := g.SetView("input", 0, th-ih, tw-1, th-1)
	if err != nil && err != c.ErrUnknownView {
		log.Println("Failed to create input view:", err)
		return
	}
	iv.Title = "Input"
	iv.FgColor = c.ColorYellow
	// The input view shall be editable.
	iv.Editable = true
	err = iv.SetCursor(0, 0)
	if err != nil {
		log.Println("Failed to set cursor:", err)
		return
	}

	// Make the enter key copy the input to the output.
	err = g.SetKeybinding("input", c.KeyEnter, c.ModNone, func(g *c.Gui, iv *c.View) error {
		// We want to read the view's buffer from the beginning.
		iv.Rewind()

		// Get the output view via its name.
		ov, e := g.View("output")
		if e != nil {
			log.Println("Cannot get output view:", e)
			return e
		}
		// Thanks to views being an io.Writer, we can simply Fprint to a view.
		_, e = fmt.Fprint(ov, iv.Buffer())
		if e != nil {
			log.Println("Cannot print to output view:", e)
		}
		// Clear the input view
		iv.Clear()
		// Put the cursor back to the start.
		e = iv.SetCursor(0, 0)
		if e != nil {
			log.Println("Failed to set cursor:", e)
		}
		return e

	})
	if err != nil {
		log.Println("Cannot bind the enter key:", err)
	}

	// Set the focus to the input view.
	_, err = g.SetCurrentView("input")
	if err != nil {
		log.Println("Cannot set focus to input view:", err)
	}

	// Start the main loop.
	err = g.MainLoop()
	log.Println("Main loop has finished:", err)
}

// The layout handler calculates all sizes depending
// on the current terminal size.
func layout(g *c.Gui) error {
	// Get the current terminal size.
	tw, th := g.Size()

	// Update the views according to the new terminal size.
	_, err := g.SetView("output", 0, 0, tw-1, th-ih-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update output view")
	}
	_, err = g.SetView("input", 0, th-ih, tw-1, th-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update input view.")
	}
	return nil
}

// `quit` is a handler that gets bound to Ctrl-C.
// It signals the main loop to exit.
func quit(g *c.Gui, v *c.View) error {
	return c.ErrQuit
}

/*
Our main func just needs to read the name from the TUI lib from the command line
and execute the respective code.
*/

//
func main() {
	runGocui()
}
