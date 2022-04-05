package lister

import (
	"github.com/abdfnx/bubbles/list"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/internal/shared"
)

type model struct {
	width, height 	int
	styles 			shared.Styles

	loginsList 		list.Model
	creditCardsList list.Model
	emailsList 		list.Model
	notesList 		list.Model
	serversList 	list.Model

	viewport  		viewport.Model
	spinner   		spinner.Model
	state    		int
}

const (
	LOGIN = iota
	CC
	EMAIL
	NOTE
	SERVER
)

var (
	st = shared.DefaultStyles()
	states = []int{LOGIN, CC, EMAIL, NOTE, SERVER}
)

type Config interface {
	PWs(p string) []list.Item
}

var conf Config = SPW()

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick)
}

func (m *model) handleKeys(msg tea.KeyMsg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg.Type {
		case tea.KeyCtrlC:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		case tea.KeyTab:
			m.switchState(+1)
		default:
			m.loginsList, cmd = m.loginsList.Update(msg)
			cmds = append(cmds, cmd)
	
			m.creditCardsList, cmd = m.creditCardsList.Update(msg)
			cmds = append(cmds, cmd)
	
			m.emailsList, cmd = m.emailsList.Update(msg)
			cmds = append(cmds, cmd)
	
			m.notesList, cmd = m.notesList.Update(msg)
			cmds = append(cmds, cmd)

			m.serversList, cmd = m.serversList.Update(msg)
			cmds = append(cmds, cmd)
	
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
	
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.width, m.height = msg.Width, msg.Height
			height := m.height

			listViewWidth := int(constants.LIST_PROPORTION * float64(m.width))
			listWidth := listViewWidth - st.ListView.GetHorizontalFrameSize()
			m.loginsList.SetSize(listWidth, height)
			m.creditCardsList.SetSize(listWidth, height)
			m.emailsList.SetSize(listWidth, height)
			m.notesList.SetSize(listWidth, height)
			m.serversList.SetSize(listWidth, height)

			detailViewWidth := m.width - listViewWidth
			m.viewport = viewport.New(detailViewWidth, height)
			m.viewport.SetContent(m.detailView())

		case tea.KeyMsg:
			cmds = append(cmds, m.handleKeys(msg))
	}

	return m, tea.Batch(cmds...)
}

func Lister() *model {
	var tab = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch"),
	)

	loginPWs := conf.PWs("-l")
	ccPWs := conf.PWs("-c")
	emailPWs := conf.PWs("-e")
	notePWs := conf.PWs("-n")
	serverPWs := conf.PWs("-s")

	l := list.NewModel(loginPWs, list.NewDefaultDelegate(), constants.SECMAN_LIST_WIDTH, constants.SECMAN_LIST_HEIGHT)
	l.Title = "Logins List"
	l.Styles.Title = st.ListTitle
	l.SetShowFilter(false)
	l.SetFilteringEnabled(false)

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{tab}
	}

	c := list.NewModel(ccPWs, list.NewDefaultDelegate(), constants.SECMAN_LIST_WIDTH, constants.SECMAN_LIST_HEIGHT)
	c.Title = "Credit Cards List"
	c.Styles.Title = st.ListTitle
	c.SetShowFilter(false)
	c.SetFilteringEnabled(false)
	c.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{tab}
	}
	
	e := list.NewModel(emailPWs, list.NewDefaultDelegate(), constants.SECMAN_LIST_WIDTH, constants.SECMAN_LIST_HEIGHT)
	e.Title = "Emails List"
	e.Styles.Title = st.ListTitle
	e.SetShowFilter(false)
	e.SetFilteringEnabled(false)
	e.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{tab}
	}

	n := list.NewModel(notePWs, list.NewDefaultDelegate(), constants.SECMAN_LIST_WIDTH, constants.SECMAN_LIST_HEIGHT)
	n.Title = "Notes List"
	n.Styles.Title = st.ListTitle
	n.SetShowFilter(false)
	n.SetFilteringEnabled(false)
	n.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{tab}
	}

	r := list.NewModel(serverPWs, list.NewDefaultDelegate(), constants.SECMAN_LIST_WIDTH, constants.SECMAN_LIST_HEIGHT)
	r.Title = "Servers List"
	r.Styles.Title = st.ListTitle
	r.SetShowFilter(false)
	r.SetFilteringEnabled(false)
	r.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{tab}
	}

	s := spinner.New()
	s.Spinner = spinner.Dot

	m := &model{
		styles: 	     st,
		loginsList:      l,
		creditCardsList: c,
		emailsList:      e,
		notesList:       n,
		serversList:     r,
		spinner:         s,
		state:           LOGIN,
	}

	return m
}

func (m *model) switchState(direction int) {
	newState := (m.state + direction) % len(states)

	m.state = newState
}
