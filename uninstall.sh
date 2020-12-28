#!/bin/bash

goodBye() {
    echo "secman was uninstalled successfully... thanks you to trying secman"
}

smLoc=/usr/local/bin/secman
SECDIR=~/.secman

clearAllData() {
    echo -e "clear all data?\n[y/N]"
    read -n 1 accept

    if [[ $accept == "Y" || $accept == "y" ]]; then
        if [ -x "$(command -v sudo)" ]; then
            if [ -x "$(command -v manx)" ]; then
                sudo manx $SECDIR

            else
                installManx

                sudo manx $SECDIR
            fi
        fi
    elif [[ $accept == "" || $accept == "N" || $accept == "n" ]]; then
        echo "ok"
    fi

    echo "after clear, you can find your old .secman in ~/.local/share/Trash if you want to restore it"
}

installManx() {
    if [ -x "$(command -v sudo)" ]; then
        sudo npm i -g @abdfnx/manx
    else
        npm i -g @abdfnx/manx
    fi
}

manxopa() {
    if [ -x "$(command -v sudo)" ]; then
        sudo manx $smLoc
    else
        manx $smLoc
    fi

    clearAllData
}

rmOpa() {
    if [ -x "$(command -v sudo)" ]; then
        sudo rm -rf $smLoc
    else
        rm -rf $smLoc
    fi

    clearAllData
}

if [ -x "$(command -v secman)" ]; then
    echo "do you want to uninstall it by manx or rm"
    read un

    if [ "$un" == "manx" ]; then
        if [ -x "$(command -v manx)" ]; then
            manxopa
        else
            # if [ -x "$(command -v sudo)" ]; then
            #     sudo npm i -g @abdfnx/manx

            #     if [ -x "$(command -v manx)" ]; then
            #         manxopa
            #     fi
            # else
            #     npm i -g @abdfnx/manx

            #     if [ -x "$(command -v manx)" ]; then
            #         manxopa
            #     fi
            # fi

            installManx
            if [ -x "$(command -v manx)" ]; then
                manxopa
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
