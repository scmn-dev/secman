# [<img src="https://imgs-secman.web.app/logo.png" align="center">][smUrl] **Secman**

[![RELEASE](https://img.shields.io/github/v/release/secman-team/secman?style=for-the-badge)](https://github.com/secman-team/secman/releases/latest)

## Code Status

![CodeQL](https://img.shields.io/github/workflow/status/secman-team/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github&style=for-the-badge)
![Go](https://img.shields.io/github/workflow/status/secman-team/secman/Go%20CI?color=blue&label=Go%20Build&logo=go&style=for-the-badge)
![Secman CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20CI?color=blue&label=Secman%20CI&logo=github-actions&logoColor=white&style=for-the-badge)
![Secman Docker CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20Docker%20CI?color=blue&label=Secman%20Docker%20CI&logo=docker&style=for-the-badge)

> `secman` is a passowrd manager can store, retrieves, generates, synchronizes passwords and save files securely, and is written in *go*! The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

## Installation ⬇

### Pre-requisites

> secman needs [**git**](https://git-scm.com) and [**ruby**](https://www.ruby-lang.org)

- [![git](https://imgs-secman.web.app/badges/git.svg)](https://git-scm.com)
- [![ruby](https://imgs-secman.web.app/badges/ruby.svg)](https://www.ruby-lang.org/en/)
- _ruby for windows_: **https://rubyinstaller.org**

### Using Shell (macOS and Linux)

```bash
curl -fsSL https://secman-team.github.io/install.sh | bash
```

### Powershell (Windows)

```powershell
iwr -useb https://secman-team.github.io/install.ps1 | iex
```

### Using [Homebrew](https://brew.sh) (macOS and Linux)

```bash
brew tap secman-team/smx
brew install secman
```

### Using [Scoop](https://scoop.sh) (Windows)

```powershell
scoop bucket add secman https://github.com/secman-team/sm-scoop
scoop install secman
```

## Build from source

see [installing from source](https://secman.vercel.app/docs/getting_started/installation#installing-from-source)

## Getting started with secman

> Initializing Vault

```bash
secman init
```

> Start using `secman`

```bash
secman insert MY_SECRET_TOKEN
```

> Sync your passwords

```bash
secman start-sync
```

> see [commands](https://secman.vercel.app/docs/commands)

## License

[secman][smUrl] is licensed under the terms of [MIT][mitUrl] License

[MIT][mitUrl]

## Our Websites

- [**main website**](https://secman.vercel.app)
- [**deps website**](https://secman-team.github.io)
- [**changelog website**](https://secman-chlog.web.app)

[goUrl]: https://goland.org
[smUrl]: https://secman.vercel.app
[mitUrl]: https://github.com/abdfnx/secman/blob/main/LICENSE
