# <img src="https://assets.secman.dev/apps/sm-win.svg" width="300px" />

sub-app of secman for windows

![sm-win version](https://img.shields.io/npm/v/@secman/sm-win?color=blue&label=version&logo=npm&style=flat-square)
![npm](https://img.shields.io/npm/dw/@secman/sm-win?style=flat-square)

* [Usage](#usage)
* [Commands](#commands)
* [Usage](#usage)
* [Commands](#commands)

# Install

```bash
npm i -g @secman/sm-win
```

# Usage
```sh-session
$ sm-win COMMAND

running command...
$ sm-win (-v | --version | version)

@secman/sm-win/0.2.4 win32-x64
$ sm-win --help [COMMAND]
USAGE
  $ sm-win COMMAND
```

# Commands
* [`sm-win help [COMMAND]`](#sm-win-help-command)
* [`sm-win fetch`](#sm-win-fetch)
* [`sm-win start`](#sm-win-start)
* [`sm-win uninstall`](#sm-win-uninstall)

## `sm-win help [COMMAND]`

display help for sm-win

```
USAGE
  $ sm-win help [COMMAND]

ARGUMENTS
  COMMAND  command to show help for

OPTIONS
  --all  see all commands in CLI
```

## `sm-win fetch`

Fetch if there's a new release.

```
USAGE
  $ sm-win fetch
```

## `sm-win start`

Start Upgrade secman

```
USAGE
  $ sm-win start
```

## `sm-win uninstall`

Uninstall your secman

```
USAGE
  $ sm-win uninstall

OPTIONS
  -d, --delete-data  delete data (~/.secman)
```
