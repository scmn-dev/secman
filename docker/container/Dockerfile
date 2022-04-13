### image name: smcr/secman ###

FROM debian:11

### begin ###

### variables ###
ENV UPD="apt-get update"
ENV UPD_s="sudo $UPD"
ENV INS="apt-get install"
ENV INS_s="sudo $INS"
ENV PKGS="zip unzip multitail curl lsof wget ssl-cert asciidoctor apt-transport-https ca-certificates gnupg-agent bash-completion build-essential htop jq software-properties-common less llvm locales man-db nano vim ruby-full"
ENV BUILDS="build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev libbz2-dev"
ENV SECMAN_DOCKER_UTILS_URL="https://scmn-dev.github.io/secman/scripts/docker"

RUN $UPD && $INS -y $PKGS && $UPD && \
    locale-gen en_US.UTF-8 && \
    mkdir /var/lib/apt/abdcodedoc-marks && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* && \
    $UPD

ENV LANG=en_US.UTF-8

### git ###
RUN $INS -y git && \
    rm -rf /var/lib/apt/lists/* && \
    $UPD

### sudo ###
RUN $UPD && $INS -y sudo && \
    adduser --disabled-password --gecos '' secman_yo && \
    adduser secman_yo sudo && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

### nodejs & npm ###
RUN curl -sL https://deb.nodesource.com/setup_17.x -o nodesource_setup.sh && \
    sudo bash nodesource_setup.sh && \
    $INS_s nodejs build-essential -y && \
    sudo rm -rf nodesource_setup.sh && \
    $UPD_s
RUN npm i -g npm@latest

ENV HOME="/home/secman_yo"
WORKDIR $HOME
USER secman_yo

### secman ###
RUN sudo curl -sL https://u.secman.dev | bash
RUN sudo npm i -g @secman/scc

### zsh ###
ENV src=".zshrc"

RUN $INS_s zsh -y
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    $UPD_s && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

### rm old ~/.zshrc ###
RUN sudo rm -rf $src

### wget new files ###
RUN wget $SECMAN_DOCKER_UTILS_URL/zshrc -O $src
RUN wget $SECMAN_DOCKER_UTILS_URL/README

CMD /bin/bash -c "cat README && zsh"

### end ###
