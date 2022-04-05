package shared

import (
	"github.com/abdfnx/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/constants"
)

type Styles struct {
	BaseStyle,
	Cursor,
	Tab,
	ActiveTab,
	ListTitle,
	Wrap,
	Doc,
	Divider,
	Help,
	Item,
	Subtle,
	Success,
	Error,
	Bold,
	Prompt,
	ListView,
	FilterPrompt,
	FilterCursor,
	FocusedPrompt,
	InactivePagination,
	Pagination,
	SelectionMarker,
	SelectedMenuItem,
	Checkmark lipgloss.Style
}

var ActiveTabBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      " ",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "┘",
	BottomRight: "└",
}

var TabBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "┴",
	BottomRight: "┴",
}

func DefaultStyles() Styles {
	s := Styles{}

	s.Cursor = lipgloss.NewStyle().Foreground(primary)
	s.Wrap = lipgloss.NewStyle().Width(58)
	s.Checkmark = lipgloss.NewStyle().
		SetString("✔ ").
		Foreground(lipgloss.Color(constants.GREEN_COLOR))
	s.Subtle = lipgloss.NewStyle().
		Foreground(constants.SUBTITLE_COLOR)
	s.Success = lipgloss.NewStyle().Foreground(lipgloss.Color(constants.GREEN_COLOR))
	s.Error = lipgloss.NewStyle().Foreground(lipgloss.Color(constants.RED_COLOR))
	s.Bold = lipgloss.NewStyle().Bold(true)
	s.Prompt = lipgloss.NewStyle().MarginRight(1).SetString(">")
	s.FocusedPrompt = s.Prompt.Copy().Foreground(primary)
	s.FilterPrompt = lipgloss.NewStyle().Foreground(lipgloss.Color(constants.PRIMARY_COLOR))
	s.FilterCursor = lipgloss.NewStyle().Foreground(lipgloss.Color(constants.PRIMARY_COLOR))
	s.InactivePagination = lipgloss.NewStyle().
		Foreground(constants.INACTIVE_COLOR)
	s.SelectionMarker = lipgloss.NewStyle().
		Foreground(primary).
		PaddingRight(1).
		SetString(">")
	s.SelectedMenuItem = lipgloss.NewStyle().PaddingLeft(2).Foreground(primary)
	s.BaseStyle = lipgloss.NewStyle()
	s.Item = lipgloss.NewStyle().PaddingLeft(4)
	s.Pagination = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	s.Doc = lipgloss.NewStyle().Align(lipgloss.Center)
	s.Help = list.DefaultStyles().HelpStyle.PaddingLeft(4)
	s.Tab = lipgloss.NewStyle().
		Border(TabBorder, true).
		BorderForeground(lipgloss.Color(constants.PRIMARY_COLOR)).
		Padding(0, 2)
	s.ActiveTab = s.Tab.Copy().Border(ActiveTabBorder, true).Bold(true)
	s.ListTitle = lipgloss.NewStyle().Background(lipgloss.Color(constants.PRIMARY_COLOR)).Padding(0, 1).Foreground(constants.WHITE_COLOR)
	s.ListView = lipgloss.NewStyle().
			PaddingRight(5).
			PaddingLeft(5).
			MarginRight(1).
			Border(lipgloss.RoundedBorder(), false, true, false, false)
	s.Divider = lipgloss.NewStyle()

	return s
}

var (
	primary = lipgloss.AdaptiveColor{Light: constants.PRIMARY_COLOR, Dark: constants.PRIMARY_COLOR}
	spinnerStyle = lipgloss.NewStyle().Foreground(constants.GRAY_COLOR)

	blurredButtonStyle = lipgloss.NewStyle().
		Foreground(constants.WHITE_COLOR).
		Background(lipgloss.AdaptiveColor{Light: constants.SECONDARY_COLOR, Dark: constants.SECONDARY_COLOR}).
		Padding(0, 3)

	focusedButtonStyle = blurredButtonStyle.Copy().
		Background(primary)
)

func OKButton(focused, defaultButton bool, value string) string {
	return styledButton(value, defaultButton, focused)
}

func CancelButton(focused, defaultButton bool, value string) string {
	return styledButton(value, defaultButton, focused)
}

func styledButton(str string, underlined, focused bool) string {
	var st lipgloss.Style

	if focused {
		st = focusedButtonStyle.Copy()
	} else {
		st = blurredButtonStyle.Copy()
	}

	if underlined {
		st = st.Underline(true)
	}

	return st.Render(str)
}

func NewSpinner() spinner.Model {
	s := spinner.NewModel()

	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return s
}
