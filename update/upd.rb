require 'colorize'
require 'rbconfig'

$l = `verx abdfnx/secman -l`
$c = `secman verx`

smLoc = "/usr/local/bin/secman"

def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared_gh_url = "https://raw.githubusercontent.com/abdfnx/secman/HEAD/tools/install"

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            system("/bin/bash -c \"$(curl -fsSL #{shared_gh_url}_win.sh)\"")
        when /darwin|mac os/
            :macosx
            system("/bin/bash -c \"$(curl -fsSL #{shared_gh_url}_osx.sh)\"")
        when /linux/
            :linux
            system("/bin/bash -c \"$(curl -fsSL #{shared_gh_url}_linux.sh)\"")
        else
            raise Error::WebDriverError, "unknown os: #{host_os.inspect}"
        end
    )
end

sm = "secman".blue

if $l == $c
    al = "is already up-to-date and it's the latest release"
    puts "#{sm} #{al} #{$l.yellow}"
    
elsif $l != $c
    system("sudo rm -rf #{smLoc}*")

    os

    puts "#{sm} was updated successfully"
end
