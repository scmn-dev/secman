package constants

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const SMUI_PORT = "3750"

const PRIMARY_COLOR = "#1163E6"
const SECONDARY_COLOR = "#B4B4B4"
const GREEN_COLOR = "#04B575"
const RED_COLOR = "#FF4141"
const YELLOW_COLOR = "178" 

var GRAY_COLOR = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}
var WHITE_COLOR = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
var INACTIVE_COLOR = lipgloss.AdaptiveColor{Light: "#CACACA", Dark: "#4F4F4F"}
var SUBTITLE_COLOR = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}

const LIST_HEIGHT = 14
const LIST_WIDTH = 10
const SECMAN_LIST_HEIGHT = 34
const SECMAN_LIST_WIDTH = 14
const LIST_HEADER_HEIGHT = 6
const LIST_FOOTER_HEIGHT = 3
const LIST_PROPORTION = 0.3

var TABLE_BORDER_DESIGN = table.Border{
	Top:    "─",
	Left:   "│",
	Right:  "│",
	Bottom: "─",

	TopRight:    "╮",
	TopLeft:     "╭",
	BottomRight: "╯",
	BottomLeft:  "╰",

	TopJunction:    "┬",
	LeftJunction:   "├",
	RightJunction:  "┤",
	BottomJunction: "┴",
	InnerJunction:  "┼",

	InnerDivider: "│",
}

func Logo(text string) string {
	return "\n" + lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff")).
		Background(lipgloss.Color(PRIMARY_COLOR)).
		Padding(0, 1).
		SetString(text).
		String()
} 

var HelpExamples = fmt.Sprintf(`
%s

$ secman init
%s

$ secman auth
%s

$ secman insert --[PASSWORD_TYPE]
%s

$ secman .
%s

$ secman read --[PASSWORD_TYPE] <PASSWORD_NAME>
%s

$ secman edit --[PASSWORD_TYPE] <PASSWORD_NAME>
%s

$ secman delete --[PASSWORD_TYPE] <PASSWORD_NAME>
%s

$ secman generate
`, Logo("Initialize ~/.secman"), Logo("Authorize With Secman"), Logo("Insert a New Password"), Logo("List All Passwords"), Logo("Read a Password"), Logo("Edit a Password"), Logo("Delete"), Logo("Generate"))

var (
	homeDir, homeErr = dfs.GetHomeDirectory()
	DotSecmanPath = filepath.Join(homeDir, ".secman")
	SecmanConfigPath = filepath.Join(DotSecmanPath, "secman.json")
	SMUIPath = filepath.Join(DotSecmanPath, "ui")
	SMUIIndexPath = filepath.Join(SMUIPath, "index.html")
	secmanConfig, SmErr = ioutil.ReadFile(SecmanConfigPath)
)

func SecmanConfig() []byte {
	if SmErr != nil {
		return []byte("")
	}

	return secmanConfig
}

var (
	checkmarkChar = lipgloss.NewStyle().
		SetString("✔").
		Foreground(lipgloss.Color(GREEN_COLOR))

	xChar = lipgloss.NewStyle().
		SetString("✘").
		Foreground(lipgloss.Color(RED_COLOR))

	Checkmark = "[" + checkmarkChar.String() + "] "
	X = "[" + xChar.String() + "] "
)
