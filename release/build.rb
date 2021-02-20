require "colorize"

puts "Enter the SECRET_KEY"
sk = gets.chomp

require "./secret_key"

if sk == "#{$code}"
    core()
else
    puts "wrong".red
end
