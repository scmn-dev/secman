# Secman CLI

[**secman cli**](https://secman.dev) is a TUI password manager can store, retrieves, generates, and synchronizes passwords, The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. and you can easily manage your passwords from everywhere with **Secman Cloud** 😉.

### Featuers

- Not GPG cored.
- It uses a master password to securely store your passwords.
- It syncs your passwords.
- Easy to use.
- It is written in Go.
- You can easily manage your passwords from everywhere, desktop, web, terminal, and more.

### Installation ⬇

### Using script

- Shell

```bash
curl -sL https://u.secman.dev | bash
```

- PowerShell

```powershell
iwr -useb https://w.secman.dev | iex
```

**then restart your powershell**

### Homebrew

```
brew install scmn-dev/tap/secman
```

### Scoop

```
scoop bucket add secman https://github.com/scmn-dev/scoop
scoop install secman
```

## Usage

##### to learn about secman types, check out the [secman types](https://secman.dev/docs/password-types) page.

### Initialize `~/.secman`

**Command**

```
secman init
```

**Flags**

no flags

### Manage secman's authentication state.

**Command**

```
secman auth
```

**Other Commands**

```
create     Create a new secman account.
login      Authenticate with secman.
logout     Logout of the current user account.
refresh    Refresh the current user account.
```

**Flags**

no flags

### Insert a New Password

**Command**

```
secman insert --[PASSWORD_TYPE]
```

**Flags**

```
-c, --credit-cards   Insert a credit card to your vault.
-e, --emails         Insert a email to your vault.
-l, --logins         Insert a login password to your vault.
-n, --notes          Insert a note to your vault.
-s, --servers        Insert a server to your vault.
```

### List all your passwords

**Command**

```
secman .
```

**Flags**

no flags

### Read a password

**Command**

```
secman read --[PASSWORD_TYPE] <PASSWORD_NAME>
```

**Flags**

```
-c, --credit-cards   Read password from credit cards type.
-e, --emails         Read password from emails type.
-j, --json           Print password in JSON view.
-l, --logins         Read password from logins type.
-n, --notes          Read password from notes type.
-s, --servers        Read password from servers type.
-p, --show-hidden    Show hidden values.
```

### Update/Edit a password value

**Command**

```
secman edit --[PASSWORD_TYPE] <PASSWORD_NAME>
```

**Flags**

```
-c, --credit-cards   Edit password from credit cards type.
-e, --emails         Edit password from emails type.
-l, --logins         Edit password from logins type.
-n, --notes          Edit password from notes type.
-s, --servers        Edit password from servers type.
```

### Delete a password

**Command**

```
secman delete --[PASSWORD_TYPE] <PASSWORD_NAME>
```

**Flags**

```
-c, --credit-cards   Delete password from credit cards type.
-e, --emails         Delete password from emails type.
-l, --logins         Delete password from logins type.
-n, --notes          Delete password from notes type.
-s, --servers        Delete password from servers type.
```

### Generate a password

**Command**

```
secman generate --length 20
```

**Flags**

```
-l, --length int   Set the length of the password. (default 10)
-r, --raw          Generate a password and print it.
```

> to learn more about secman commands run `secman help`

## Some Resources

- [**secman website**][smweb]
- [**docs**](https://secman.dev/docs)
- [**changelog**](https://secman.dev/changelog)
- [**privacy policy of secman**](https://secman.dev/privacy)

## Other Secman Products

- [**Secman Desktop**](https://github.com/scmn-dev/desktop)
- [**Secman Extension**](https://secman.dev/extension)

## Contributing

Thanks for your interest in contributing to `secman`. You can start a development environment with [gitpod](https://www.gitpod.io):

[![open in gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scmn-dev/secman)

## Special thanks ❤

thanks to [**@charmbracelet**](https://github.com/charmbracelet) for thier awesome TUI libraries 🏗.

## License

[secman][smweb] is licensed under the terms of [MIT][licurl] License

[smweb]: https://secman.dev
[licurl]: https://github.com/scmn-dev/secman/blob/main/LICENSE
