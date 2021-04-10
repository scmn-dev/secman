require "colorize"

$l = `powershell.exe ./vx.ps1 -l`
$c = `secman verx`

$loc = "$HOME/sm/AppData/Local/secman"

def pre_upgrade
    system("powershell.exe Remove-Item #{$loc} -Recurse -Force")
end

def core
    system("powershell.exe Remove-Item $HOME/sm -Recurse -Force")
    system("powershell.exe iwr -useb https://secman-team.github.io/install.ps1 | iex")
end

sm = "secman".cyan

if $l == $c
    al = "is already up-to-date and it's the latest release"
    puts "#{sm} #{al} #{$l.yellow}"
    
elsif $l != $c
    core()

    puts "#{sm} was upgraded successfully 🎊"
end
