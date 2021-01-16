import subprocess as sp

f = "touch xcode.secman"

_file = sp.getoutput(f)

print(_file)
