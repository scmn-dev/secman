require "colorize"

system("if [ -d ~/.secman ]; then rm -rf ~/.secman; fi")

puts "#{"secman".cyan.bold} was cleaned #{"successfully".green.bold}"
