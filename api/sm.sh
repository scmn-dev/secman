if [ -x "$(command -v curl)" ]; then
    curl --silent "https://api.github.com/repos/secman-team/secman/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
else
    echo "you should install curl"
fi
