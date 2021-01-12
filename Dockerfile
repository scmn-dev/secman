# The reason of choosing gitpod/workspace-full, Because it contains everything, like brew, apt/apt-get, npm
FROM gitpod/workspace-full

# pkgs
ARG PKGS="curl wget"
ARG GITHUB_URL="https://raw.githubusercontent.com"

RUN sudo apt-get update
RUN sudo apt-get install $PKGS -y
RUN brew install dep
RUN sudo apt-get update

# install cli apps (gh, corgit, manx, verx)
RUN brew install gh
RUN /bin/bash -c "$(curl -fsSL $GITHUB_URL/Dev-x-Team/corgit/main/setup)"
RUN npm i -g @abdfnx/manx
RUN /bin/bash -c "$(curl -fsSL $GITHUB_URL/abdfnx/verx/HEAD/install.sh)"
RUN sudo apt-get update

# zsh
RUN brew install zsh

# secman
RUN /bin/bash -c "$(curl -fsSL $GITHUB_URL/abdfnx/secman/HEAD/packages/install_linux.sh)"
RUN sudo apt-get update

WORKDIR /core
