FROM debian:latest

# vars
ARG UPD="apt-get update"
ARG UPD_s="sudo $UPD"
ARG INS="apt-get install"
ARG INS_s="sudo $INS"
ARG APT_REPO="add-apt-repository"
ARG APT_REPO_s="sudo $APT_REPO"
ENV GITHUB_URL="https://raw.githubusercontent.com"
ENV PKGS="zip unzip multitail curl lsof wget ssl-cert asciidoctor apt-transport-https ca-certificates gnupg-agent bash-completion build-essential htop jq software-properties-common less llvm locales man-db nano vim ruby-full "
ENV BUILDS="build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev libbz2-dev"

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

# sudo
RUN $UPD && $INS -y sudo && \
    adduser --disabled-password --gecos '' smx && \
    adduser smx sudo && \
    echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

### docker ###
LABEL abdcodex/layer=tool-docker
LABEL abdcodex/test=tools/tool-docker.yml
USER root

RUN $UPD_s
RUN curl -o /var/lib/apt/abdcodedoc-marks/docker.gpg -fsSL https://download.docker.com/linux/debian/gpg && \
    sudo apt-key add /var/lib/apt/abdcodedoc-marks/docker.gpg && \
    $APT_REPO_s "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs)  stable" && \
    $UPD_s && \
    $INS_s -y docker-ce docker-ce-cli containerd.io docker-compose && \
    sudo cp /var/lib/dpkg/status /var/lib/apt/abdcodedoc-marks/tool-docker.status && \
    sudo apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* && \
    $UPD_s

ENV HOME="/home/smx"
WORKDIR $HOME
USER smx

### homebrew ###
LABEL abdcodex/layer=tool-brew
LABEL abdcodex/test=tools/tool-brew.yml

RUN $UPD_s
ENV TRIGGER_BREW_REBUILD=1
RUN mkdir ~/.cache && /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
ENV PATH="$PATH:/home/linuxbrew/.linuxbrew/bin:/home/linuxbrew/.linuxbrew/sbin/" \
    MANPATH="$MANPATH:/home/linuxbrew/.linuxbrew/share/man" \
    INFOPATH="$INFOPATH:/home/linuxbrew/.linuxbrew/share/info" \
    HOMEBREW_NO_AUTO_UPDATE=1

### nodejs & npm ###
RUN curl -sL https://deb.nodesource.com/setup_15.x -o nodesource_setup.sh && \
    sudo bash nodesource_setup.sh && \
    $INS_s nodejs build-essential -y && \
    sudo rm -rf nodesource_setup.sh && \
    $UPD_s

# install pkgs
RUN $UPD_s && \
    brew install dep && \
    sudo gem install bundler && \
    $UPD_s

# install cli apps (gh, corgit, manx and verx) & install ruby deps from gemfile
RUN brew install gh && \
    /bin/bash -c "$(curl -fsSL $GITHUB_URL/Dev-x-Team/corgit/main/setup)" && \
    sudo npm i -g @abdfnx/manx && \
    /bin/bash -c "$(curl -fsSL $GITHUB_URL/abdfnx/verx/HEAD/install.sh)" && \
    $UPD_s

COPY Gemfile ./
RUN sudo bundle install && sudo rm -rf Gemfile*

### secman ###
RUN /bin/bash -c "$(curl -fsSL $GITHUB_URL/abdfnx/secman/HEAD/tools/install_linux.sh)" && \
    $UPD_s

# zsh
RUN brew install zsh
ENV src=".zshrc"
RUN zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" && \
    $UPD_s && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

RUN echo 'ZSH_THEME="af-magic"' >> $src && \
    echo 'plugins=( git zsh-syntax-highlighting zsh-autosuggestions )' >> $src

ENTRYPOINT ["zsh"]
