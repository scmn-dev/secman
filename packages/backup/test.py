import subprocess as sp
import platform

f = "touch xcode.secman"

_file = sp.getoutput(f)

print(_file)

if platform.system() == "Linux":
    print("Linux is")
