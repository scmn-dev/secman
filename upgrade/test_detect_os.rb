require "rbconfig"

def os
    @os ||= (
        host_os = RbConfig::CONFIG['host_os']
        case host_os
        when /mswin|msys|mingw|cygwin|bccwin|wince|emc/
            :windows
            puts "win"
        when /darwin|mac os/
            :macosx
            puts "mac"
        when /linux/
            :linux
            puts "linux"
        when /solaris|bsd/
            :unix
            puts "solaris/bsd"
        else
            raise Error::WebDriverError, "unknown os: #{host_os.inspect}"
        end
    )
end

os
