require 'optparse'
require 'colorize'

image = `docker images -q abdcodedoc 2> /dev/null`
wel = "Welcome to".green
smvm = "secman vm".cyan

options = {}
OptionParser.new do |opts|
  opts.banner = "secman vm (virtual machine) with #{"docker".blue}"

  opts.on("-l", "--login", "login to your vm") do |l|
    puts "#{wel} #{smvm}"

    if image == ""
        system("docker pull abdcodedoc/secman:latest")
        # system("docker run -t -i --privileged abdcodedoc/secman")
    else
        # system("docker run -t -i --privileged abdcodedoc/secman")
    end
  end
  opts.on("-q", "--quit", "quit from secman vm") do |q|
    system("exit")
  end
end.parse!
