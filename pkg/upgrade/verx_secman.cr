l = `verx secman-team/secman -l`
c = `secman verx`

if l == c
    puts "same"
elsif l != c
    puts "not same"
end
