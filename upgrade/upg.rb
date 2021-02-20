require "rbconfig"
require "colorize"

$l = `verx secman-team/secman -l`
$c = `secman verx`

smLoc = "/usr/local/bin"

def _os
    @_os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared_gh_url = "https://raw.githubusercontent.com/secman-team/install/HEAD/install"
        lin = "linux".yellow
        osx = "osx".black

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            puts "secman upd command is only supported for #{lin} and #{osx}"
        when /darwin|mac os/
            :macosx
            system("$(curl -fsSL #{shared_gh_url}_osx.sh) | bash")
        when /linux/
            :linux
            system("$(curl -fsSL #{shared_gh_url}_linux.sh) | bash")
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
        when /darwin|mac os|linux/
            :macosx_linux
            system("sudo rm -rf #{smLoc}/secman*")
            system("sudo rm -rf #{smLoc}/cgit*")
            system("sudo rm -rf #{smLoc}/verx*")
            _os()
        else
            raise Error::WebDriverError, "unknown os: #{host_os.inspect}"
        end
    )

    puts "#{sm} was upgraded successfully"
end
