//go:generate go run assets_generate.go

package runtime

import editor "github.com/scmn-dev/editor/core"

var Files = editor.NewRuntimeFiles(files)
