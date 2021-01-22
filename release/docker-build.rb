puts "Enter docker tag: "
tag = gets

system("docker build -t abdcodedoc/secman:#{tag} .")
