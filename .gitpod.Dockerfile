FROM gitpod/workspace-full
# curl & wget

ARG PKGS="curl wget"

RUN apt-get update
RUN apt-get install $PKGS -y
RUN apt-get update

# install packages & cli apps (corgit, manx, secman)
RUN brew install gh
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Dev-x-Team/corgit/main/setup)"
RUN npm i -g @abdfnx/manx
# secman
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/HEAD/packages/install_linux.sh)"
RUN apt-get update
