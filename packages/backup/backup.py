import os
import subprocess as sp

# git & github
SM_GH_UN = sp.getoutput("git config user.name")
create = "gh repo create {}/.secman -y --private".format(SM_GH_UN)

# dirs
SECDIR = "~/.secman.bk"
SECDIR_primary = "~/.secman"
cd_SECDIR = "cd {}".format(SECDIR)

# pkgs
ghraw_url = "https://raw.githubusercontent.com"

install_brew = "/bin/bash -c \"$(curl -fsSL {}/Homebrew/install/HEAD/install.sh)\"".format(ghraw_url)
install_cgit = "/bin/bash -c \"$(curl -fsSL {}/Dev-x-Team/corgit/main/setup)\"".format(ghraw_url)
install_verx = "/bin/bash -c \"$(curl -fsSL {}/abdfnx/verx/HEAD/install.sh)\"".format(ghraw_url)

brew_gh = "brew install gh"

def _help():
    pass

def repo_work():
    csi = "cgit secman-ibk"
    rdm = 'touch {}/README.md && echo "# My secman backup passwords - {}" >> {}/README.md'.format(SECDIR, SM_GH_UN, SECDIR)

    # copy ~/.secman to ~/.secman.bk
    sp.getoutput("cp -rf {} {} && cd {}".format(SECDIR_primary, SECDIR, SECDIR))
    sp.getoutput("git init")
    sp.getoutput(rdm)
    sp.getoutput(create)
    sp.getoutput(csi)

def repo():
    try:
        devnull = open(os.devnull)
        sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
        repo_work()
    except OSError:
        sp.getoutput(brew_gh)

        try:
            devnull = open(os.devnull)
            sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
            repo_work()
        
        except OSError:
            sp.getoutput(install_brew)

            try:
                devnull = open(os.devnull)
                sp.Popen(["brew"], stdout=devnull, stderr=devnull).communicate()
                sp.getoutput(brew_gh)

                try:
                    devnull = open(os.devnull)
                    sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
                    repo_work()
                
                except OSError:
                    print("sorry, there's an error while initialize a backup, try again")
        
            except OSError:
                print("you should install brew\nhttps://brew.sh")
        return False
    
    return True

def _ph():
    pass

def _pl():
    pass

def _clone():
    pass

def badUsage():
    pass

def version():
    pass

repo()
