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

def core
    system("curl -fsSL https://secman-team.github.io/install/install.sh | bash")
end

sm = "secman".cyan

if $l == $c
    al = "is already up-to-date and it's the latest release"
    puts "#{sm} #{al} #{$l.yellow}"
    
elsif $l != $c
    pre_upgrade()
    core()

    puts "#{sm} was upgraded successfully"
end
