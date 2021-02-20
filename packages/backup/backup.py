import os
import sys, getopt
import subprocess as sp
import pathlib
import platform

# dirs
SECDIR = "~/.secman.bk"
SECDIR_primary = "~/.secman"
SECDIR_url = "secman.bk"
SECDIR_primary_url = ".secman"
cd_SECDIR = "cd {}".format(SECDIR)

# git & github
SM_GH_UN = sp.getoutput("git config user.name")
create = "gh repo create {}/{} -y --private".format(SM_GH_UN, SECDIR_url)
clone = "gh repo clone {}/{} {}".format(SM_GH_UN, SECDIR_url, SECDIR)
remote = "git remote set-url origin https://github.com/{}/.secman".format(
    SM_GH_UN)

# pkgs
ghraw_url = "https://raw.githubusercontent.com"

def install(_url):
    return "curl -fsSL {}/{} | bash".format(ghraw_url, _url)

install_brew = install("Homebrew/install/HEAD/install.sh")

brew_gh = "brew install gh"
choco_gh = "choco install gh"

def _help():
    print(
        "Flags:\n\t-h: help\n\t-i: init ~/.secman.bk folder\n\t-c: clone the backup repo\n\t-p: push new backup passwords\n\t-l: pull passwords\n\t-M: make ~/.secman.bk is the main folder\n"
    )

def repo_work():
    csi = "bash ~/sm/cgit secman-ibk"
    rdm = 'touch {}/README.md && echo "# My secman backup passwords - {}" >> {}/README.md'.format(
        SECDIR, SM_GH_UN, SECDIR)

    # copy ~/.secman to ~/.secman.bk
    sp.getoutput("cp -rf {} {}".format(SECDIR_primary, SECDIR))
    os.system("{} && rm -rf .git && git init && {} && {} && {}".format(
        cd_SECDIR, rdm, create, csi))

def install_installer():
    if platform.system() == "Linux" or platform.system() == "Darwin":
        sp.getoutput(brew_gh)
    elif platform.system() == "Windows":
        sp.getoutput(choco_gh)

def repo():
    try:
        devnull = open(os.devnull)
        sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
        repo_work()
    except OSError:
        install_installer()

        try:
            devnull = open(os.devnull)
            sp.Popen(["gh"], stdout=devnull, stderr=devnull).communicate()
            repo_work()

        except OSError:
            if platform.system() == "Linux" or platform.system() == "Darwin":
                sp.getoutput(install_brew)
            elif platform.system() == "Windows":
                print(
                    "you should install Chocolatey from this url https://chocolatey.org/install"
                )

            try:
                devnull = open(os.devnull)

                if platform.system() == "Linux" or platform.system(
                ) == "Darwin":
                    sp.Popen(["brew"], stdout=devnull,
                             stderr=devnull).communicate()
                elif platform.system() == "Windows":
                    sp.Popen(["choco"], stdout=devnull,
                             stderr=devnull).communicate()

                install_installer()

                try:
                    devnull = open(os.devnull)
                    sp.Popen(["gh"], stdout=devnull,
                             stderr=devnull).communicate()
                    repo_work()

                except OSError:
                    print(
                        "sorry, there's an error while initialize a backup, try again"
                    )

            except OSError:
                if platform.system() == "Linux" or platform.system(
                ) == "Darwin":
                    print("you should install brew\nhttps://brew.sh")
                elif platform.system() == "Windows":
                    print(
                        "you should install Chocolatey from this url https://chocolatey.org/install"
                    )
        return False

    return True

def _ph():
    os.system("{} && bash ~/sm/cgit ph".format(cd_SECDIR))

def _pl():
    os.system("{} && bash ~/sm/cgit pl".format(cd_SECDIR))

def _clone():
    os.system(clone)

def badUsage():
    print(
        "Flag not recognized.\nFor an overview of the command, execute: secman-sync backup -h"
    )

def version():
    os.system("secman ver")

def make_main():
    """
    1. check if ~/.secman is not exist
     then
        1. copy ~/.secman.bk to ~/.secman
        2. change git remote

    2. if ~/.secman is exist
     then
        execute this msg '~/.secman is exist'
    """

    home = os.path.expanduser('~/.secman')
    SECDIR_path = pathlib.Path(home)

    if not SECDIR_path.exists():
        os.system("cp -rf {} {}".format(SECDIR, SECDIR_primary))
        os.system("cd {} && {}".format(SECDIR_primary, remote))

    else:
        print("{} is exist".format(SECDIR_primary))

def main(argv):
    try:
        opts, args = getopt.getopt(
            argv, "hicvplM", ["help", "init", "clone", "push", "pull", "Main"])

    except getopt.GetoptError:
        badUsage()
        sys.exit(2)

    for opt, arg in opts:
        if opt in ("-h", "--help"):
            _help()
            sys.exit()

        elif opt in ("-i", "--init"):
            repo_work()
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

        elif opt in ("-m", "-M", "--Main"):
            make_main()
            sys.exit()

        else:
            badUsage()
            sys.exit()

if __name__ == "__main__":
    main(sys.argv[1:])
