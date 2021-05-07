<p align="center">
   <img src="https://imgs-secman.web.app/logo.svg" width="500" />
</p>

<p align="center">
   <img src="https://imgs-secman.web.app/assets/Secman.svg" />
</p>

[![RELEASE](https://img.shields.io/github/v/release/secman-team/secman?style=for-the-badge)](https://github.com/secman-team/secman/releases/latest)

## Code Status

![CodeQL](https://img.shields.io/github/workflow/status/secman-team/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github&style=for-the-badge)
![Go](https://img.shields.io/github/workflow/status/secman-team/secman/Go%20CI?color=blue&label=Go%20Build&logo=go&style=for-the-badge)
![Secman CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20CI?color=blue&label=Secman%20CI&logo=github-actions&logoColor=white&style=for-the-badge)
![Secman Docker CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20Docker%20CI?color=blue&label=Secman%20Docker%20CI&logo=docker&style=for-the-badge)
[![Secman With GitPod](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod&style=for-the-badge)](https://gitpod.io/#https://github.com/secman-team/secman)
![Codacy grade](https://img.shields.io/codacy/grade/8c1ede5d80d2489c9c041e99d67a42f3?color=blue&logo=codacy&style=for-the-badge)

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
- [![git](https://imgs-secman.web.app/badges/git.svg)](https://git-scm.com)
- [![npm](https://imgs-secman.web.app/badges/npm.svg)](https://nodejs.org)

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

see [installing from source](https://secman.vercel.app/docs/secman/build_from_source)

## Getting started with secman

> Initializing Vault

<img src="https://imgs-secman.web.app/assets/Init.svg" />

> Start using `secman`

<img src="https://imgs-secman.web.app/assets/Insert.svg" />

> Authenticate With **Github**

<img src="https://imgs-secman.web.app/assets/Auth-Login.svg" />


> Sync your passwords

<img src="https://imgs-secman.web.app/assets/Sync-Start.svg" />

> see [commands](https://secman.vercel.app/docs/commands)

## License

[secman][smUrl] is licensed under the terms of [MIT][mitUrl] License

## Our Websites

- [**main website**](https://secman.vercel.app)
- [**deps website**](https://secman-team.github.io)
- [**changelog website**](https://secman-chlog.web.app)

[goUrl]: https://goland.org
[smUrl]: https://secman.vercel.app
[mitUrl]: https://github.com/secman-team/secman/blob/main/LICENSE
