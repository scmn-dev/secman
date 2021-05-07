sm-upg
======

sub-app of secman for windows

![sm-upg version](https://img.shields.io/npm/v/@secman/sm-upg?color=blue&label=version&logo=npm&style=flat-square)
![npm](https://img.shields.io/npm/dw/@secman/sm-upg?style=flat-square)

* [Usage](#usage)
* [Commands](#commands)
* [Usage](#usage)
* [Commands](#commands)

# Install

```bash
npm i -g @secman/sm-upg
```

# Usage
```sh-session
$ sm-upg COMMAND

running command...
$ sm-upg (-v | --version | version)

@secman/sm-upg/0.1.5 win32-x64
$ sm-upg --help [COMMAND]
USAGE
  $ sm-upg COMMAND
```

# Commands
* [`sm-upg help [COMMAND]`](#sm-upg-help-command)
* [`sm-upg fetch`](#sm-upg-fetch)
* [`sm-upg start`](#sm-upg-start)
* [`sm-upg uninstall`](#sm-upg-uninstall)

## `sm-upg help [COMMAND]`

display help for sm-upg

```
USAGE
  $ sm-upg help [COMMAND]

ARGUMENTS
  COMMAND  command to show help for

OPTIONS
  --all  see all commands in CLI
```

## `sm-upg fetch`

Fetch if there's a new release.

```
USAGE
  $ sm-upg fetch
```

## `sm-upg start`

Start Upgrade secman

```
USAGE
  $ sm-upg start
```

## `sm-upg uninstall`

Uninstall your secman

```
USAGE
  $ sm-upg uninstall

OPTIONS
  -d, --delete-data  delete data (~/.secman)
```
