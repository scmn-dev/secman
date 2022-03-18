package options

type PasswordsOptions struct {
	Password     string
	Logins       bool
	CreditCards  bool
	Emails       bool
	Notes	     bool
	Servers	     bool
	ShowHidden   bool
	ShowJsonView bool
	ShowTreeView bool
}

type GenOptions struct {
	Length int
	Raw    bool
}

type AuthOptions struct {
	ConfirmLogout bool
}

type RootOptions struct {
	Version bool
}

type WhoamiOptions struct {
	ShowUser bool
}
