require 'colorize'

image = `docker images -q abdcodedoc 2> /dev/null`
wel = "Welcome to secman vm".green

if image == ""
    puts wel
    system("docker pull abdcodedoc/secman:latest")
    system("docker run -t -i --privileged abdcodedoc/secman")
else
    puts wel
    system("docker run -t -i --privileged abdcodedoc/secman")
end
