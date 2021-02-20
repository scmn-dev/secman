$gh_raw = "https://raw.githubusercontent.com"

def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared_gh_url = "#{$gh_raw}/secman-team/install/HEAD/install"

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            system("curl -fsSL #{shared_gh_url}_win.sh | bash")
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
