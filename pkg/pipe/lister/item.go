package lister

type item struct {
	title, url, username, password, extra,
	cardHolderName, cType, number, expiryDate, verificationNumber,
	email,
	hiddenNote, note,
	ip, hosting_username, hosting_password, admin_username, admin_password string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string {
	switch currentType {
		case "Login":
			return i.username
		case "Credit Card":
			return i.number
		case "Email":
			return i.email
		case "Server":
			return i.ip
		default:
			return ""
	}
}

func (i item) FilterValue() string { return i.title }

func NewLoginListItem(title, url, username, password, extra string) item {
	return item{title, url, username, password, extra, "", "", "", "", "", "", "", "", "", "", "", "", ""}
}

func NewCCListItem(title, cardHolderName, cType, number, expiryDate, verificationNumber string) item {
	return item{title, "", "", "", "", cardHolderName, cType, number, expiryDate, verificationNumber, "", "", "", "", "", "", "", ""}
}

func NewEmailListItem(title, email, password string) item {
	return item{title, "", "", password, "", "", "", "", "", "", email, "", "", "", "", "", "", ""}
}

func NewNoteListItem(title, hiddenNote, note string) item {
	return item{title, "", "", "", "", "", "", "", "", "", "", hiddenNote, note, "", "", "", "", ""}
}

func NewServerListItem(title, ip, url, username, password, hosting_username, hosting_password, admin_username, admin_password, extra string) item {
	return item{title, url, username, password, extra, "", "", "", "", "", "", "", "", ip, hosting_username, hosting_password, admin_username, admin_password}
}
