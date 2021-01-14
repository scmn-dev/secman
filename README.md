# [<img src=".github/assets/secman.svg" align="center">][smUrl]

[!["GitHub Discussions"](https://img.shields.io/badge/%20GitHub-%20Discussions-gray.svg?longCache=true&logo=github&colorB=purple)](https://github.com/abdfnx/secman/discussions)
[![MIT LICENSE](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/abdfnx/secman/blob/main/LICENSE)

> stores, retrieves, generates, and synchronizes passwords and files securely and is written in [<img src=".github/assets/go.svg" align="center" width="30">][goUrl] 💪! The most important difference is secman is not GPG based. Instead it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

`secman is meant to be secure enough that you can publicly post your vault.`

## Installation ⬇

## secman with [docker][dkUrl]  (_Recommended_)

> you can create secman virtual machine by [docker][dkUrl]

```sh
❯ docker pull abdcodedoc/secman:latest
❯ docker run -t -i --privileged abdcodedoc/secman
```

## without docker

`secman` requires [Go][goUrl] version 1.11 or later.

```sh
# wsl/linux
❯ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/HEAD/tools/install_linux.sh)"

# macOS
❯ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/HEAD/tools/install_osx.sh)"
```

## Getting started with secman

Create a vault and specify the directory to store passwords in. You will be prompted for your master password:

```sh
❯ secman init
Please enter a strong master password:
2020/12/23 09:54:31 Created directory to store passwords: ~/.secman
```

Finally, to learn more you can either read about the commands listed in this README or run:

```sh
❯ secman help
```

The `--help` argument can be used on any subcommand to describe it and see documentation or examples 😉.

## Configuring secman

The `SECDIR` environment variable specifies the directory that your vault is in.

it's store the vault in the default location `~/.secman`. All subcommands will respect this environment variable, including `init`

## COMMANDS

### Listing Passwords

```code
❯ secman
├──ionic
|  └──pass
└──dev
   └──dev.to
```

This basic command is used to print out the contents of your password vault. It doesn't require you to enter your master password.

### Initializing Vault

```sh
❯ secman init
```

Init should only be run one time, before running any other command. It is used for generating your master public private keypair.

By default, secman will create your password vault in the `.secman` directory within your home directory. You can override this location using the `SECDIR` environment variable.

### Inserting a password

```code
❯ secman insert accounts/ionic
Enter password for accounts/ionic: 
```

Inserting a password in to your vault is easy. If you wish to group multiple entries together, it can be accomplished by prepending a group name followed by a slash to the pass-name.

Here we are adding ionic to the password store within the money group.

### Inserting a file 📝

```sh
❯ secman insert money/budget.csv budget.csv
```

Adding a file works almost the same as insert. Instead it has an extra argument. The file that you want to add to your vault is the final argument.

### Retrieving a password

```code
❯ secman show accounts/ionic
Enter master password:
ionic_is_😎_js_platform
```

Show is used to display a password in standard out.

### Rename a password

```code
❯ secman rename accounts/ionic-hub
Enter new site name for accounts/ionic-hub: accounts/ionic
```

If a password is added with the wrong name it can be updated later. Here we rename ionic site after misspelling the group name.

### Updating/Editing a password

```code
❯ secman edit dev/dev.to
Enter new password for dev/dev.to:
```

If you want to securely update a password for an already existing site, the edit command is helpful.

### Generating a password

```code
❯ secman generate
%L4^!s,Rry!}s:U<QwliL{vQKow321-!tr}:232

❯ secman generate 8
#%Xy1t7E
```

secman can also create randomly generated passwords. The default length of secman generated passwords is 24 characters. This length can be changed by passing an optional length to the generate subcommand.

### Searching the vault

```code
❯ secman find git
└──git
   └──github.com

❯ secman ls dev
└──dev
   └──dev.to
```

`find` and `ls` can both be used to search for all sites that contain a particular substring. It's good for printing out groups of sites as well. `secman ls` is an alias of `secman find`.

### Deleting a vault entry

```code
❯ secman
├──bb
|  └──ff
├──something
|  └──somethingelse.com
└──code.com
   └──dex.io

❯ secman remove bb/ff

❯ secman
├──something
|  └──somethingelse.com
└──code.com
   └──dex.io
```

remove is used for removing sites from the password vault. `secman rm` is an alias of `secman remove`.

### Getting Help

```code
❯ secman --help
```

All subcommands support the `--help` flag.

## `secman-sync`

#### auth

you should authenticate by [`gh cli`](https://cli.github.com) to use **sync** feature

```sh
❯ gh auth login
```

#### sync

```sh
❯ secman-sync sync
```

if you sync your passwords for first time, `create` command will create a private github repo and store the passwords on it

`secman-sync sy` is an alias of `secman-sync sync`

#### clone

```sh
❯ secman-sync clone
```

if you lose your passwords, or you use more than device, you can clone your private repo

`secman-sync cn` is an alias of `secman-sync clone`

#### push

```sh
❯ secman-sync push
```

if there's a new password/s, it's well push it to the repo, like git

`secman-sync ph` is an alias of `secman-sync push`

#### pull

```sh
❯ secman-sync pull
```

we know what `pull` do

alias: `secman-sync pl`

#### getting help

```sh
❯ secman-sync --help | -h
```

## CRYPTOGRAPHY DETAILS

### Generating Passwords

Password generation takes place in the pc package by using the GeneratePassword function. GeneratePassword creates a random password by reading a large amount of randomness using the `func Read([]byte) (int, error)` function in the `crypto/rand` package.

The block of randomness is then read byte-by-byte. Printable characters that match the desired password specification (uppercase, lowercase, symbols, and digits) are then included in the generated password.

### Adding A Site

When a site is added to the password store, a new public private key pair is generated. The newly generated private key, the user's master public key, and a securely generated nonce are used to encrypt the sites data.

The encryption and key computation are done using the `golang.org/x/crypto/nacl/box` package which uses Curve25519, XSalsa20, and Poly1305 to encrypt and authenticate the site's data.

After the site information is added, the site's generated private key is thrown away.

## Update/Uninstall [secman][smUrl]

if you want yo update/uninstall `secman`, you should type

### Update

> `update by secman upd`

```sh
❯ secman upd
```

### Uninstall

> `uninstall by secman-un`

```sh
❯ secman-un
```

## License

[secman][smUrl] is licensed under the terms of [MIT](https://github.com/abdfnx/secman/blob/main/LICENSE) License

```
MIT License

Copyright (c) 2020 abdfnx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

[goUrl]: https://goland.org
[smUrl]: https://secman.web.app
[dkUrl]: https://docker.com
