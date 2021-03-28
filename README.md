# [<img src="https://github.com/secman-team/secman/blob/main/.github/assets/secman.svg" width="300" align="center">][smUrl]

[![RELEASE](https://img.shields.io/github/v/release/secman-team/secman?style=flat)](https://github.com/secman-team/secman/releases/latest)

## Code Status

![CircleCI](https://circleci.com/gh/secman-team/secman.svg?style=svg)
![CodeQL](https://img.shields.io/github/workflow/status/secman-team/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github)
![Go](https://img.shields.io/github/workflow/status/secman-team/secman/Go%20CI?color=blue&label=Go%20Build&logo=go)
![Secman CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20CI?color=blue&label=Secman%20CI&logo=github-actions&logoColor=white)

---

> `secman` is a passowrd manager can store, retrieves, generates, synchronizes passwords and save files securely, and is written in [<img src=".github/assets/go.svg" width="23" align="center">][smUrl]! The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

## Installation ⬇

### Pre-requisites

> secman needs [**go**][goUrl], [**git**](https://git-scm.com), [**ruby**](https://www.ruby-lang.org) and [**gh cli**](https://cli.github.com)

- ![go](https://img.shields.io/static/v1?label=%20&message=v1.11%20and%20above&color=9cf&logo=go)
- ![git](https://img.shields.io/static/v1?label=%20&message=git&color=9cf&logo=git)
- ![ruby](https://img.shields.io/static/v1?label=%20&message=ruby&color=9cf&logo=ruby&logoColor=red)
- ![gh cli](https://img.shields.io/static/v1?label=%20&message=gh%20cli&color=9cf&logo=github&logoColor=black)
- ![windows needs bash](https://img.shields.io/static/v1?label=%20&message=windows%20needs%20bash&color=9cf&logo=gnu-bash&logoColor=black)

### Using Shell (macOS and Linux)

```sh
curl -fsSL https://secman-team.github.io/install/install.sh | bash
```

### Using PowerShell (Windows)

```sh
iwr -useb https://secman-team.github.io/install/install.ps1 | iex
```

### Using [Homebrew](https://brew.sh) (macOS and Linux)

```sh
brew tap secman-team/smx
brew install secman
```

### Using [Scoop](https://scoop.sh) (Windows)

```pwsh
scoop bucket add secman https://github.com/secman-team/secman
scoop install secman
```

### MSI Installer

> MSI installer is available for download on the [releases](https://github.com/secman-team/secman/releases/latest).

## Build from source

see [docs/from_source.md](https://github.com/secman-team/secman/blob/main/docs/from_source.md)

## Getting started with secman

> Initializing Vault

```sh
secman init
```

> Start using `secman`

```sh
secman insert MY_SECRET_TOKEN
```

> Sync your passwords

```sh
secman-sync sync
```

> see [commands](https://github.com/secman-team/secman/blob/main/docs/commands.md)

## License

[secman][smUrl] is licensed under the terms of [MIT][mitUrl] License

[MIT][mitUrl]

[goUrl]: https://goland.org
[smUrl]: https://secman.vercel.app
[mitUrl]: https://github.com/abdfnx/secman/blob/main/LICENSE
