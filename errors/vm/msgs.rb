require "optparse"
require "colorize"

wel = "Welcome to".blue
smvm = "secman vm".cyan

OptionParser.new do |opts|
  opts.on("-w", "--wel", "") do |l|
    puts "#{wel} #{smvm} !".bold
  end

  opts.on("-s", "--secmandocker", "") do |q|
    puts "secman vm (virtual machine) with #{"docker".blue}"
  end

  opts.on("-c", "--cid", "") do |q|
    puts "cid (container id) was created successfully".green
  end
end.parse!
