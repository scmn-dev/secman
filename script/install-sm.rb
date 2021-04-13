def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        shared = "https://secman-team.github.io/install"

        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            system("powershell.exe iwr -useb #{shared}.ps1 | iex")

        when /linux|darwin|mac os/
            :linux_macos
            system("curl -fsSL #{shared}.sh | bash")
        else
            raise Error::WebDriverError, "unknown os: #{host_os.inspect}"
        end
    )
end

os()
