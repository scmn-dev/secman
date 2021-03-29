# Installation

Secman works on Windows, Linux, and macOS.

## Download and install

secman provides convenience scripts to download and install the binary.

### Using Shell (macOS and Linux)

```sh
curl -fsSL https://secman-team.github.io/install/install.sh | bash
```

### Using PowerShell (Windows)

```sh
iwr -useb https://secman-team.github.io/install/install.ps1 | iex
```

### Using [Homebrew](https://brew.sh) (macOS and Linux)

```sh
brew tap secman-team/smx
brew install secman
```

### Using [Scoop](https://scoop.sh) (Windows)

```pwsh
scoop bucket add secman https://github.com/secman-team/secman
scoop install secman
```

### MSI Installer

> MSI installer is available for download on the [releases](https://github.com/secman-team/secman/releases/latest).

Secman binaries can also be installed manually, by downloading a zip file at github.com/secman-team/secman/releases.

### Installing from source

> `secman` requires [Go](https://golang.org) version 1.11+

If `go` is not installed, follow steps on the [Go website](https://golang.org/doc/install).

1. clone secman repo

    ```sh
    # gh cli
    gh repo clone secman-team/secman

    # git
    git clone https://github.com/secman-team/secman

    # after clone
    cd secman
    ```

2. Build and install it

    ```sh
    # linux/macOS: by default, it's installs to '/usr/local'; maybe you'll need sudo
    # windows: it's installs to '%AppData\Local%';
    make
    ```

3. Run `secman ver` to check if it worked.

### After installation

To test your installation, run `secman ver`. If this prints the Secman version to the console the installation was successful.

Use secman help to see help text documenting Secman's flags and usage. Get a detailed guide on the CLI here.

```sh
secman -h
```

### Upgrading

To upgrade a previously installed version of Secman, you can run:

`secman upg/upgrade` This will fetch the latest release from **https://github.com/secman-team/secman/releases**, unzip it, and replace your current executable with it.
