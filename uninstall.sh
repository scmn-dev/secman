#!/bin/bash

goodBye() {
    echo "secman was uninstalled successfully... thanks you to trying secman"
}

smLoc=/usr/local/bin/secman
SECDIR=~/.secman

manxopa() {
    if [ -x "$(command -v sudo)" ]; then
        sudo manx $smLoc $SECDIR
    else
        manx $smLoc $SECDIR
    fi
}

rmOpa() {
    if [ -x "$(command -v sudo)" ]; then
        sudo rm -rf $smLoc $SECDIR
    else
        rm -rf $smLoc $SECDIR
    fi
}

if [ -x "$(command -v secman)" ]; then
    echo "do you want to uninstall it by manx or rm"
    read un

    if [ "$un" == "manx" ]; then
        if [ -x "$(command -v manx)" ]; then
            manxopa
        else
            if [ -x "$(command -v sudo)" ]; then
                sudo npm i -g @abdfnx/manx

                if [ -x "$(command -v manx)" ]; then
                    manxopa
                fi
            else
                npm i -g @abdfnx/manx

                if [ -x "$(command -v manx)" ]; then
                    manxopa
                fi
            fi
        fi
    elif [ "$un" == "rm" ]; then
        rmOpa
    fi

    if ! [ -f "$smLoc" ]; then
        goodBye
    else
        echo "there's error while uninstalling secman"
    fi
else
    echo "there's no secman 😐"
fi
