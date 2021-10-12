<p align="center">
   <img src="https://assets.secman.dev/logo.svg" width="500" />
</p>

<p align="center">
   <img src="https://assets.secman.dev/assets/Secman.svg" />
</p>

[![RELEASE](https://img.shields.io/github/v/release/scmn-dev/secman?style=flat-square)](https://github.com/scmn-dev/secman/releases/latest)

## Secman Products

- [**Secman Desktop**](https://d.secman.dev)
- [**Secman Extension**](https://secman.dev/extension)

## Code Status

[![CodeQL](https://img.shields.io/github/workflow/status/scmn-dev/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github&style=flat-square)](https://github.com/scmn-dev/secman/actions/workflows/codeql.yml)
[![Secman CI](https://img.shields.io/github/workflow/status/scmn-dev/secman/Secman%20CI?color=blue&label=Secman%20CI&logo=github-actions&logoColor=white&style=flat-square)](https://github.com/scmn-dev/secman/actions/workflows/secman.yml)
[![Secman Docker CI](https://img.shields.io/github/workflow/status/scmn-dev/secman/Secman%20Docker%20CI?color=blue&label=Secman%20Docker%20CI&logo=docker&style=flat-square)](https://github.com/scmn-dev/secman/actions/workflows/docker.yml)
[![Secman With GitPod](https://img.shields.io/badge/Gitpod-Ready%20to%20Code-blue?logo=gitpod&style=flat-square)](https://gitpod.io/#https://github.com/scmn-dev/secman)
![Codacy grade](https://img.shields.io/codacy/grade/c434720ddcc84bea982475063f903a81?color=blue&logo=codacy&style=flat-square)

> Open In [**VSCode**](https://code.visualstudio.com)

[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/scmn-dev/secman)

> `secman` is a password manager can store, retrieves, generates, synchronizes passwords and save files securely, and is written in _**go**_! The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

## Featurs

- **Not GPG cored**.
- **It uses a master password to securely store your passwords**.
- **It syncs your passwords**.
- **Supports encrypting arbitrary files**.

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
curl -fsSL https://unix.secman.dev | bash
```

### Powershell (Windows)

```powershell
iwr -useb https://win.secman.dev | iex
```

> if you get an error you might need to change the **execution policy** _**(i.e. enable Powershell)**_ via

```powershell
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
```

### Using [Homebrew](https://brew.sh) (macOS and Linux)

```bash
brew tap scmn-dev/secman
brew install secman
```

### Using [Scoop](https://scoop.sh) (Windows)

```powershell
scoop bucket add secman https://github.com/scmn-dev/secman
scoop install secman
```

### Via [Docker](https://docker.com)

> `secman cli` image

```bash
docker run -it smcr/secman-cli
```

see [**secman cli docs**](https://docker.secman.dev/docs/sm-cli)

> `secman container` image

```bash
docker run -it smcr/secman
```

see [**secman container docs**](https://docker.secman.dev/docs/sm-container)

## Build from source

see [**building from source**](https://docs.secman.dev/contributing/build_from_source) doc.

## Getting started with secman

> Initializing Vault

<img src="https://assets.secman.dev/assets/Init.svg" />

> Start using `secman`

<img src="https://assets.secman.dev/assets/Insert.svg" />

> Authenticate With **Github**

<img src="https://assets.secman.dev/assets/Auth-Login.svg" />

> Sync your passwords

<img src="https://assets.secman.dev/assets/Sync-Start.svg" />

> see [commands](https://docs.secman.dev/guides)

## License

[secman][smUrl] is licensed under the terms of [MIT][mitUrl] License

## Some Resources

- [**secman website**](https://secman.dev)
- [**docs**](https://secman.dev/docs)
- [**changelog**](https://secman.dev/changelog)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/scmn-dev/secman.svg)](https://starchart.cc/scmn-dev/secman)

[goUrl]: https://goland.org
[smUrl]: https://secman.dev
[mitUrl]: https://github.com/scmn-dev/secman/blob/main/LICENSE
