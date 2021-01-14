require 'optparse'
$l = `verx abdfnx/secman -l`
# $c = "./test"

def check()
    if $l == ""
        # _n
        puts "same"
    else
        puts "not same"
    end
end

OptionParser.new do |opts|
  opts.on("-c", "--check", "check the version") do |c|
    check
  end
end.parse!
