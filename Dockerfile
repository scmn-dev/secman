# The reason of choosing gitpod/workspace-full, Because it contains everything, like brew, apt/apt-get, npm & more...
FROM gitpod/workspace-full

# pkgs
ARG PKGS="curl wget ruby-full"
ARG UPD="sudo apt-get update"
ARG GITHUB_URL="https://raw.githubusercontent.com"

RUN $UPD && \
    sudo apt-get install $PKGS -y && \
    brew install dep && \
    $UPD

# install cli apps (gh, corgit, manx and verx) & install deps from gemfile
RUN brew install gh && \
    /bin/bash -c "$(curl -fsSL $GITHUB_URL/Dev-x-Team/corgit/main/setup)" && \
    npm i -g @abdfnx/manx && \
    /bin/bash -c "$(curl -fsSL $GITHUB_URL/abdfnx/verx/HEAD/install.sh)" && \
    bundle install && \
    $UPD

# secman
RUN /bin/bash -c "$(curl -fsSL $GITHUB_URL/abdfnx/secman/HEAD/tools/install_linux.sh)" && \
    $UPD

# zsh & omz (oh my zsh)
RUN brew install zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    $UPD

WORKDIR /core
