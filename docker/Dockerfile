FROM gitpod/workspace-full

WORKDIR ../
# curl & wget

ARG PKGS="curl wget"

RUN sudo apt-get update
RUN sudo apt-get install $PKGS -y
RUN sudo apt-get update

# install cli apps (gh, corgit, manx, verx)
RUN brew install gh
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Dev-x-Team/corgit/main/setup)"
# 👇 to fix some pull errors
RUN git config pull.ff only
RUN npm i -g @abdfnx/manx
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/verx/HEAD/install.sh)"

# secman
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/HEAD/packages/install_linux.sh)"
RUN sudo apt-get update
