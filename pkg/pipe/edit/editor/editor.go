package editor

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdfnx/gosh"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/internal/shared"
)

const (
	textInput shared.Index = iota
	okButton
	cancelButton
)

type model struct {
	styles   shared.Styles
	state    shared.State
	password string
	index    shared.Index
	message  string
	value    string
	input    textinput.Model
	spinner  spinner.Model
	pwType   string
	field    string
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

func Editor(pwType, field, password string) model {
	st := shared.DefaultStyles()

	t := textinput.NewModel()
	t.CursorStyle = st.Cursor
	t.Placeholder = "Enter the new value of " + password + " " + field
	t.Prompt = st.FocusedPrompt.String()
	t.CharLimit = 50

	if strings.Contains(field, "Password") || field == "Verification Number" {
		t.EchoMode = textinput.EchoPassword
		t.EchoCharacter = '•'
	}

	t.Focus()

	m := model{
		styles:   st,
		state:    shared.Ready,
		password: fmt.Sprintf("\"%s\"", password),
		index:    textInput,
		message:  "",
		value:    "",
		input:    t,
		spinner:  shared.NewSpinner(),
		pwType:   pwType,
		field:    fmt.Sprintf("\"%s\"", field),
	}

	return m
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
									m.value = strings.TrimSpace(m.input.Value())

									return m, tea.Batch(
										sme(m),
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
			m.message = m.styles.Subtle.Render("Sorry, We couldn't update your password, maybe the password name, field, or value is wrong.\n")

			return m, nil
		
		case shared.OtherMsg:
			m.state = shared.Ready
			head := m.styles.Error.Render("Your Authentication is Expired.")
			body := m.styles.Subtle.Render(" Refresh your authentication via `secman auth refresh`.")
			m.message = m.styles.Wrap.Render(head + body)

			return m, nil

		case shared.SuccessMsg:
			m.state = shared.Ready
			m.message = m.styles.Error.Render(m.styles.Checkmark.String() + "Password Updated")

			return m, nil

		case shared.Message:
			m.state = shared.Ready
			head := m.styles.Error.Render("Oh, what? There was a curious error we were not expecting. ")
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
	s := m.input.View() + "\n\n"

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

	return lipgloss.NewStyle().Padding(0, 2).SetString(s).String()
}

func spinnerView(m model) string {
	return m.spinner.View() + "✏️ Editing..."
}

func sme(m model) tea.Cmd {
	return func() tea.Msg {
		err, out, errout := gosh.RunOutput("scc edit " + m.pwType + " " + m.password + " -f " +  fmt.Sprintf("\"%s\"", m.field) + " -v " + fmt.Sprintf("\"%s\"", m.value))

		if err != nil {
			fmt.Println(errout)
			return shared.Message{err}
		}

		cmdOut := strings.TrimSuffix(out, "\n")

		if strings.Contains(cmdOut, "200") {
			return shared.SuccessMsg{}
		} else if strings.Contains(cmdOut, "401") {
			return shared.OtherMsg{}
		} else if strings.Contains(cmdOut, "404") {	
			return shared.ErrorMsg{}
		}

		return shared.SetMsg(cmdOut)
	}
}
