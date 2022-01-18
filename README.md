<div align="center">
  <h1>Secman CLI</h1>
  <p align="center">
   <img src="https://assets.secman.dev/assets/Secman.svg" />
  </p>
  <a href="https://npm.im/secman"><img alt="npm" src="https://img.shields.io/npm/v/secman?label=npm&logo=npm&style=flat-square" /></a>
  <a href="https://github.com/scmn-dev/secman/blob/main/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/scmn-dev/secman?style=flat-square" /></a>
  <br />
  <br />
  <a href="https://secman.dev">Website</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://secman.dev/docs">Docs</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://secman.dev/blog">Blog</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://twitter.com/_secman">Twitter</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://github.com/scmn-dev/secman/tree/v1">Secman V1</a>
  <br />
  <hr />
</div>

## Secman Products

- [**Secman Desktop**](https://d.secman.dev)
- [**Secman Extension**](https://secman.dev/extension)

---

> `secman` is a password manager can store, retrieves, generates, and synchronizes passwords, and is written in _**TypeScript**_! The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. and you can easily manage your passwords from everywhere.

## Features

- **Not GPG cored**.
- **It uses a master password to securely store your passwords**.
- **It syncs your passwords**.
- **Easy to use**.
- **It is written in TypeScript**.
- **You can easily manage your passwords from everywhere, desktop, web, terminal, and more**.

## Examples

> Initialize `~/.secman`

```bash
secman init
```

> Authenticate

```bash
secman auth

# with one command
secman auth --username EMAIL --password MASTER_PASSWORD

# read the master password from stdin
cat password.txt | secman auth -u EMAIL --password-stdin
```

> Create a new password

```bash
secman new -l

✔ Title › Twitter
✔ URL › https://twitter.com
✔ Username › _secman
✔ Password › *********
✔ Extra › no extra

Password created
```

> List passwords

```
secman .
.
├──Logins
│  └──Twitter
├──Credit Cards
├──Emails
├──Notes
└──Servers
```

> Read It

```
secman read -l Twitter

╭─────────┬─────────────────────┬──────────────────┬───────────┬──────────╮
│ Title   │ URL                 │ Username         │ Password  │ Extra    │
├─────────┼─────────────────────┼──────────────────┼───────────┼──────────┤
│ Twitter │ https://twitter.com │ hello@secman.dev │ ••••••••• │ no extra │
╰─────────┴─────────────────────┴──────────────────┴───────────┴──────────╯
```

#### show password

```
secman read -lp Twitter

╭─────────┬─────────────────────┬──────────────────┬───────────┬──────────╮
│ Title   │ URL                 │ Username         │ Password  │ Extra    │
├─────────┼─────────────────────┼──────────────────┼───────────┼──────────┤
│ Twitter │ https://twitter.com │ hello@secman.dev │ hitwitter │ no extra │
╰─────────┴─────────────────────┴──────────────────┴───────────┴──────────╯
```

> Edit password field

```bash
secman edit -l Twitter

? Pick a field › - Use arrow-keys. Return to submit.
❯   Title
    URL
    Username
    Password
    Extra
```

#### edit multiple fields

```code
secman edit -lm Twitter

? Pick a field ›
? Pick a field ›
Instructions:
    ↑/↓: Highlight option
    ←/→/[space]: Toggle selection
    a: Toggle all
    enter/return: Complete answer
◯   Title
◉   URL
◯   Username
◯   Password
◉   Extra
```

## Installation ⬇

### Using npm

```bash
npm i -g secman
```

> (Windows): if you get an error you might need to change the **execution policy** _**(i.e. enable Powershell)**_ via

```powershell
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
```

### Using [Homebrew](https://brew.sh) (macOS and Linux)

```bash
brew tap scmn-dev/secman
brew install secman
```

### Using script (Ubuntu/Debian)

```bash
curl -sL https://cli.secman.dev | bash
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

see [**building from source**](https://secman.dev/docs/contributing/building_from_source) doc.

## Getting started with secman

> Initializing

```bash
secman init
```

> Authenticate

```bash
secman auth
```

> Insert a New Password

```bash
secman insert --[PASSWORD_TYPE]
```

> List Passwords

```bash
secman .
```

> Read The Password

```
secman read --[PASSWORD_TYPE] <PASSWORD_NAME>
```

> Edit Password

```bash
secman edit --[PASSWORD_TYPE] <PASSWORD_NAME>
```

> Generate

```bash
secman generate
```

> Edit Settings

```bash
secman settings
```

> see [commands](https://secman.dev/docs/cli)

## License

[secman][smurl] is licensed under the terms of [MIT][miturl] License

## Some Resources

- [**secman website**](https://secman.dev)
- [**docs**](https://secman.dev/docs)
- [**changelog**](https://secman.dev/changelog)
- [**privacy policy of secman**](https://secman.dev/privacy)

## Contributing

Thanks for your interest in contributing to `secman` . You can start a development environment with [gitpod](https://www.gitpod.io):

[![open in gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scmn-dev/secman)

## Code Status

[![CodeQL](https://img.shields.io/github/workflow/status/scmn-dev/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github&style=flat-square)](https://github.com/scmn-dev/secman/actions/workflows/codeql.yml)
[![Secman CI](https://img.shields.io/github/workflow/status/scmn-dev/secman/Secman%20CI?color=blue&label=Secman%20CI&logo=github-actions&logoColor=white&style=flat-square)](https://github.com/scmn-dev/secman/actions/workflows/secman.yml)
[![Secman Docker CI](https://img.shields.io/github/workflow/status/scmn-dev/secman/Secman%20Docker%20CI?color=blue&label=Secman%20Docker%20CI&logo=docker&style=flat-square)](https://github.com/scmn-dev/secman/actions/workflows/docker.yml)
![Codacy grade](https://img.shields.io/codacy/grade/c434720ddcc84bea982475063f903a81?color=blue&logo=codacy&style=flat-square)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/scmn-dev/secman.svg)](https://starchart.cc/scmn-dev/secman)

[smurl]: https://secman.dev
[miturl]: https://github.com/scmn-dev/secman/blob/main/LICENSE
