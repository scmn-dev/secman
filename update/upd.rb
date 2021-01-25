require 'colorize'
require 'rbconfig'

$l = `verx abdfnx/secman -l`
$c = `secman verx`

smLoc = "/usr/local/bin"

def deps()
    system("/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Dev-x-Team/corgit/main/setup)\"")
    system("/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/abdfnx/verx/HEAD/install.sh)\"")
end

def _os
    @_os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared_gh_url = "https://raw.githubusercontent.com/secman-team/install/HEAD/install"

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            system("/bin/bash -c \"$(curl -fsSL #{shared_gh_url}_win.sh)\"")
            deps()
        when /darwin|mac os/
            :macosx
            system("/bin/bash -c \"$(curl -fsSL #{shared_gh_url}_osx.sh)\"")
            deps()
        when /linux/
            :linux
            system("/bin/bash -c \"$(curl -fsSL #{shared_gh_url}_linux.sh)\"")
            deps()
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
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            smLoc_win = "/usr/bin"

            system("rm -rf #{smLoc_win}/secman*")
            system("rm -rf #{smLoc_win}/cgit*")
            system("rm -rf #{smLoc_win}/verx*")
            deps()
        when /darwin|mac os|linux/
            :macosx:linux
            system("sudo rm -rf #{smLoc}/secman*")
            system("sudo rm -rf #{smLoc}/cgit*")
            system("sudo rm -rf #{smLoc}/verx*")
            deps()
        else
            raise Error::WebDriverError, "unknown os: #{host_os.inspect}"
        end
    )

    _os()

    puts "#{sm} was updated successfully"
end
