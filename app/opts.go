package app

import "github.com/scmn-dev/secman/pkg/options"

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
}

var GenOpts = options.GenOptions{
	Length: 10,
	Raw:    false,
}

var AuthOpts = options.AuthOptions{
	ConfirmLogout: false,
}

var WhoamiOpts = options.WhoamiOptions{
	ShowUser: false,
}
