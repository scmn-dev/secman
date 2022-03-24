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
	AutoGenerate bool
}

type GenOptions struct {
	Length int
	Raw    bool
}

type AuthOptions struct {
	Username	  string
	Password	  string
	PasswordStdin bool
	ConfirmLogout bool
}

type RootOptions struct {
	Version bool
}

type WhoamiOptions struct {
	ShowUser bool
}

type EncryptOptions struct {
	AES    bool
	SHA256 bool
	SHA512 bool
	MD5    bool
	AESKey string
}