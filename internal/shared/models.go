package shared

import (
	"io"
	"fmt"

	"github.com/abdfnx/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	Ready State = iota
	Loading
)

type SetMsg       string
type ErrorMsg     struct{}
type OtherMsg     struct{}
type SuccessMsg   struct{}
type Message      struct{ Err error }
type State        int
type Index        int
type Item         string
type ItemDelegate struct{}

func (e Message) 	  Error() string { return e.Err.Error() }
func (i Item) 		  FilterValue() string { return "" }
func (d ItemDelegate) Height() int                               { return 1 }
func (d ItemDelegate) Spacing() int                              { return 0 }
func (d ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	st := DefaultStyles()

	i, ok := listItem.(Item)

	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := st.Item.Render

	if index == m.Index() {
		fn = func(s string) string {
			return st.SelectedMenuItem.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}
