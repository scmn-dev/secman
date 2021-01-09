#!/bin/bash

# Installation
# 1- check if curl command is exist
# 2- check if wget command // //
# 3- check if brew command // //
# 4- some Linux platforms don't have git installed, so it's well checking is git command is exist

UNAME=$(uname)
smUrl=https://raw.githubusercontent.com/abdfnx/secman/HEAD/release/linux/secman
sm_unUrl=https://raw.githubusercontent.com/abdfnx/secman/HEAD/core/secman-un
smLocLD=/usr/local/bin

successInstall() {
    echo "yesss, secman was installed successfully 😎, you can type secman --help"
}

installBrew() {
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
}

installSecman() {
    if [ -x "$(command -x sudo)" ]; then
        sudo wget -P $smLocLD $smUrl

        sudo chmod 755 $smLocLD/secman
    else
        wget -P $smLocLD $smUrl

        chmod 755 $smLocLD/secman
    fi
}

installSecman_un() {
    if [ -x "$(command -x sudo)" ]; then
        sudo wget -P $smLocLD $sm_unUrl

        sudo chmod 755 $smLocLD/secman-un
    else
        wget -P $smLocLD $sm_unUrl

        chmod 755 $smLocLD/secman-un
    fi
}

checkWget() {
    if [ -x "$(command -x wget)" ]; then
        installSecman

        installSecman_un
    else
        brew install wget

        if [ -x "$(command -x wget)" ]; then
            installSecman

            installSecman_un
        fi
    fi
}

gitAndBrew() {
    if [ -x "$(command -x git)" ]; then
        installBrew

        if [ -x "$(command -x brew)" ]; then
            checkWget
        fi
    fi
}

checkGit() {
    if [ -x "$(command -x sudo)" ]; then
        sudo apt install git
        gitAndBrew
    else
        apt install git
        gitAndBrew
    fi

}

mainCheck() {
    if [ -x "$(command -x brew)" ]; then
        checkWget

    else
        if [ -x "$(command -x git)" ]; then
            installBrew

            if [ -x "$(command -x brew)" ]; then
                checkWget
            fi
        else
            checkGit
        fi
    fi
}

if [ -x "$(command -x curl)" ]; then
    mainCheck

    if [ -x "$(command -x secman)" ]; then
        successInstall
    else
        echo "Download failed 😔"
    fi

else
    echo "You should install curl"
    exit 0
fi

if [ -x "$(command -x secman)" ]; then
    echo "Enter your github username: "
    echo " "
    read sgu

    echo "export $SM_GH_UN=$sgu"
fi
