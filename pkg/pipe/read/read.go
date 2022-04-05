package read

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abdfnx/gosh"
	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/scmn-dev/secman/v6/constants"
	"github.com/scmn-dev/secman/v6/internal/shared"
	"github.com/scmn-dev/secman/v6/pkg/options"
	"github.com/tidwall/gjson"
)

type model struct {
	styles  shared.Styles
	smTable table.Model
}

func Read(o *options.PasswordsOptions) model {
	st := shared.DefaultStyles()

	m := model{
		styles: st,
		smTable: table.New([]table.Column{
			table.NewColumn("Data", "Data", 9).WithStyle(
				lipgloss.NewStyle().
					Align(lipgloss.Center),
			),
		}).WithRows([]table.Row{
			table.NewRow(table.RowData{
				"Data": "No data",
			}),
		}).Border(constants.TABLE_BORDER_DESIGN),
	}

	var p = fmt.Sprintf("\"%s\"", o.Password)

	var isHidden = func () string {
		if o.ShowHidden {
			return "-p"
		}

		return ""
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 📡 Preparing & Getting data..."
	s.Start()

	var err, out, errout = gosh.RunOutput("scc read " + shared.PasswordType(o) + " " + p + " " + isHidden())

	if err != nil {
		fmt.Println(err)
		fmt.Println(errout)
		os.Exit(0)
	}

	s.Stop()

	if strings.Contains(out, "401") {
		head := m.styles.Error.Render("\n\nYour Authentication is Expired.")
		body := m.styles.Subtle.Render(" Refresh your authentication via `secman auth refresh`.")

		fmt.Println(lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Reader") + m.styles.Wrap.Render(head + body)).String())

		os.Exit(0)
	} else if o.Logins {
		title := gjson.Get(out, "title")
		url := gjson.Get(out, "url")
		username := gjson.Get(out, "username")
		password := gjson.Get(out, "password")
		extra := gjson.Get(out, "extra")

		titleLen := len(title.String())
		urlLen := len(url.String()) + 2
		usernameLen := len(username.String()) + 2
		passwordLen := len(password.String()) + 2
		extraLen := len(extra.String()) + 2

		if titleLen < 8 {
			titleLen = 8
		}

		if urlLen < 10 {
			urlLen = 10
		}

		if usernameLen < 10 {
			usernameLen = 15
		}

		if passwordLen < 10 {
			passwordLen = 10
		}

		if extraLen < 10 {
			extraLen = 10
		}

		if title.Exists() && url.Exists() && username.Exists() && password.Exists() && extra.Exists() {
			return model{
				smTable: table.New([]table.Column{
					table.NewColumn("Title", "Title", titleLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("URL", "URL", urlLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Username", "Username", usernameLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Password", "Password", passwordLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Extra", "Extra", extraLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
				}).WithRows([]table.Row{
					table.NewRow(table.RowData{
						"Title": title.String(),
						"URL": url.String(),
						"Username": username.String(),
						"Password": password.String(),
						"Extra": extra.String(),
					}),
				}).Border(constants.TABLE_BORDER_DESIGN),
			}
		}
	} else if o.CreditCards {
		cardName := gjson.Get(out, "card_name")
		cardholderName := gjson.Get(out, "cardholder_name")
		cardType := gjson.Get(out, "type")
		number := gjson.Get(out, "number")
		expiryDate := gjson.Get(out, "expiry_date")
		verificationNumber := gjson.Get(out, "verification_number")

		cardNameLen := len(cardName.String()) + 2
		cardholderNameLen := len(cardholderName.String()) + 2
		cardTypeLen := len(cardType.String()) + 3
		numberLen := len(number.String()) + 2
		expiryDateLen := 13
		verificationNumberLen := 21

		if cardNameLen < 11 {
			cardNameLen = 11
		}

		if cardholderNameLen < 18 {
			cardholderNameLen = 18
		}

		if cardTypeLen < 11 {
			cardTypeLen = 11
		}

		if numberLen < 10 {
			numberLen = 10
		}

		if cardName.Exists() && cardholderName.Exists() && cardType.Exists() && number.Exists() && expiryDate.Exists() && verificationNumber.Exists() {
			return model{
				smTable: table.New([]table.Column{
					table.NewColumn("Card Name", "Card Name", cardNameLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Cardholder Name", "Card Holder Name", cardholderNameLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Card Type", "Card Type", cardTypeLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Number", "Number", numberLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Expiry Date", "Expiry Date", expiryDateLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Verification Number", "Verification Number", verificationNumberLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
				}).WithRows([]table.Row{
					table.NewRow(table.RowData{
						"Card Name": cardName.String(),
						"Cardholder Name": cardholderName.String(),
						"Card Type": cardType.String(),
						"Number": number.String(),
						"Expiry Date": expiryDate.String(),
						"Verification Number": verificationNumber.String(),
					}),
				}).Border(constants.TABLE_BORDER_DESIGN),
			}
		}
	} else if o.Emails {
		email := gjson.Get(out, "email")
		password := gjson.Get(out, "password")

		emailLen := len(email.String()) + 2
		passwordLen := len(password.String()) + 2

		if emailLen < 10 {
			emailLen = 10
		}

		if passwordLen < 10 {
			passwordLen = 10
		}

		if email.Exists() && password.Exists() {
			return model{
				smTable: table.New([]table.Column{
					table.NewColumn("Email", "Email", emailLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Password", "Password", passwordLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
				}).WithRows([]table.Row{
					table.NewRow(table.RowData{
						"Email": email.String(),
						"Password": password.String(),
					}),
				}).Border(constants.TABLE_BORDER_DESIGN),
			}
		}
	} else if o.Notes {
		title := gjson.Get(out, "title")
		note := m.styles.Wrap.Render(gjson.Get(out, "note").String())

		titleLen := len(title.String()) + 2
		noteLen := len(note) + 2

		if titleLen < 10 {
			titleLen = 10
		}

		if noteLen < 10 {
			noteLen = 10
		}

		if title.Exists() && note != "" {
			return model{
				smTable: table.New([]table.Column{
					table.NewColumn("Title", "Title", titleLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Note", "Note", noteLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
				}).WithRows([]table.Row{
					table.NewRow(table.RowData{
						"Title": title.String(),
						"Note": note,
					}),
				}).Border(constants.TABLE_BORDER_DESIGN),
			}
		}
	} else if o.Servers {
		title := gjson.Get(out, "title")
		ip := gjson.Get(out, "ip")
		url := gjson.Get(out, "url")
		username := gjson.Get(out, "username")
		password := gjson.Get(out, "password")
		hosting_username := gjson.Get(out, "hosting_username")
		hosting_password := gjson.Get(out, "hosting_password")
		admin_username := gjson.Get(out, "admin_username")
		admin_password := gjson.Get(out, "admin_password")
		extra := gjson.Get(out, "extra")

		titleLen := len(title.String()) + 2
		ipLen := len(ip.String()) + 2
		urlLen := len(url.String()) + 2
		usernameLen := len(username.String()) + 2
		passwordLen := len(password.String()) + 2
		hosting_usernameLen := len(hosting_username.String()) + 2
		hosting_passwordLen := len(hosting_password.String()) + 2
		admin_usernameLen := len(admin_username.String()) + 2
		admin_passwordLen := len(admin_password.String()) + 2
		extraLen := len(extra.String()) + 2

		if titleLen < 10 {
			titleLen = 10
		}

		if ipLen < 10 {
			ipLen = 10
		}

		if urlLen < 10 {
			urlLen = 10
		}

		if usernameLen < 10 {
			usernameLen = 10
		}

		if passwordLen < 10 {
			passwordLen = 10
		}

		if hosting_usernameLen < 18 {
			hosting_usernameLen = 18
		}

		if hosting_passwordLen < 18 {
			hosting_passwordLen = 18
		}

		if admin_usernameLen < 16 {
			admin_usernameLen = 16
		}

		if admin_passwordLen < 16 {
			admin_passwordLen = 16
		}

		if extraLen < 10 {
			extraLen = 10
		}

		if title.Exists() && url.Exists() && username.Exists() && password.Exists() && hosting_username.Exists() && hosting_password.Exists() && admin_username.Exists() && admin_password.Exists() && extra.Exists() {
			return model{
				smTable: table.New([]table.Column{
					table.NewColumn("Title", "Title", titleLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("IP", "IP", ipLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("URL", "URL", urlLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Username", "Username", usernameLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Password", "Password", passwordLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Hosting Username", "Hosting Username", hosting_usernameLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Hosting Password", "Hosting Password", hosting_passwordLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Admin Username", "Admin Username", admin_usernameLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Admin Password", "Admin Password", admin_passwordLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
					table.NewColumn("Extra", "Extra", extraLen).WithStyle(
						lipgloss.NewStyle().
							Align(lipgloss.Center),
					),
				}).WithRows([]table.Row{
					table.NewRow(table.RowData{
						"Title": title.String(),
						"IP": ip.String(),
						"URL": url.String(),
						"Username": username.String(),
						"Password": password.String(),
						"Hosting Username": hosting_username.String(),
						"Hosting Password": hosting_password.String(),
						"Admin Username": admin_username.String(),
						"Admin Password": admin_password.String(),
						"Extra": extra.String(),
					}),
				}).Border(constants.TABLE_BORDER_DESIGN),
			}
		}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.smTable, _ = m.smTable.Update(msg)

	return m, tea.Quit
}

func (m model) View() string {
	s := "\n\n" + m.smTable.View() + "\n"

	return lipgloss.NewStyle().PaddingLeft(2).SetString(constants.Logo("Secman Reader") + s).String()
}
