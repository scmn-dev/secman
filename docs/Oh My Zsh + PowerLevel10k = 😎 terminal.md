# Oh My Zsh + PowerLevel10k = 😎 terminal

Hi there, The developers always use the terminal, but it is boring and has no colors or shapes.
Today we will transform our terminals into a wonderful, colorful, supportive terminal that offers suggestions and has a memory, we'll do it by [OMZ][omzUrl] and design it by [PowerLevel10k][p10kUrl]...

## Pre-requisites

> if you're using [Windows](https://www.microsoft.com/en-us/windows), you can install and configure [WSL][wslUrl]

and I recommended to use [Ubuntu](https://ubuntu.com/) or [Debian](https://www.debian.org/) wsl plugin

* [Homebrew](https://brew.sh) is installed

## Setup [zsh](https://www.zsh.org/)

in the command line type

```sh
brew install zsh
```

type **zsh**

```sh
zsh
```

> ### Install [Oh My Zsh][omzUrl]

```sh
sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```

## [PowerLevel10k][p10kUrl]

* Install Powerlevel10k using the following command

```sh
# gh cli
gh repo clone romkatv/powerlevel10k $ZSH_CUSTOM/themes/powerlevel10k

# git
git clone https://github.com/romkatv/powerlevel10k.git $ZSH_CUSTOM/themes/powerlevel10k
```

Then you need to enable it, change the value of ZSH_THEME to following in `~/.zshrc` file :

```zsh
ZSH_THEME="powerlevel10k/powerlevel10k"
```

> ### Configure Powerlevel10k Theme

* Make sure your terminal font is `MesloLGS NF` or `Hack Nerd Font`.

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/vb56rxkxktyjaocbrh6s.png)

#### Cheat-sheet for Windows

if you've `Windows terminal` you can open your settings and  in **UNIX** preferences and add `fontFace` prop,
assign it to `MesloLGS NF` or `Hack Nerd Font`.

```json
{
  "guid": "{YOUR_UNIX_GUID}",
  "hidden": false,
  "name": "Ubuntu",
  "source": "Windows.Terminal.Wsl",
  "fontFace": "Hack Nerd Font",
  "snapOnInput": true,
  "useAcrylic": true
}
```

> Windows Terminal url in Microsoft Store: [url](https://www.microsoft.com/en-us/p/windows-terminal/9n0dx20hk701)
> Windows Terminal repo: [url](https://github.com/microsoft/terminal)

### **p10k configure**

type

```sh
p10k configure
```

![x](https://dev-to-uploads.s3.amazonaws.com/i/xf9fk2sgux1niog4vhpy.gif)

you can choose your style...

> ## Plugins (Optional, Good to have!)

### Clone plugins

* zsh-syntax-highlighting - It enables highlighting of commands whilst they are typed at a zsh prompt into an interactive terminal. This helps in reviewing commands before running them, particularly in catching syntax errors.

```sh
# gh cli
gh repo clone zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# git
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
```

* zsh-autosuggestions - It suggests commands as you type based on history and completions.

```sh
# gh cli
gh repo clone zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

# git
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```

### **ls** tools

* [colorls](https://github.com/athityakumar/colorls): A Ruby script that colorizes the `ls` output with color and icons

* [exa](https://the.exa.website): is a modern replacement for _ls_

#### colorls

```sh
sudo gem install colorls
```

#### warn 🙃

#### maybe you'll get some `gem` errors, you should fix it

> _Linux_

```sh
sudo apt install ruby-full
```

#### exa

```sh
brew install exa
```

> #### Activate the plugins

In `~/.zshrc` file replace the line starting with `plugins=()` to below line.

```zsh
plugins=( git zsh-syntax-highlighting zsh-autosuggestions )
```

> colorls

```zsh
alias ls="colorls"
alias la="colorls -al"
```

> or exa

```zsh
alias ls="exa"
alias la="exa -al"
```

> Some more official plugins - [ohmyzsh plugins][szpUrl]

after all these steps type

```sh
source ~/.zshrc
```

## Finally it should be like this 👇

> Mac

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/7yxpyhy9lj36ks178ywt.png)

> Windows(WSL) or Linux

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/suf79s6ur03owctdq0l4.png)

![Alt Text](https://dev-to-uploads.s3.amazonaws.com/i/gppctwt70q58skp24vmz.png)

That's it, see you next time 👋

[omzUrl]: https://ohmyz.sh
[p10kUrl]: https://github.com/romkatv/powerlevel10k
[wslUrl]: https://docs.microsoft.com/en-us/windows/wsl/
[szpUrl]: https://github.com/ohmyzsh/ohmyzsh/wiki/Plugins
