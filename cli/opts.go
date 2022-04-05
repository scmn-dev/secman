package cli

import "github.com/scmn-dev/secman/v6/pkg/options"

var PwOpts = options.PasswordsOptions{
	Password:     "",
	Logins: 	  false,
	CreditCards:  false,
	Emails: 	  false,
	Notes: 		  false,
	Servers: 	  false,
	ShowHidden:   false,
	ShowJsonView: false,
	ShowTreeView: false,
	AutoGenerate: false,
}

var GenOpts = options.GenOptions{
	Length: 10,
	Raw:    false,
}

var AuthOpts = options.AuthOptions{
	Username:      "",
	Password:      "",
	PasswordStdin: false,
	ConfirmLogout: false,
}

var WhoamiOpts = options.WhoamiOptions{
	ShowUser: false,
}

var EncryptOpts = options.EncryptOptions{
	AES:	false,
	SHA256: false,
	SHA512: false,
	MD5:    false,
	AESKey: "",
}
