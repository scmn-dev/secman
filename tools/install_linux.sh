#!/bin/bash

# Installation
# 1- check if curl command is exist
# 2- check if wget command // //
# 3- check if brew command // //
# 4- some Linux platforms don't have git installed, so it's well checking is git command is exist

GH_RAW_URL=https://raw.githubusercontent.com/abdfnx
smUrl=$GH_RAW_URL/secman/HEAD/release/linux/secman
sm_unUrl=$GH_RAW_URL/secman/HEAD/packages/secman-un
sm_syUrl=$GH_RAW_URL/secman/HEAD/api/sync/secman-sync
vrb=$GH_RAW_URL/secman/HEAD/tools/v_checker.rb
smLocLD=/usr/local/bin

successInstall() {
    echo "yesss, secman was installed successfully 😎, you can type secman --help"
}

installBrew() {
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
}

installSecman_Tools() {
    # install deps
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Dev-x-Team/corgit/main/setup)"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/verx/HEAD/install.sh)"

    # secman
    sudo wget -P $smLocLD $smUrl

    sudo chmod 755 $smLocLD/secman

    # secman-un
    sudo wget -P $smLocLD $sm_unUrl

    sudo chmod 755 $smLocLD/secman-un

    # secman-sync
    sudo wget -P $smLocLD $sm_syUrl

    sudo chmod 755 $smLocLD/secman-sync

    ${vrb}
}

checkWget() {
    if [ -x "$(command -v wget)" ]; then
        installSecman_Tools
    else
        brew install wget

        if [ -x "$(command -v wget)" ]; then
            installSecman_Tools
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
    if [ -x "$(command -v sudo)" ]; then
        sudo apt install git
        gitAndBrew
    else
        apt install git
        gitAndBrew
    fi

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
