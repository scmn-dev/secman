package shared

import "github.com/scmn-dev/secman/pkg/options"

func PasswordType(o *options.PasswordsOptions) string {
	if o.Logins {
		return "-l"
	} else if o.CreditCards {
		return "-c"
	} else if o.Emails {
		return "-e"
	} else if o.Notes {
		return "-n"
	} else if o.Servers {
		return "-s"
	}

	return ""
}
