$gh_raw = "https://raw.githubusercontent.com"

def deps()
    system("/bin/bash -c \"$(curl -fsSL #{$gh_raw}/Dev-x-Team/corgit/main/setup)\"")
    system("/bin/bash -c \"$(curl -fsSL #{$gh_raw}/abdfnx/verx/HEAD/install.sh)\"")
end

def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared_gh_url = "#{$gh_raw}/secman-team/install/HEAD/install"

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

os()
