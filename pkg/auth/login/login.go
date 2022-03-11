package login

import (
	"os"
	"fmt"
	"log"
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
	verfiyButton
	cancelButton
)

type model struct {
	styles     shared.Styles
	focusIndex int
	index      shared.Index
	inputs     []textinput.Model
	spinner    spinner.Model
	state      shared.State
	message    string
}

func Login() model {
	st := shared.DefaultStyles()
	var t textinput.Model

	m := model{
		styles:  st,
		inputs:  make([]textinput.Model, 2),
		spinner: shared.NewSpinner(),
		state:   shared.Ready,
		message: "",
	}

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = st.Cursor
		t.Prompt = st.FocusedPrompt.String()
		t.CharLimit = 32

		switch i {
			case 0:
				t.Focus()
				t.Placeholder = "E-Mail"
				t.CharLimit = 64

			case 1:
				t.Placeholder = "Master Password"
				t.EchoMode = textinput.EchoPassword
				t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.updateInputs(msg)

	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyCtrlC, tea.KeyCtrlQ:
					return m, tea.Quit

				default:
					if m.state == shared.Loading {
						return m, nil
					}

					switch msg.String() {
						case "tab", "shift+tab", "enter", "up", "down":
							s := msg.String()

							if s == "enter" {
								if m.focusIndex == 2 {
									m.state = shared.Loading
									m.message = ""
					
									return m, tea.Batch(
										sma(m),
										spinner.Tick,
									)
								}
							}
				
							if s == "up" || s == "shift+tab" {
								m.focusIndex--
							} else {
								m.focusIndex++
							}
				
							if m.focusIndex > len(m.inputs) {
								m.focusIndex = 0
							} else if m.focusIndex < 0 {
								m.focusIndex = len(m.inputs)
							}
				
							cmds := make([]tea.Cmd, len(m.inputs))

							for i := 0; i <= len(m.inputs)-1; i++ {
								if i == m.focusIndex {
									cmds[i] = m.inputs[i].Focus()
									continue
								}

								m.inputs[i].Blur()
							}

							return m, tea.Batch(cmds...)
						}

					return m, nil
			}

		case shared.SuccessMsg:
			m.state = shared.Ready
			m.message = m.styles.Wrap.Render(m.styles.Success.Render("🎉 Welcome back!"))

			return m, nil

		case shared.OtherMsg:
			m.state = shared.Ready
			username := m.styles.Bold.Render(config.Config("config.name"))
			m.message = "You are already logged in as " + username + ", if you want to re-authenticate run " + m.styles.Subtle.Render("`secman auth refresh`")

			return m, nil

		case shared.ErrorMsg:
			m.state = shared.Ready
			m.message = m.styles.Error.Render("Invalid email or master password. if you don't have an account, please create one using the command ") + m.styles.Subtle.Render("`secman auth create`")

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
		}

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	s := "\n\nEnter your secman account credentials. press `Ctrl+q` to quit\n\n"

	if m.state == shared.Loading {
		s += spinnerView(m)
	} else {
		for i := range m.inputs {
			s += m.inputs[i].View() + "\n"
		}

		s += "\n" + shared.OKButton(m.focusIndex == 2, true, "Yes")

		if m.message != "" {
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(m.message).String())
			os.Exit(0)
		}
	}

	return lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Auth") + s).String()
}

func spinnerView(m model) string {
	return m.spinner.View() + "🔒 Authenticating..."
}

func sma(m model) tea.Cmd {
	return func() tea.Msg {
		err, out, errout := gosh.RunOutput("sc auth -e " + m.inputs[0].Value() + " -m " + m.inputs[1].Value())

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		}

		out = strings.TrimSuffix(out, "\n")
		
		if strings.Contains(out, "401") {
			return shared.ErrorMsg{}
		} else if strings.Contains(out, "406") {
			return shared.OtherMsg{}
		} else if strings.Contains(out, "200") {
			return shared.SuccessMsg{}
		}

		return shared.SetMsg(out)
	}
}
