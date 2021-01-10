#!/bin/bash

# Installation
# 1- check if curl command is exist
# 2- check if wget command // /////
# 3- check if brew command // /////

UNAME=$(uname)
smUrl=https://raw.githubusercontent.com/abdfnx/secman/HEAD/release/osx/secman
sm_unUrl=https://raw.githubusercontent.com/abdfnx/secman/HEAD/packages/secman-un
sm_syUrl=https://raw.githubusercontent.com/abdfnx/secman/HEAD/api/secman-sync
smLocLD=/usr/local/bin

successInstall() {
    echo "yesss, secman was installed successfully 😎, you can type secman --help"
}

installBrew() {
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
}

installSecman_&_Tools() {
    # secman
    sudo wget -P $smLocLD $smUrl

    sudo chmod 755 $smLocLD/secman

    # secman-un
    sudo wget -P $smLocLD $sm_unUrl

    sudo chmod 755 $smLocLD/secman-un

    # secman-sync
    sudo wget -P $smLocLD $sm_syUrl

    sudo chmod 755 $smLocLD/secman-sync
}

checkWget() {
    if [ -x "$(command -v wget)" ]; then
        installSecman_&_Tools
    else
        brew install wget

        if [ -x "$(command -v wget)" ]; then
            installSecman_&_Tools
        fi
    fi
}

gitAndBrew() {
    if [ -x "$(command -v git)" ]; then
        installBrew

        if [ -x "$(command -v brew)" ]; then
            checkWget
        fi
    fi
}

checkGit() {
    xcode-select --install

    gitAndBrew
}

mainCheck() {
    if [ -x "$(command -v brew)" ]; then
        checkWget

    else
        if [ -x "$(command -v git)" ]; then
            installBrew

            if [ -x "$(command -v brew)" ]; then
                checkWget
            fi
        else
            checkGit
        fi
    fi
}

if [ -x "$(command -v curl)" ]; then
    mainCheck

    if [ -x "$(command -v secman)" ]; then
        successInstall
    else
        echo "Download failed 😔"
    fi

else
    echo "You should install curl"
    exit 0
fi
