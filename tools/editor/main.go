package main

import (
	"log"
	"path"
	"io/ioutil"

	editor "github.com/scmn-dev/editor/core"
	"github.com/scmn-dev/editor/core/runtime"

	"github.com/rivo/tview"
	"github.com/tidwall/gjson"
	"github.com/gdamore/tcell/v2"
	"github.com/mitchellh/go-homedir"
)

var homeDir, _ = homedir.Dir()
var settingsFile = path.Join(homeDir, ".secman", "settings.json")

func saveBuffer(b *editor.Buffer, path string) error {
	return ioutil.WriteFile(path, []byte(b.String()), 0600)
}

func main() {
	content, err := ioutil.ReadFile(settingsFile)

	if err != nil {
		log.Fatalf("could not read %v: %v", settingsFile, err)
	}

	var colorscheme editor.Colorscheme

	vs := gjson.Get(string(content), "rs_settings.request_body.theme")
	tm := ""

	if vs.Exists() {
		tm = vs.String()
	} else {
		tm = "railscast"
	}

	if theme := runtime.Files.FindFile(editor.RTColorscheme, tm); theme != nil {
		if data, err := theme.Data(); err == nil {
			colorscheme = editor.ParseColorscheme(string(data))
		}
	}

	app := tview.NewApplication()

	buffer := editor.NewBufferFromString(string(content), settingsFile)
	root := editor.NewView(buffer)
	root.SetRuntimeFiles(runtime.Files)
	root.SetColorscheme(colorscheme)

	root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyCtrlS:
				saveBuffer(buffer, settingsFile)
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
