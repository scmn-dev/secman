require "colorize"

system("if [ -d ~/.secman ]; then rm -rf ~/.secman; fi")
system("if [ -d ~/.secman.bk ]; then rm -rf ~/.secman.bk; fi")

puts "#{"secman".cyan.bold} was cleaned #{"successfully".green.bold}"
