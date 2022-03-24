package generator

import (
    "fmt"
    "time"
    "strings"
    "math/rand"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/constants"
	"github.com/scmn-dev/secman/pkg/options"
)

func Generator(opts *options.GenOptions) {
    if opts.Raw {
		fmt.Println(SMGenerator(opts.Length))
	} else {
		out, err := glamour.Render("> " + SMGenerator(opts.Length), "dark")

		if err != nil {
			fmt.Println("could not render")
			return
		}

		out = out[:len(out)-1]

		fmt.Print(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Generator")).String() + out)
	}
}

func SMGenerator(length int) string {
	rand.Seed(time.Now().Unix())
    charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP1234567890")

    var output strings.Builder

	for i := 0; i < length; i++ {
        random := rand.Intn(len(charSet))
        randomChar := charSet[random]
        output.WriteRune(randomChar)
    }

	return output.String()
}
