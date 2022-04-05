package delete

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/gosh"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/internal/shared"
	"github.com/scmn-dev/secman/v6/pkg/options"
)

const (
	okButton shared.Index = iota
	cancelButton
)

type model struct {
	styles   shared.Styles
	state    shared.State
	index    shared.Index
	spinner  spinner.Model
	message  string
	out      string
	err   	 error
	password string
}

func (m *model) indexForward() {
	m.index = (m.index + 1) % 2
}

func (m *model) indexBackward() {
	m.index = (m.index - 1) % 2
}

func Delete(o *options.PasswordsOptions) model {
	st := shared.DefaultStyles()
	var p = fmt.Sprintf("\"%s\"", o.Password)

	var err, out, _ = gosh.RunOutput("scc delete " + shared.PasswordType(o) + " " + p)

	return model{
		styles:   st,
		state:    shared.Ready,
		index:    okButton,
		message:  "",
		out:      out,
		spinner:  shared.NewSpinner(),
		password: p,
		err:      err,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.out != "" {
		switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.Type {
					case tea.KeyCtrlC:
						return m, tea.Quit

					default:
						if m.state == shared.Loading {
							return m, nil
						}

						switch msg.String() {
							case "tab":
								m.indexForward()

							case "shift+tab":
								m.indexBackward()

							case "l", "k", "right":
								m.indexForward()

							case "h", "j", "left":
								m.indexBackward()

							case "enter":
								switch m.index {
									case okButton:
										m.state = shared.Loading
										m.message = ""

										return m, tea.Batch(
											smd(m),
											spinner.Tick,
										)

									case cancelButton:
										return m, tea.Quit
								}
						}

						return m, nil
				}

			case shared.SuccessMsg:
				m.state = shared.Ready
				head := m.styles.Success.Render(m.styles.Checkmark.String() + m.password)
				body := m.styles.Subtle.Render(" was deleted successfully.")
				m.message = m.styles.Wrap.Render(head + body)

				return m, nil

			case shared.OtherMsg:
				m.state = shared.Ready
				head := m.styles.Error.Render("Your Authentication is Expired.")
				body := m.styles.Subtle.Render(" Refresh your authentication via `secman auth refresh`.\n")
				m.message = m.styles.Wrap.Render(head + body)

				return m, nil

			case shared.Message:
				m.state = shared.Ready
				head := m.styles.Error.Render("Oh, what? There was a curious error we were not expecting.\n")
				body := m.styles.Subtle.Render(msg.Error())
				m.message = m.styles.Wrap.Render(head + body)

				return m, nil

			case spinner.TickMsg:
				var cmd tea.Cmd
				m.spinner, cmd = m.spinner.Update(msg)

				return m, cmd

			default:
				var cmd tea.Cmd

				return m, cmd
		}
	} else {
		return m, tea.Quit
	}
}

func (m model) View() string {
	s := ""

	if m.out != "" {
		s += "\n\nAre you sure you want to delete " + m.password + " ?\n\n"
	} else {
		s += "\n\n" + m.password + " is not found.\n"
	}

	if m.state == shared.Loading {
		s += spinnerView(m)
	} else {
		s += shared.OKButton(m.index == 0, true, "Yes")
		s += " " + shared.CancelButton(m.index == 1, false, "Exit")

		if m.message != "" {
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(m.message).String())
			os.Exit(0)
		}
	}

	return lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Deleter") + s).String()
}

func spinnerView(m model) string {
	return m.spinner.View() + "🪓 Deleting..."
}

func smd(m model) tea.Cmd {
	return func() tea.Msg {
		cmdOut := strings.TrimSuffix(m.out, "\n")

		if strings.Contains(cmdOut, "401") {	
			return shared.OtherMsg{}
		} else if strings.Contains(cmdOut, "200") {
			return shared.SuccessMsg{}
		} else if cmdOut == "" {
			return shared.ErrorMsg{}
		}

		return shared.SetMsg(cmdOut)
	}
}
