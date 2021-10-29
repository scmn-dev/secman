l = `verx scmn-dev/secman -l`
c = `secman verx`

if l == c
    puts "same"
elsif l != c
    puts "not same"
end
