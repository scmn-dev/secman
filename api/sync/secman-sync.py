import os
import sys, getopt
import subprocess as sp

# dirs
SECDIR = "~/.secman"
SECDIR_url = ".secman"
cd_SECDIR = "cd {}".format(SECDIR)

# git & github
SM_GH_UN = sp.getoutput("git config user.name")
create = "gh repo create {}/{} -y --private".format(SM_GH_UN, SECDIR_url)
clone = "gh repo clone {}/{} {}".format(SM_GH_UN, SECDIR_url, SECDIR)

# pkgs
ghraw_url = "https://raw.githubusercontent.com"

def install(_url):
    return "/bin/bash -c \"$(curl -fsSL {}/{})\"".format(ghraw_url, _url)

install_brew = install("Homebrew/install/HEAD/install.sh")
install_cgit = install("secman-team/corgit/main/setup")
install_verx = install("abdfnx/verx/HEAD/install.sh")

brew_gh = "brew install gh"

def _help():
    print("Flags:\n\t-h | --help: help about any command\n\t-s | --sync: create private github repo and sync your passwords on it by git\n\t-c | --clone: clone .secman repo\n\t-p | --push: push and commit a new secret\n\t-l | --pull: pull secret/s")

def sync():
    csi = "cgit secman-i"
    rdm = 'touch {}/README.md && echo "# My secman passwords - {}" >> {}/README.md'.format(SECDIR, SM_GH_UN, SECDIR)

    os.system("{} && {} && {} && {}".format(cd_SECDIR, rdm, create, csi))

def repo():
    try:
        devnull = open(os.devnull)
        sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
        sync()
    except OSError:
        sp.getoutput(brew_gh)

        try:
            devnull = open(os.devnull)
            sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
            sync()
        
        except OSError:
            sp.getoutput(install_brew)

            try:
                devnull = open(os.devnull)
                sp.Popen(["brew"], stdout=devnull, stderr=devnull).communicate()
                sp.getoutput(brew_gh)

                try:
                    devnull = open(os.devnull)
                    sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
                    sync()
                
                except OSError:
                    print("sorry, there's an error while initialize a backup, try again")
        
            except OSError:
                print("you should install brew\nhttps://brew.sh")
        return False
    
    return True

def _ph():
    os.system("{} && cgit ph".format(cd_SECDIR))

def _pl():
    os.system("{} && cgit pl".format(cd_SECDIR))

def _clone():
    os.system(clone)

def badUsage():
    print("Flag not recognized.\nFor an overview of the command, execute: secman-sync -h")

def version():
    os.system("secman ver")

def main(argv):
    try:
      opts, args = getopt.getopt(argv, "hscvpl", ["help", "sync", "clone", "push", "pull"])

    except getopt.GetoptError:
      badUsage()
      sys.exit(2)
    
    for opt, arg in opts:
        if opt in ("-h", "--help"):
            _help()
            sys.exit()
        
        elif opt in ("-s", "--sync"):
            sync()
            sys.exit()

        elif opt in ("-c", "--clone"):
            _clone()
            sys.exit()

        elif opt in ("-v", "--version"):
            version()
            sys.exit()
        
        elif opt in ("-p", "--push"):
            _ph()
            sys.exit()

        elif opt in ("-l", "--pull"):
            _pl()
            sys.exit()

        else:
            badUsage()
            sys.exit()
        
if __name__ == "__main__":
   main(sys.argv[1:])
