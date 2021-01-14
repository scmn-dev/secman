require 'optparse'
require 'colorize'

$l = `verx abdfnx/secman -l`
$c = `secman ver`

def _n()
    nr = "there's a new release of secman is avalaible:".yellow
    ly = $l.cyan.bold
    up = "to update it run".yellow
    smu = "secman upd".blue
    puts new_r = "#{nr} #{ly}#{up} #{smu}"
end

def check()
    if "#{$l}" != "#{$c}"
        _n
    end

    if "#{$l}" == "#{$c}"
        puts ""
    end
end

options = {}
OptionParser.new do |opts|
  opts.on("-c", "--check", "check the version") do |c|
    check
  end
end.parse!
