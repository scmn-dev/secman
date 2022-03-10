package refresh

import (
	"os"
	"fmt"
	"strings"

	"github.com/abdfnx/gosh"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/scmn-dev/secman/internal/config"
	"github.com/scmn-dev/secman/internal/shared"
)

const (
	textInput shared.Index = iota
	okButton
	cancelButton
)

var user = config.Config("config.user")

type model struct {
	styles  shared.Styles
	state   shared.State
	ms      string
	index   shared.Index
	message  string
	input   textinput.Model
	spinner spinner.Model
}

func (m *model) updateFocus() {
	if m.index == textInput && !m.input.Focused() {
		m.input.Focus()
		m.input.Prompt = m.styles.FocusedPrompt.String()
	} else if m.index != textInput && m.input.Focused() {
		m.input.Blur()
		m.input.Prompt = m.styles.Prompt.String()
	}
}

func (m *model) indexForward() {
	m.index++
	if m.index > cancelButton {
		m.index = textInput
	}

	m.updateFocus()
}

func (m *model) indexBackward() {
	m.index--
	if m.index < textInput {
		m.index = cancelButton
	}

	m.updateFocus()
}

func Refresh() model {
	st := shared.DefaultStyles()
	t := textinput.NewModel()

	t.CursorStyle = st.Cursor
	t.Placeholder = "Your Master Password"
	t.Prompt = st.FocusedPrompt.String()
	t.CharLimit = 50
	t.EchoMode = textinput.EchoPassword
	t.EchoCharacter = '•'
	t.Focus()

	return model{
		styles:  st,
		state:   shared.Ready,
		ms: "",
		index:   textInput,
		message:  "",
		input:   t,
		spinner: shared.NewSpinner(),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
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
							if m.index != textInput {
								m.indexForward()
							}

						case "h", "j", "left":
							if m.index != textInput {
								m.indexBackward()
							}

						case "up", "down":
							if m.index == textInput {
								m.indexForward()
							} else {
								m.index = textInput
								m.updateFocus()
							}

						case "enter":
							switch m.index {
								case textInput:
									fallthrough

								case okButton:
									m.state = shared.Loading
									m.message = ""
									m.ms = strings.TrimSpace(m.input.Value())

									return m, tea.Batch(
										smr(m),
										spinner.Tick,
									)

								case cancelButton:
									return m, tea.Quit
						}
					}

					if m.index == textInput {
						var cmd tea.Cmd
						m.input, cmd = m.input.Update(msg)

						return m, cmd
					}

					return m, nil
				}

			case shared.ErrorMsg:
				m.state = shared.Ready
				m.message = m.styles.Error.Render("Invalid master password. or if you don't have an account, please create one using the command ") + m.styles.Subtle.Render("`secman auth create`")

				return m, nil

			case shared.SuccessMsg:
				m.state = shared.Ready
				m.message = m.styles.Wrap.Render(m.styles.Success.Render("🔗 Refreshed"))

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
				m.input, cmd = m.input.Update(msg)

				return m, cmd
	}
}

func (m model) View() string {
	if user == "" {
		fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Auth") + m.styles.Error.Render("\n\nYou are not logged in. Please use ") + m.styles.Subtle.Render("`secman auth`") + m.styles.Error.Render(" command to login.")))

		os.Exit(0)

		return ""
	} else {
		s := "\n\nPlease enter your Master Password\n\n"
		s += m.input.View() + "\n\n"

		if m.state == shared.Loading {
			s += spinnerView(m)
		} else {
			s += shared.OKButton(m.index == 1, true, "Yes")
			s += " " + shared.CancelButton(m.index == 2, false, "Exit")

			if m.message != "" {
				fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(m.message).String())
				os.Exit(0)
			}
		}

		return lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Auth") + s).String()
	}
}

func spinnerView(m model) string {
	return m.spinner.View() + "🔗 Refreshing..."
}

func smr(m model) tea.Cmd {
	return func() tea.Msg {
		err, out, _ := gosh.RunOutput("sc auth -e " + user + " -m " + m.ms)

		if err != nil {
			return shared.Message{err}
		}

		out = strings.TrimSuffix(out, "\n")

		if strings.Contains(out, "401") {		
			return shared.ErrorMsg{}
		} else if strings.Contains(out, "406") {
			return shared.SuccessMsg{}
		}

		return shared.SetMsg(out)
	}
}
