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

        when /linux|darwin|mac os/
            :linux_macos
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
