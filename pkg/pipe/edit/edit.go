package edit

import (
	"fmt"
	"os"

	"github.com/abdfnx/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/internal/shared"
	"github.com/scmn-dev/secman/v6/pkg/options"
	"github.com/scmn-dev/secman/v6/pkg/pipe/edit/editor"
)

type model struct {
	styles   shared.Styles
	list     list.Model
	pwType   string
	password string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.list.SetWidth(msg.Width)
			return m, nil

		case tea.KeyMsg:
			switch keypress := msg.String(); keypress {
				case "ctrl+c":
					return m, tea.Quit

				case "enter":
					i, ok := m.list.SelectedItem().(shared.Item)

					if ok {
						if err := tea.NewProgram(editor.Editor(m.pwType, string(i), m.password)).Start(); err != nil {
							fmt.Printf("could not start editor: %s\n", err)
							os.Exit(1)
						}
					}

					return m, tea.Quit
		}
	}

	var cmd tea.Cmd

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Editor")).String() + "\n" + m.list.View()
}

func Edit(o *options.PasswordsOptions) {
	items := []list.Item{}
	st := shared.DefaultStyles()
	var p = fmt.Sprintf("\"%s\"", o.Password)

	if o.Logins {
		items = append(items, shared.Item("Title"), shared.Item("Username"), shared.Item("Password"), shared.Item("URL"), shared.Item("Extra"))
	} else if o.CreditCards {
		items = append(items, shared.Item("Card Name"), shared.Item("Card Holder Name"), shared.Item("Card Type"), shared.Item("Number"), shared.Item("Expiry Date"), shared.Item("Verification Number"))
	} else if o.Emails {
		items = append(items, shared.Item("Email Address"), shared.Item("Password"))
	} else if o.Notes {
		items = append(items, shared.Item("Title"), shared.Item("Note"))
	} else if o.Servers {
		items = append(items, shared.Item("Title"), shared.Item("URL"), shared.Item("Username"), shared.Item("Password"), shared.Item("Hosting Username"), shared.Item("Hosting Password"), shared.Item("Admin Username"), shared.Item("Admin Password"), shared.Item("Extra"))
	}

	l := list.New(items, shared.ItemDelegate{}, constants.LIST_WIDTH, constants.LIST_HEIGHT)
	l.Title = "\nChoose which password to edit, or press `q` to quit."
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = st.BaseStyle
	l.Styles.PaginationStyle = st.Pagination

	m := model{
		styles:   st,
		list:     l,
		pwType:   shared.PasswordType(o),
		password: p,
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
