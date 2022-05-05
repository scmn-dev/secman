## Archived

since 30 april 2022, secman has archived after two years of working, the reason is that there is not enough interest in this project

---

<p align="center">
  <a href="https://secman.dev/#gh-light-mode-only" target="_blank">
    <img src="https://assets.secman.dev/logo.svg" alt="Secman" width="300">
  </a>
  <a href="https://secman.dev/#gh-dark-mode-only" target="_blank">
    <img src="https://assets.secman.dev/logo_w.svg" alt="Secman" width="300">
  </a>
</p>

[**secman**][smweb] is a password manager can store, retrieves, generates, and synchronizes passwords, The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. and you can easily manage your passwords from everywhere with **Secman Cloud** 😉.

<p align="center">
  <img src="https://raw.githubusercontent.com/scmn-dev/.github/main/assets/secman.svg" />
</p>

## Projects

| Project     | Package                                                                   | Version                                                                                                                                                                          | Links                                 |
| ----------- | ------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------- |
| **CLI**     | [`scmn-dev/secman`](https://pkg.go.dev/github.com/scmn-dev/secman/v6)     | ![version](https://img.shields.io/github/v/release/scmn-dev/secman?label=go%40latest&logo=go&style=flat-square)                                                                  | [`README.md`](README.md)              |
| **Hub** | [`@secman/hub`](https://github.com/scmn-dev/secman/tree/main/hub) | ![version](https://img.shields.io/github/package-json/v/scmn-dev/secman?filename=hub%2Fpackage.json&label=npm%40latest&logo=npm&style=flat-square)                           | [`README.md`](hub/README.md)      |
| **Scc**    | [`@secman/scc`](https://github.com/scmn-dev/secman/tree/main/scc)       | ![version)](https://img.shields.io/github/package-json/v/scmn-dev/secman?color=blue&filename=scc%2Fpackage.json&label=npm%40latest&logo=npm&logoColor=blue&style=flat-square)   | [`README.md`](scc/README.md)         |
| **Crypto**  | [`@secman/crypto`](https://github.com/scmn-dev/secman/tree/main/crypto)   | ![version)](https://img.shields.io/github/package-json/v/scmn-dev/secman?color=blue&filename=crypto%2Fpackage.json&label=npm%40latest&logo=npm&logoColor=blue&style=flat-square) | [`README.md`](crypto/README.md) |

## Featuers

- **Not GPG cored.**
- **It uses a master password to securely store your passwords.**
- **It syncs your passwords.**
- **Easy to use.**
- **It is written in Go to be fast.**
- **You can easily manage your passwords from everywhere, from desktop, web, command line, and more.**

#### to learn about secman types, check out the [secman types](https://secman.dev/docs/password-types) page.

## Installation ⬇

> ### Secman CLI

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

> ### Secman Hub

Refer to [secman hub](https://secman.dev/docs/hub) instructions.

> ### Secman Web Extension

Check out [secman web extension](https://secman.dev/docs/extension) for installation instructions.

## Usage

## Some Resources

- [**website**][smweb]
- [**docs**](https://secman.dev/docs)
- [**blog**](https://secman.dev/blog)
- [**privacy policy of secman**](https://secman.dev/privacy)

## Contributing

Thanks for your interest in contributing to `secman`. You can start a development environment with [gitpod](https://www.gitpod.io)

[![open in gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scmn-dev/secman)

## Special thanks ❤

thanks to [**@charmbracelet**](https://github.com/charmbracelet) for thier awesome TUI libraries 🏗.

## License

[**secman**][smweb] is licensed under the terms of [MIT](https://github.com/scmn-dev/secman/blob/main/LICENSE) License

[smweb]: https://secman.dev
