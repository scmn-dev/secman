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
        shared_gh_url = "https://raw.githubusercontent.com/secman-team/install/HEAD/install"
        lin = "linux".yellow
        osx = "osx".black

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            puts "secman upg command is only supported for #{lin} and #{osx}"
        when /darwin|mac os/
            :macosx
            system("curl -fsSL #{shared_gh_url}_osx.sh | bash")
        when /linux/
            :linux
            system("curl -fsSL #{shared_gh_url}_linux.sh | bash")
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
