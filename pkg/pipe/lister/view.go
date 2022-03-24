package lister

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"github.com/scmn-dev/secman/constants"
	"github.com/evertras/bubble-table/table"
)

var (
	smTable table.Model
	currentType string
)

func (m model) listView() string {
	if m.state == LOGIN {
		return st.ListView.Render(m.loginsList.View())
	} else if m.state == CC {
		return st.ListView.Render(m.creditCardsList.View())
	} else if m.state == EMAIL {
		return st.ListView.Render(m.emailsList.View())
	} else if m.state == NOTE {
		return st.ListView.Render(m.notesList.View())
	} else if m.state == SERVER {
		return st.ListView.Render(m.serversList.View())
	}

	return st.ListView.Render("")
}

func (m model) detailView() string {
	builder := &strings.Builder{}
	divider := st.Divider.Render(strings.Repeat("-", m.viewport.Width)) + "\n"
	showTable := true
	pwType := ""
	pwTitle := ""

	switch m.state {
		case LOGIN:
			currentType = "Login"
			pwType = "-l"

			if it := m.loginsList.SelectedItem(); it != nil {
				title := it.(item).title
				url := it.(item).url
				username := it.(item).username
				password := it.(item).password
				extra := it.(item).extra

				builder.WriteString(divider)
				builder.WriteString("Title" + "\n\n")
				builder.WriteString(title + "\n")
				builder.WriteString(divider)
				builder.WriteString("URL" + "\n\n")
				builder.WriteString(url + "\n")
				builder.WriteString(divider)
				builder.WriteString("Username" + "\n\n")
				builder.WriteString(username + "\n")
				builder.WriteString(divider)
				builder.WriteString("Password" + "\n\n")
				builder.WriteString(password + "\n")
				builder.WriteString(divider)
				builder.WriteString("Extra" + "\n\n")
				builder.WriteString(extra + "\n")
				builder.WriteString(divider)
		
				titleLen := len(title)
				urlLen := len(url) + 2
				usernameLen := len(username) + 2
				passwordLen := len(password) + 2
				extraLen := len(extra) + 2
		
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
		
				smTable = table.New([]table.Column{
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
						"Title": title,
						"URL": url,
						"Username": username,
						"Password": password,
						"Extra": extra,
					}),
				}).Border(constants.TABLE_BORDER_DESIGN)

				pwTitle = title
			} else {
				builder.WriteString("No item selected")
			}
		case CC:
			currentType = "Credit Card"
			pwType = "-c"

			if it := m.loginsList.SelectedItem(); it != nil {
			cardName := m.creditCardsList.SelectedItem().(item).title
			cardholderName := m.creditCardsList.SelectedItem().(item).cardHolderName
			cardType := m.creditCardsList.SelectedItem().(item).cType
			number := m.creditCardsList.SelectedItem().(item).number
			expiryDate := m.creditCardsList.SelectedItem().(item).expiryDate
			verificationNumber := m.creditCardsList.SelectedItem().(item).verificationNumber

			builder.WriteString(divider)
			builder.WriteString("Card Name" + "\n\n")
			builder.WriteString(cardName + "\n")
			builder.WriteString(divider)
			builder.WriteString("Card Holder" + "\n\n")
			builder.WriteString(cardholderName + "\n")
			builder.WriteString(divider)
			builder.WriteString("Type" + "\n\n")
			builder.WriteString(cardType + "\n")
			builder.WriteString(divider)
			builder.WriteString("Number" + "\n\n")
			builder.WriteString(number + "\n")
			builder.WriteString(divider)
			builder.WriteString("Expiry Date" + "\n\n")
			builder.WriteString(expiryDate + "\n")
			builder.WriteString(divider)
			builder.WriteString("Verification Number" + "\n\n")
			builder.WriteString(verificationNumber + "\n")
			builder.WriteString(divider)

			cardNameLen := len(cardName) + 2
			cardholderNameLen := len(cardholderName) + 2
			cardTypeLen := len(cardType) + 3
			numberLen := len(number) + 2
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

			smTable = table.New([]table.Column{
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
					"Card Name": cardName,
					"Cardholder Name": cardholderName,
					"Card Type": cardType,
					"Number": number,
					"Expiry Date": expiryDate,
					"Verification Number": verificationNumber,
				}),
			}).Border(constants.TABLE_BORDER_DESIGN)
			
			pwTitle = cardName
		}

		case EMAIL:
			currentType = "Email"
			pwType = "-e"

			title := m.emailsList.SelectedItem().(item).title
			email := m.emailsList.SelectedItem().(item).email
			password := m.emailsList.SelectedItem().(item).password

			builder.WriteString(divider)
			builder.WriteString("Title" + "\n\n")
			builder.WriteString(title + "\n")
			builder.WriteString(divider)
			builder.WriteString("Email" + "\n\n")
			builder.WriteString(email + "\n")
			builder.WriteString(divider)
			builder.WriteString("Password" + "\n\n")
			builder.WriteString(password + "\n")
			builder.WriteString(divider)

			titleLen := len(title) + 2
			emailLen := len(email) + 2
			passwordLen := len(password) + 2

			if titleLen < 8 {
				titleLen = 8
			}

			if emailLen < 10 {
				emailLen = 10
			}

			if passwordLen < 10 {
				passwordLen = 10
			}

			smTable = table.New([]table.Column{
				table.NewColumn("Title", "Title", titleLen).WithStyle(
					lipgloss.NewStyle().
						Align(lipgloss.Center),
				),
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
					"Title": title,
					"Email": email,
					"Password": password,
				}),
			}).Border(constants.TABLE_BORDER_DESIGN)

			pwTitle = title

		case NOTE:
			pwType = "-n"

			title := m.notesList.SelectedItem().(item).title
			note := m.notesList.SelectedItem().(item).note

			builder.WriteString(divider)
			builder.WriteString("Title" + "\n\n")
			builder.WriteString(title + "\n")
			builder.WriteString(divider)
			builder.WriteString("Note" + "\n\n")
			builder.WriteString(m.styles.Wrap.Render(note) + "\n")
			builder.WriteString(divider)

			showTable = false
			pwTitle = title

		case SERVER:
			pwType = "-s"

			currentType = "Server"
			title := m.serversList.SelectedItem().(item).title
			ip := m.serversList.SelectedItem().(item).ip
			url := m.serversList.SelectedItem().(item).url
			username := m.serversList.SelectedItem().(item).username
			password := m.serversList.SelectedItem().(item).password
			hosting_username := m.serversList.SelectedItem().(item).hosting_username
			hosting_password := m.serversList.SelectedItem().(item).hosting_password
			admin_username := m.serversList.SelectedItem().(item).admin_username
			admin_password := m.serversList.SelectedItem().(item).admin_password
			extra := m.serversList.SelectedItem().(item).extra

			builder.WriteString(divider)
			builder.WriteString("Title" + "\n\n")
			builder.WriteString(title + "\n")
			builder.WriteString(divider)
			builder.WriteString("IP" + "\n\n")
			builder.WriteString(ip + "\n")
			builder.WriteString(divider)
			builder.WriteString("URL" + "\n\n")
			builder.WriteString(url + "\n")
			builder.WriteString(divider)
			builder.WriteString("Username" + "\n\n")
			builder.WriteString(username + "\n")
			builder.WriteString(divider)
			builder.WriteString("Password" + "\n\n")
			builder.WriteString(password + "\n")
			builder.WriteString(divider)
			builder.WriteString("Hosting Username" + "\n\n")
			builder.WriteString(hosting_username + "\n")
			builder.WriteString(divider)
			builder.WriteString("Hosting Password" + "\n\n")
			builder.WriteString(hosting_password + "\n")
			builder.WriteString(divider)
			builder.WriteString("Admin Username" + "\n\n")
			builder.WriteString(admin_username + "\n")
			builder.WriteString(divider)
			builder.WriteString("Admin Password" + "\n\n")
			builder.WriteString(admin_password + "\n")
			builder.WriteString(divider)
			builder.WriteString("Extra" + "\n\n")
			builder.WriteString(extra + "\n")
			builder.WriteString(divider)

			titleLen := len(title) + 2
			ipLen := len(ip) + 2
			urlLen := len(url) + 2
			usernameLen := len(username) + 2
			passwordLen := len(password) + 2
			hosting_usernameLen := len(hosting_username) + 2
			hosting_passwordLen := len(hosting_password) + 2
			admin_usernameLen := len(admin_username) + 2
			admin_passwordLen := len(admin_password) + 2
			extraLen := len(extra) + 2

			if titleLen < 8 {
				titleLen = 8
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

			if hosting_usernameLen < 10 {
				hosting_usernameLen = 10
			}

			if hosting_passwordLen < 10 {
				hosting_passwordLen = 10
			}

			if admin_usernameLen < 10 {
				admin_usernameLen = 10
			}

			if admin_passwordLen < 10 {
				admin_passwordLen = 10
			}

			if extraLen < 10 {
				extraLen = 10
			}

			smTable = table.New([]table.Column{
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
					"Title": title,
					"IP": ip,
					"URL": url,
					"Username": username,
					"Password": password,
					"Hosting Username": hosting_username,
					"Hosting Password": hosting_password,
					"Admin Username": admin_username,
					"Admin Password": admin_password,
					"Extra": extra,
				}),
			}).Border(constants.TABLE_BORDER_DESIGN)

			pwTitle = title
	}

	if showTable {
		builder.WriteString(smTable.View())
	}

	if strings.Contains(pwTitle, " ") {
		pwTitle = fmt.Sprintf("\"%s\"", pwTitle)
	}

	builder.WriteString("\n\nto read more about this password, run `secman read " + pwType + " " + pwTitle + "`\n")

	return wordwrap.String(builder.String(), m.viewport.Width)
}

func (m model) View() string {
	m.viewport.SetContent(m.detailView())

	return lipgloss.JoinHorizontal(lipgloss.Bottom, m.listView(), m.viewport.View())
}
