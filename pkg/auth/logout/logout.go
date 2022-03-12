package logout

import (
	"os"
	"fmt"

	"github.com/abdfnx/gosh"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/scmn-dev/secman/internal/config"
	"github.com/scmn-dev/secman/internal/shared"
)

const (
	okButton shared.Index = iota
	cancelButton
)

var user = config.Config("config.name")

type model struct {
	styles  shared.Styles
	state   shared.State
	index   shared.Index
	message string
	spinner spinner.Model
}

func (m *model) indexForward() {
	m.index = (m.index + 1) % 2
}

func (m *model) indexBackward() {
	m.index = (m.index - 1) % 2
}

func Logout() model {
	st := shared.DefaultStyles()

	return model{
		styles:  st,
		state:   shared.Ready,
		index:   okButton,
		message: "",
		spinner: shared.NewSpinner(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
										sml(m),
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
			m.message = m.styles.Error.Render(m.styles.Checkmark.String() + "Logged out successfully")

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
}

func (m model) View() string {
	if user == "" {
		fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Auth") + m.styles.Error.Render("\n\nYou are not logged in. Please use ") + m.styles.Subtle.Render("`secman auth login`") + m.styles.Error.Render(" command to login.")))

		os.Exit(0)

		return ""
	} else {
		username := m.styles.Bold.Render(user)
		s := "\n\nAre you sure you want to logout of secman " + username + "?\n\n"

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

		return lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Auth") + s).String()
	}
}

func spinnerView(m model) string {
	return m.spinner.View() + "🔓 Logging out..."
}

func sml(m model) tea.Cmd {
	return func() tea.Msg {
		gosh.Run("sc logout")

		return shared.SuccessMsg{}
	}
}
