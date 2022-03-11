package insert

import (
	"os"
	"fmt"
	"log"
	"strings"

	"github.com/abdfnx/gosh"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/pkg/options"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/scmn-dev/secman/internal/shared"
)

const okButton shared.Index = iota

type model struct {
	styles     shared.Styles
	focusIndex int
	index      shared.Index
	inputs     []textinput.Model
	spinner    spinner.Model
	state      shared.State
	pwType     string
	message    string
}

func Insert(o *options.PasswordsOptions) model {
	st := shared.DefaultStyles()

	var inps = func() int {
		if o.Logins {
			return 5
		} else if o.CreditCards {
			return 6
		} else if o.Emails {
			return 3
		} else if o.Notes {
			return 2
		} else if o.Servers {
			return 10
		}

		return 0
	}

	m := model{
		styles:  st,
		inputs:  make([]textinput.Model, inps()),
		spinner: shared.NewSpinner(),
		state:   shared.Ready,
		pwType:  shared.PasswordType(o),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = st.Cursor
		t.Prompt = st.FocusedPrompt.String()
		t.CharLimit = 32

		if o.Logins {
			switch i {
				case 0:
					t.Focus()
					t.Placeholder = "Title"
					t.CharLimit = 50

				case 1:
					t.Placeholder = "URL"
					t.CharLimit = 50
				
				case 2:
					t.Placeholder = "Username"
					t.CharLimit = 50

				case 3:
					t.Placeholder = "Password"
					t.EchoMode = textinput.EchoPassword
					t.EchoCharacter = '•'

				case 4:
					t.Placeholder = "Extra"
					t.CharLimit = 50
			}
		} else if o.CreditCards {
			switch i {
				case 0:
					t.Focus()
					t.Placeholder = "Title"
					t.CharLimit = 50

				case 1:
					t.Placeholder = "Card Holder"
					t.CharLimit = 50

				case 2:
					t.Placeholder = "Card Type"
					t.CharLimit = 50

				case 3:
					t.Placeholder = "Card Number"
					t.CharLimit = 16

				case 4:
					t.Placeholder = "Expiry Date"
					t.CharLimit = 5

				case 5:
					t.Placeholder = "Verification Number"
					t.CharLimit = 3
			}
		} else if o.Emails {
			switch i {
				case 0:
					t.Focus()
					t.Placeholder = "Title"
					t.CharLimit = 50

				case 1:
					t.Placeholder = "Email"
					t.CharLimit = 50

				case 2:
					t.Placeholder = "Password"
					t.EchoMode = textinput.EchoPassword
					t.EchoCharacter = '•'
			}
		} else if o.Notes {
			switch i {
				case 0:
					t.Focus()
					t.Placeholder = "Title"
					t.CharLimit = 50

				case 1:
					t.Placeholder = "Note"
					t.CharLimit = 50
			}
		} else if o.Servers {
			switch i {
				case 0:
					t.Focus()
					t.Placeholder = "Title"
					t.CharLimit = 50

				case 1:
					t.Placeholder = "Ip Address"
					t.CharLimit = 50

				case 2:
					t.Placeholder = "Username"
					t.CharLimit = 50

				case 3:
					t.Placeholder = "Password"
					t.EchoMode = textinput.EchoPassword
					t.EchoCharacter = '•'

				case 4:
					t.Placeholder = "URL"
					t.CharLimit = 50

				case 5:
					t.Placeholder = "Hosting Username"
					t.CharLimit = 50

				case 6:
					t.Placeholder = "Hosting Password"
					t.EchoMode = textinput.EchoPassword
					t.EchoCharacter = '•'

				case 7:
					t.Placeholder = "Admin Username"
					t.CharLimit = 50

				case 8:
					t.Placeholder = "Admin Password"
					t.EchoMode = textinput.EchoPassword
					t.EchoCharacter = '•'

				case 9:
					t.Placeholder = "Extra"
					t.CharLimit = 50
			}
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
				case tea.KeyCtrlC:
					return m, tea.Quit

				default:
					if m.state == shared.Loading {
						return m, nil
					}

					switch msg.String() {
					case "tab", "shift+tab", "enter", "up", "down":
						s := msg.String()
			
						if s == "enter" {
							if m.index == okButton {
								m.state = shared.Loading
								m.message = ""

								return m, tea.Batch(
									smi(m),
									spinner.Tick,
								)
							} else {
								return m, tea.Quit
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
			m.message = m.styles.Success.Render("🔑 Password created successfully!")

			return m, nil

		case shared.OtherMsg:
			m.state = shared.Ready
			head := m.styles.Error.Render("Your Authentication is Expired.")
			body := m.styles.Subtle.Render(" Refresh your authentication via `secman auth refresh`.\n")
			m.message = m.styles.Wrap.Render(head + body)

			return m, nil

		case shared.ErrorMsg:
			m.state = shared.Ready
			m.message = m.styles.Error.Render("An error occurred while creating your password. Please try again.")

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
	s := "\n\nCreate a new password for your account.\n\n"

	if m.state == shared.Loading {
		s += spinnerView(m)
	} else {
		for i := range m.inputs {
			s += m.inputs[i].View() + "\n"
		}
	
		s += "\n" + shared.OKButton(m.focusIndex == len(m.inputs), true, "Ok")

		if m.message != "" {
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(m.message).String())
			os.Exit(0)
		}
	}

	return lipgloss.NewStyle().Padding(0, 2).SetString(constants.Logo("Secman Password Creator") + s).String()
}

func spinnerView(m model) string {
	return m.spinner.View() + "🔑 Creating..."
}

func smi(m model) tea.Cmd {
	return func() tea.Msg {
		insCmd := "sc insert "
		cmd := ""
		extra := "no-extra"

		if m.pwType == "-l" {
			if m.inputs[4].Value() != "" {
				extra = m.inputs[4].Value()
			}
		} else if m.pwType == "-s" {
			if m.inputs[9].Value() != "" {
				extra = m.inputs[9].Value()
			}
		}

		var CMD = func() string {
			if m.pwType == "-l" {
				cmd = insCmd + m.pwType + " -t " + fmt.Sprintf("\"%s\"", m.inputs[0].Value()) + " -u " + fmt.Sprintf("\"%s\"", m.inputs[1].Value()) + " -U " + fmt.Sprintf("\"%s\"", m.inputs[2].Value()) + " -p " + fmt.Sprintf("\"%s\"", m.inputs[3].Value()) + " -x " + extra
			} else if m.pwType == "-c" {
				cmd = insCmd + m.pwType + " -t " + fmt.Sprintf("\"%s\"", m.inputs[0].Value()) + " -C " + fmt.Sprintf("\"%s\"", m.inputs[1].Value()) + " -T " + fmt.Sprintf("\"%s\"", m.inputs[2].Value()) + " -N " + fmt.Sprintf("\"%s\"", m.inputs[3].Value()) + " -E " + fmt.Sprintf("\"%s\"", m.inputs[4].Value()) + " -V " + fmt.Sprintf("\"%s\"", m.inputs[5].Value())
			} else if m.pwType == "-e" {
				cmd = insCmd + m.pwType + " -t " + fmt.Sprintf("\"%s\"", m.inputs[0].Value()) + " -m " + fmt.Sprintf("\"%s\"", m.inputs[1].Value()) + " -p " + fmt.Sprintf("\"%s\"", m.inputs[2].Value())
			} else if m.pwType == "-n" {
				cmd = insCmd + m.pwType + " -t " + fmt.Sprintf("\"%s\"", m.inputs[0].Value()) + " -N " + fmt.Sprintf("\"%s\"", m.inputs[1].Value())
			} else if m.pwType == "-s" {
				cmd = insCmd + m.pwType + " -t " + fmt.Sprintf("\"%s\"", m.inputs[0].Value()) + " -i " + fmt.Sprintf("\"%s\"", m.inputs[1].Value()) + " -U " + fmt.Sprintf("\"%s\"", m.inputs[2].Value()) + " -p " + fmt.Sprintf("\"%s\"", m.inputs[3].Value()) + " -u " + fmt.Sprintf("\"%s\"", m.inputs[4].Value()) + " -H " + fmt.Sprintf("\"%s\"", m.inputs[5].Value()) + " -O " + fmt.Sprintf("\"%s\"", m.inputs[6].Value()) + " -a " + fmt.Sprintf("\"%s\"", m.inputs[7].Value()) + " -w " + fmt.Sprintf("\"%s\"", m.inputs[8].Value()) + " -x " + fmt.Sprintf("\"%s\"", m.inputs[9].Value())
			}

			return cmd
		}

		err, out, errout := gosh.RunOutput(CMD())

		if err != nil {
			log.Printf("error: %v\n", err)
			fmt.Print(errout)
		}

		out = strings.TrimSuffix(out, "\n")

		if strings.Contains(out, "401") {
			return shared.OtherMsg{}
		} else if strings.Contains(out, "404") {
			return shared.ErrorMsg{}
		} else if strings.Contains(out, "200") {
			return shared.SuccessMsg{}
		} 

		return shared.SetMsg(out)
	}
}
