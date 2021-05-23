<p align="center">
   <img src="https://assets.secman.dev/logo.svg" width="500" />
</p>

<p align="center">
   <img src="https://assets.secman.dev/assets/Secman.svg" />
</p>

[![RELEASE](https://img.shields.io/github/v/release/secman-team/secman?style=for-the-badge)](https://github.com/secman-team/secman/releases/latest)

## Code Status

[![CodeQL](https://img.shields.io/github/workflow/status/secman-team/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github&style=for-the-badge)](https://github.com/secman-team/secman/actions/workflows/codeql.yml)
[![Go](https://img.shields.io/github/workflow/status/secman-team/secman/Go?color=blue&label=Go%20Build&logo=go&style=for-the-badge)](https://github.com/secman-team/secman/actions/workflows/go.yml)
[![Secman CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20CI?color=blue&label=Secman%20CI&logo=github-actions&logoColor=white&style=for-the-badge)](https://github.com/secman-team/secman/actions/workflows/secman.yml)
[![Secman Docker CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20Docker%20CI?color=blue&label=Secman%20Docker%20CI&logo=docker&style=for-the-badge)](https://github.com/secman-team/secman/actions/workflows/docker.yml)
[![Secman With GitPod](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod&style=for-the-badge)](https://gitpod.io/#https://github.com/secman-team/secman)
![Codacy grade](https://img.shields.io/codacy/grade/d222c27c970f4dc086b77e83809bffde?color=blue&logo=codacy&style=for-the-badge)

> `secman` is a password manager can store, retrieves, generates, synchronizes passwords and save files securely, and is written in *go*! The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

## Featurs

- Not GPG cored.
- It uses a master password to securely store your passwords.
- It syncs your passwords.
- Supports encrypting arbitrary files.

## Forms

1. Password Form

```x
└──PASSWORDNAME
```

```x
secman insert PASSWORDNAME
```

2. Folder Form

```x
└──FOLDERNAME
   └──PASSWORDNAME
```

```x
secman insert FOLDERNAME/PASSWORDNAME
```

## Installation ⬇

### Pre-requisites

> secman needs [**git**](https://git-scm.com) (and just [**npm**](https://nodejs.org) for windows)
- [![git](https://assets.secman.dev/badges/git.svg)](https://git-scm.com)
- [![npm](https://assets.secman.dev/badges/npm.svg)](https://nodejs.org)

### Using Shell (macOS and Linux)

```bash
curl -fsSL https://cli.secman.dev/install.sh | bash
```

### Powershell (Windows)

```powershell
iwr -useb https://cli.secman.dev/install.ps1 | iex
```

> if you get an error you might need to change the **execution policy** _**(i.e. enable Powershell)**_ via

```powershell
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
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

see [installing from source](https://secman.dev/docs/secman/build_from_source)

## Getting started with secman

> Initializing Vault

<img src="https://assets.secman.dev/assets/Init.svg" />

> Start using `secman`

<img src="https://assets.secman.dev/assets/Insert.svg" />

> Authenticate With **Github**

<img src="https://assets.secman.dev/assets/Auth-Login.svg" />

> Sync your passwords

<img src="https://assets.secman.dev/assets/Sync-Start.svg" />

> see [commands](https://secman.dev/docs/commands)

## License

[secman][smUrl] is licensed under the terms of [MIT][mitUrl] License

## Our Websites

- [**main website**](https://secman.dev)
- [**secman cli website**](https://cli.secman.dev)
- [**changelog website**](https://changelog.secman.dev)

[goUrl]: https://goland.org
[smUrl]: https://secman.dev
[mitUrl]: https://github.com/secman-team/secman/blob/main/LICENSE
