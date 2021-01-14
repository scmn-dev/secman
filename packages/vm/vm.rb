require 'optparse'
require 'colorize'

image = `docker images -q abdcodedoc/secman 2> /dev/null`
wel = "Welcome to".blue
smvm = "secman vm".cyan

def start
    system("CID=$(docker create -t -i abdcodedoc/secman) && docker start -a -i $CID")
end

OptionParser.new do |opts|
  opts.banner = "secman vm (virtual machine) with #{"docker".blue}"

  opts.on("-l", "--login", "login to your vm") do |l|
    puts "#{wel} #{smvm}"

    if image == ""
        system("docker pull abdcodedoc/secman:latest")
        start
    else
        start
    end
  end

  opts.on("-q", "--quit", "quit from secman vm") do |q|
    system("exit")
  end
end.parse!
