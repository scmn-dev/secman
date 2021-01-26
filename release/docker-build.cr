puts "Enter docker tag: "
tag = gets

system("docker build -t abdcodedoc/secman:#{tag} .")
system("docker push -t abdcodedoc/secman:#{tag}")
