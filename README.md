# [<img src=".github/assets/secman.svg" width="300" align="center">][smUrl]

[![RELEASE](https://img.shields.io/github/v/release/secman-team/secman?style=flat)](https://github.com/secman-team/secman/releases/latest)

## Code Status

![CodeQL](https://img.shields.io/github/workflow/status/secman-team/secman/CodeQL?color=blue&label=CodeQL%20Build&logo=github)
![Docker](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20Docker%20Image%20CI%20(VM)?color=blue&label=Docker%20Image%20Build&logo=docker)
![Go](https://img.shields.io/github/workflow/status/secman-team/secman/Go%20CI?color=blue&label=Go%20Build&logo=go)
![Secman CI](https://img.shields.io/github/workflow/status/secman-team/secman/Secman%20CI?color=blue&label=Secman%20CI)
![CircleCI](https://img.shields.io/circleci/build/gh/secman-team/secman?color=blue&label=CircleCI&logo=circleci)

---

> `secman` is a passowrd manager can store, retrieves, generates, synchronizes passwords and save files securely, and is written in [<img src=".github/assets/go.svg" width="23" align="center">][smUrl]! The most important difference is secman is not GPG cored. Instead, it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

## Installation ⬇

### Pre-requisites

> secman needs [**go**][goUrl], [**git**](https://git-scm.com), [**ruby**](https://www.ruby-lang.org) and [**gh cli**](https://cli.github.com)

- ![go](https://img.shields.io/static/v1?label=%20&message=v1.11%20and%20above&color=blue&logo=go)
- ![git](https://img.shields.io/static/v1?label=%20&message=%20&color=blue&logo=git)
- ![ruby](https://img.shields.io/static/v1?label=%20&message=%20&color=blue&logo=ruby&logoColor=red)
- ![gh cli](https://img.shields.io/static/v1?label=%20&message=%20&color=blue&logo=github)
- ![windows](https://img.shields.io/static/v1?label=%20&message=%20&color=blue&logo=windows) ![bash](https://img.shields.io/static/v1?label=bash&message=any%20version&color=white&logo=gnu-bash&logoColor=white)

<!-- ```sh
# wsl/linux
curl -fsSL https://raw.githubusercontent.com/secman-team/install/HEAD/install_linux.sh | bash

# macOS
curl -fsSL https://raw.githubusercontent.com/secman-team/install/HEAD/install_osx.sh | bash
``` -->

## Linux

- by [**curl**](#curl) (Recommended)
- by [**linux package managers**](#linux-package-managers)
   - [ubuntu/debian](#ubuntudebian)
   - [fedora/centOS/red hat enterprise linux](#fedoracentosred-hat-enterprise-linux)
- by [**brew**](#brew)

### curl

```sh
curl -fsSL https://raw.githubusercontent.com/secman-team/install/HEAD/install_linux.sh | bash
```

### linux package managers

> secman is avaliable for linux package managers

### _Ubuntu/Debian_

download `secman_latest_version.deb` from https://github.com/secman-team/secman/releases/latest

and install secman by `apt`

```sh
sudo apt install ./secman_latest_version_x64.deb
```

### _Fedora/CentOS/Red Hat Enterprise Linux_

install `secman_latest_version.rpm` from https://github.com/secman-team/secman/releases/latest

and install secman by `rpm`

```sh
sudo rpm -ivh ./secman_latest_version_x64.rpm --nodeps
```

for more see [docs/linux.md](https://github.com/secman-team/secman/blob/main/docs/linux.md)

### brew

```sh
brew tap secman-team/sm-tap
brew install secman
```

---

## MacOS

- by [**curl**](#curl-1) (Recommended)
- by [**brew**](#brew-1)

### curl

```sh
curl -fsSL https://raw.githubusercontent.com/secman-team/install/HEAD/install_osx.sh | bash
```

### brew

```sh
brew tap secman-team/sm-tap
brew install secman
```

---

## Windows

`secman` is available via [scoop](https://scoop.sh), and as downloadable MSI.

### scoop

```pwsh
scoop install secman
```

Note: if you use `git bash`, you can write this command

```sh
curl -fsSL https://raw.githubusercontent.com/secman-team/install/HEAD/install_win.sh | bash
```

### MSI

> MSI installers are available for download on the [releases](https://github.com/secman-team/secman/releases/latest).

## Build from source

see [docs/from_source.md](https://github.com/secman-team/secman/blob/main/docs/from_source.md)

## Getting started with secman

Create a vault and specify the directory to store passwords in. You will be prompted for your master password:

```sh
secman init
Please enter a strong master password:
2020/12/23 09:54:31 Created directory to store passwords: ~/.secman
```

Finally, to learn more you can either read about the commands listed in this README or run:

```code
secman help
```

The `--help` argument can be used on any subcommand to describe it and see documentation or examples 😉.

## Configuring secman with _*`.secman`*_

The `SECDIR` environment variable specifies the directory that your vault is in.

it's store the vault in the default location `~/.secman`. All subcommands will respect this environment variable, including `init`

## COMMANDS

### Listing Passwords

```code
secman
├──ionic
|  └──pass
└──dev
   └──dev.to
```

This basic command is used to print out the contents of your password vault. It doesn't require you to enter your master password.

### Initializing Vault

```sh
secman init
```

Init should only be run one time, before running any other command. It is used for generating your master public private keypair.

By default, secman will create your password vault in the `.secman` directory within your home directory. You can override this location using the `SECDIR` environment variable.

### Inserting a password

```code
secman insert accounts/ionic
Enter password for accounts/ionic: 
```

Inserting a password in to your vault is easy. If you wish to group multiple entries together, it can be accomplished by prepending a group name followed by a slash to the pass-name.

Here we are adding ionic to the password store within the accounts group.

### Inserting a file 📝

```sh
secman insert money/budget.csv budget.csv
```

Adding a file works almost the same as insert. Instead it has an extra argument. The file that you want to add to your vault is the final argument.

### Retrieving a password

```code
secman show accounts/ionic
Enter master password:
ionic_is_😎_js_platform
```

Show is used to display a password in standard out.

### Rename a password

```code
secman rename accounts/ionic-hub
Enter new site name for accounts/ionic-hub: accounts/ionic
```

If a password is added with the wrong name it can be updated later. Here we rename ionic site after misspelling the group name.

### Updating/Editing a password

```code
secman edit dev/dev.to
Enter new password for dev/dev.to:
```

If you want to securely update a password for an already existing site, the edit command is helpful.

### Generating a password

```code
secman gen
%L4^!s,Rry!}s:U<QwliL{vQKow321-!tr}:232

secman gen 8
#%Xy1t7E
```

secman can also create randomly generated passwords. The default length of secman generated passwords is 24 characters. This length can be changed by passing an optional length to the generate subcommand.

### Searching the vault

```code
secman find git
└──git
   └──github.com

secman ls dev
└──dev
   └──dev.to
```

`find` and `ls` can both be used to search for all sites that contain a particular substring. It's good for printing out groups of sites as well. `secman ls` is an alias of `secman find`.

### Deleting a vault entry

```code
secman
├──bb
|  └──ff
├──something
|  └──somethingelse.com
└──code.com
   └──dex.io

secman remove bb/ff

secman
├──something
|  └──somethingelse.com
└──code.com
   └──dex.io
```

remove is used for removing sites from the password vault. `secman rm` is an alias of `secman remove`.

### Getting Help

```code
secman --help
```

All subcommands support the `--help` flag.

## `secman-sync`

### auth

you should authenticate by [`gh cli`](https://cli.github.com) to use **sync** feature

```sh
gh auth login
```

#### sync

```sh
secman-sync sync
```

if you sync your passwords for first time, `sync` command will create a private github repo and store the passwords on it

`secman-sync sy` is an alias of `secman-sync sync`

#### clone

```sh
secman-sync clone
```

if you lose your passwords, or you use more than device, you can clone your private repo

`secman-sync cn` is an alias of `secman-sync clone`

#### push

```sh
secman-sync push
```

if there's a new password/s, it's well push it to the repo, like git

`secman-sync ph` is an alias of `secman-sync push`

#### pull

```sh
secman-sync pull
```

we know what `pull` do

alias: `secman-sync pl`

#### getting help

```code
secman-sync --help | -h
```

## CRYPTOGRAPHY DETAILS

- see [docs/cryptography](https://github.com/secman-team/secman/blob/main/docs/cryptography.md)

## Update/Uninstall [secman][smUrl]

if you want yo update/uninstall `secman`, you should type

### Update

Note: `secman upd` & `secman-un` are only supported in **linux/mac**, but if you use `git bash` in windows, you can use **`upd`** command

#### Linux/Mac/Git Bash

> `update by secman upd`

```sh
secman upd
```

#### in Windows

```pwsh
scoop upgrade secman
```

### Uninstall

> `uninstall by secman-un`

```sh
secman-un
```

#### For Windows

> you can uninstall secman from your **control panel**

**or you can uninstall it by `scoop`**

```sh
scoop uninstall secman
```

## License

[secman][smUrl] is licensed under the terms of [MIT](https://github.com/abdfnx/secman/blob/main/LICENSE) License

[MIT][mitUrl]

[goUrl]: https://goland.org
[smUrl]: https://secman.web.app
[dkUrl]: https://docker.com
[mitUrl]: https://github.com/abdfnx/secman/blob/main/LICENSE
