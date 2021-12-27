package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/abdfnx/resto/core/editor"
	"github.com/abdfnx/resto/core/editor/runtime"
	"github.com/rivo/tview"
)

func saveBuffer(b *editor.Buffer, path string) error {
	return ioutil.WriteFile(path, []byte(b.String()), 0600)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: secman settings [filename]\n")
		os.Exit(1)
	}

	path := os.Args[1]

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not read %v: %v", path, err)
	}

	var colorscheme editor.Colorscheme
	if railscast := runtime.Files.FindFile(editor.RTColorscheme, "railscast"); railscast != nil {
		if data, err := railscast.Data(); err == nil {
			colorscheme = editor.ParseColorscheme(string(data))
		}
	}

	app := tview.NewApplication()

	buffer := editor.NewBufferFromString(string(content), path)
	root := editor.NewView(buffer)
	root.SetRuntimeFiles(runtime.Files)
	root.SetColorscheme(colorscheme)

	root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyCtrlS:
				saveBuffer(buffer, path)
				app.Stop()
				return nil
			case tcell.KeyCtrlQ:
				app.Stop()
				return nil
			}
		return event
	})

	app.SetRoot(root, true)

	if err := app.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
