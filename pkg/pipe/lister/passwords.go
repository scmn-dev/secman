package lister

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abdfnx/bubbles/list"
	"github.com/abdfnx/gosh"
	"github.com/briandowns/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/internal/shared"
	"github.com/spf13/viper"
)

type Passwords struct {
    logins  []list.Item
	ccs     []list.Item
	emails  []list.Item
	notes   []list.Item
	servers []list.Item
}

func SPW() *Passwords {
	return &Passwords{
		logins:  readPasswords("-l"),
		ccs:     readPasswords("-c"),
		emails:  readPasswords("-e"),
		notes:   readPasswords("-n"),
		servers: readPasswords("-s"),
	}
}

func readPasswords(p string) []list.Item {
	if len(os.Args) > 1 && len(os.Args) != 3 {
		if strings.Contains(os.Args[1], "list") || strings.Contains(os.Args[1], ".") {
			s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
			s.Suffix = " 📡 Preparing & Getting data..."
			s.Start()

			err, out, errout := gosh.RunOutput("scc . " + p)

			if err != nil {
				fmt.Println(err)
				fmt.Println(errout)

				os.Exit(2)
			}

			s.Stop()

			if strings.Contains(out, "401") {
				st := shared.DefaultStyles()

				head := st.Error.Render("\n\nYour Authentication is Expired.")
				body := st.Subtle.Render(" Refresh your authentication via `secman auth refresh`.")

				fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Lister") + st.Wrap.Render(head + body)).String())

				os.Exit(2)
			}

			viper.SetConfigType("yaml")
			viper.ReadConfig(bytes.NewBuffer([]byte(out)))
			items := make([]list.Item, 0)

			for _, line := range strings.Split(viper.GetString("passwords"), "\n") {
				if line[0] == '#' || len(line) == 0 {
					continue
				} else {
					fmt.Println("Passwords Not Found")
					os.Exit(2)
				}

				fields := strings.Split(line, "-|-")

				if p == "-l" {
					if len(fields) != 5 {
						continue
					}
			
					items = append(items, NewLoginListItem(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2]), strings.TrimSpace(fields[3]), strings.TrimSpace(fields[4])))
				} else if p == "-c" {
					if len(fields) != 6 {
						continue
					}

					items = append(items, NewCCListItem(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2]), strings.TrimSpace(fields[3]), strings.TrimSpace(fields[4]), strings.TrimSpace(fields[5])))
				} else if p == "-e" {
					if len(fields) != 3 {
						continue
					}

					items = append(items, NewEmailListItem(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2])))
				} else if p == "-n" {
					items = append(items, NewNoteListItem(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2])))
				} else if p == "-s" {
					if len(fields) != 10 {
						continue
					}

					items = append(items, NewServerListItem(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]), strings.TrimSpace(fields[2]), strings.TrimSpace(fields[3]), strings.TrimSpace(fields[4]), strings.TrimSpace(fields[5]), strings.TrimSpace(fields[6]), strings.TrimSpace(fields[7]), strings.TrimSpace(fields[8]), strings.TrimSpace(fields[9])))
				}
			}

			return items
		}
	}

	return nil
}

func (c *Passwords) PWs(p string) []list.Item {
    if p == "-l" {
		return c.logins
	} else if p == "-c" {
		return c.ccs
	} else if p == "-e" {
		return c.emails
	} else if p == "-n" {
		return c.notes
	} else if p == "-s" {
		return c.servers
	}

	return nil
}
