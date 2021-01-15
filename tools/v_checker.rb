require 'optparse'
$l = `verx abdfnx/secman -l`
$c = `secman verx`

def check()
    if $l != $c
        # _n
        puts "not same"
    else
        puts "same"
    end
end

OptionParser.new do |opts|
  opts.on("-c", "--check", "check the version") do |c|
    check
  end
end.parse!

