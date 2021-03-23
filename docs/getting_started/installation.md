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

### After installation

To test your installation, run `secman ver`. If this prints the Secman version to the console the installation was successful.

Use secman help to see help text documenting Secman's flags and usage. Get a detailed guide on the CLI here.

```sh
secman -h
```

### Upgrading

To upgrade a previously installed version of Secman, you can run:

`secman upg/upgrade` This will fetch the latest release from github.com/secman-team/secman/releases, unzip it, and replace your current executable with it.
