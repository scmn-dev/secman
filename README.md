# [<img src=".github/assets/secman.svg" align="center">]()

> stores, retrieves, generates, and synchronizes passwords and files securely and is written in [<img src=".github/assets/go.svg" align="center" width="30">][goUrl] 💪! The most important difference is secman is not GPG based. Instead it uses a master password to securely store your passwords. It also supports encrypting arbitrary files.

`secman is meant to be secure enough that you can publicly post your vault.`

## Installation

`secman` requires [Go][goUrl] version 1.11 or later.

### Install by [go][goUrl]

```sh
(cd; GO111MODULE=on go install github.com/abdfnx/secman/v2)
```

### Install by `install.sh`

```sh
❯ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/main/install.sh)"
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

I store my vault in the default location `~/.secman`. All subcommands will respect this environment variable, including `init`

## COMMANDS

### Listing Passwords

```code
❯ secman
├──git
|  └──github.com
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
❯ secman insert git/github.com
Enter password for git/github.com: 
```

Inserting a password in to your vault is easy. If you wish to group multiple entries together, it can be accomplished by prepending a group name followed by a slash to the pass-name.

Here we are adding mint.com to the password store within the money group.

### Inserting a file

```sh
❯ secman insert money/budget.csv budget.csv
```

Adding a file works almost the same as insert. Instead it has an extra argument. The file that you want to add to your vault is the final argument.

### Retrieving a password

```code
❯ secman show git/github.com
Enter master password:
dolladollabills$$1
```

Show is used to display a password in standard out.

### Rename a password

```code
❯ secman rename mney/mint.com
Enter new site name for mney/mint.com: git/github.com
```

If a password is added with the wrong name it can be updated later. Here we rename our mint.com site after misspelling the group name.

### Updating/Editing a password

```code
❯ secman edit dev/dev.to
Enter new password for dev/dev.to:
```

If you want to securely update a password for an already existing site, the edit command is helpful.

### Generating a password

```code
❯ secman generate
%L4^!s,Rry!}s:U<QwliL{vQ
Kow321-!tr}:232

❯ secman generate 8
#%Xy1t7E
```

secman can also create randomly generated passwords. The default length of secman generated passwords is 24 characters. This length can be changed by passing an optional length to the generate subcommand.

### Searching the vault

```sh
❯ secman find money
└──money
   └──mint.com

❯ secman ls money
└──money
   └──mint.com
```

`find` and `ls` can both be used to search for all sites that contain a particular substring. It's good for printing out groups of sites as well. `secman ls` is an alias of `secman find`.

### Deleting a vault entry

```sh
❯ secman
├──bb
|  └──ff
├──something
|  └──somethingelse.com
└──twiinsen.com
   └──bbbbb

❯ secman remove bb/ff

❯ secman
├──something
|  └──somethingelse.com
└──twiinsen.com
   └──bbbbb
```

remove is used for removing sites from the password vault. `secman rm` is an alias of `secman remove`.

### Getting Help

```code
❯ secman --help
```

All subcommands support the `--help` flag.

## CRYPTOGRAPHY DETAILS

###### Generating Passwords

Password generation takes place in the pc package by using the GeneratePassword function. GeneratePassword creates a random password by reading a large amount of randomness using the `func Read([]byte) (int, error)` function in the `crypto/rand` package.

The block of randomness is then read byte-by-byte. Printable characters that match the desired password specification (uppercase, lowercase, symbols, and digits) are then included in the generated password.

###### Adding A Site

When a site is added to the password store, a new public private key pair is generated. The newly generated private key, the user's master public key, and a securely generated nonce are used to encrypt the sites data.

The encryption and key computation are done using the `golang.org/x/crypto/nacl/box` package which uses Curve25519, XSalsa20, and Poly1305 to encrypt and authenticate the site's data.

After the site information is added, the site's generated private key is thrown away.

[goUrl]: https://goland.org