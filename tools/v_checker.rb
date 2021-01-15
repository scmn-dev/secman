require 'optparse'
require 'colorize'

$l = `verx abdfnx/secman -l`
$c = `secman verx`

def _n()
    ly = $l.cyan.bold
    nr = "there's a new release of secman is avalaible:".yellow
    up = "to update it run".yellow
    smu = "secman upd".blue
    puts new_r = "#{nr} #{ly}#{up} #{smu}"
end

def check()
    if $l != $c
        _n
    end
end

OptionParser.new do |opts|
  opts.on("-c", "--check", "check the version") do |c|
    check
  end
end.parse!
