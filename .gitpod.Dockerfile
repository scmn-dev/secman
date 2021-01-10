FROM gitpod/workspace-full
# curl & wget

ARG PKGS="curl wget"

RUN sudo apt-get update
RUN sudo apt-get install $PKGS -y
RUN sudo apt-get update

# terraform
RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
RUN sudo apt-add-repository "deb [arch=$(dpkg --print-architecture)] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
RUN sudo apt-get install terraform
RUN sudo apt-get update

# install cli apps (gh, corgit, manx)
RUN brew install gh
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Dev-x-Team/corgit/main/setup)"
RUN npm i -g @abdfnx/manx

# secman
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/HEAD/packages/install_linux.sh)"
RUN sudo apt-get update
