# Upgrade your secman if there's a new release.

Usage:
  `secman upg [flags]`

Aliases:
  **upg**, **upgrade**

Flags:
  **-h**, **--help**   help for upg

Note: `secman upg` & `secman-un` are only supported in **linux/mac**

## Linux/MacOS

> `update by upg command`

```sh
secman upg
```

> `by brew`

```sh
brew upgrade secman
```

## in Windows

```pwsh
scoop update secman
```

## How `upg` command works ?

secman use [**verx**](https://github.com/abdfnx/verx) for this

then secman run `ruby /home/sm/v_checker.rb -c` to check versions

if there's a new release, it's well show this message

```bash
there is a new release of secman is avalaible: THE_NEW_VERSION
to upgrade run secman upg
```

### v_checker.rb

_macOS/Linux_

```ruby
require "optparse"
require "colorize"

$l = `verx secman-team/secman -l`
$c = `secman verx`

def _n()
  ly = $l.cyan.bold
  nr = "\nthere's a new release of secman is avalaible:".yellow
  up = "to upgrade run".yellow
  smu = "secman upg".cyan
  puts new_r = "#{nr} #{ly}#{up} #{smu}"
end

def check()
  if $l != $c
    _n
  end
end

OptionParser.new do |opts|
  opts.on("-c", "--check", "check the version") do |c|
    check()
  end
end.parse!
```

_Windows_

first, secman run ~/sm/vx.ps1

```ruby
require "optparse"
require "colorize"

$l = `powershell.exe ./vx.ps1 -l`
$c = `secman verx`

def _n()
  ly = $l.cyan.bold
  nr = "\nthere's a new release of secman is avalaible:".yellow
  ug = "to upgrade".yellow
  scoop = "scoop upgrade secman".cyan
  puts new_r = "#{nr} #{ly}#{ug} go to https://github.com/secman-team/secman/releases\nor if you install secman by scoop run #{scoop}"
end

def check()
  if $l != $c
    _n
  end
end

OptionParser.new do |opts|
  opts.on("-c", "--check", "check the version") do |c|
    check()
  end
end.parse!
```

## upg.rb (macOS/Linux)

#### Working on add upg command to windows

```ruby
require "rbconfig"
require "colorize"

$l = `verx secman-team/secman -l`
$c = `secman verx`

$smLoc = "/usr/local/bin"

def pre_upgrade
    system("sudo rm -rf #{$smLoc}/secman*")
    system("sudo rm -rf #{$smLoc}/cgit*")
    system("sudo rm -rf #{$smLoc}/verx*")
    system("sudo rm -rf /home/sm")
end

def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        lin = "linux".yellow
        osx = "osx".black

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            puts "secman upg command is only supported for #{lin} and #{osx}"
        when /darwin|mac os/linux/
            :macosx
            system("curl -fsSL https://secman-team.github.io/install/install.sh | bash")
        else
            raise Error::WebDriverError, "unknown os: #{host_os.inspect}"
        end
    )
end

sm = "secman".cyan

if $l == $c
    al = "is already up-to-date and it's the latest release"
    puts "#{sm} #{al} #{$l.yellow}"
    
elsif $l != $c
    pre_upgrade()
    os()

    puts "#{sm} was upgraded successfully"
end
```

> if the user run `secman upg`
> it's well remove old secman files, like `secman`, `secman-sync`, `secman-un` and remove `/home/sm` folder
> then secman download the new .zip file from https://github.com/secman-team/secman/releases/latest
> after this, the new secman was installed
