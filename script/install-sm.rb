$gh_raw = "https://raw.githubusercontent.com/secman-team/install/HEAD"

def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared_gh_url = "#{$gh_raw}/install"

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            system("powershell.exe iwr -useb #{$gh_raw/HEAD}/win/install.ps1 | iex")
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

os()
