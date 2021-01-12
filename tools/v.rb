#!/usr/bin/ruby -W

require 'open-uri'

IO.copy_stream(URI.open("https://raw.githubusercontent.com/abdfnx/secman/HEAD/tools/v_checker.rb"), 'destination.rb')

require './destination.rb'

if $l != $c
    puts "new"
end
