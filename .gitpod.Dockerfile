FROM gitpod/workspace-full

RUN brew install zsh

WORKDIR /home/gitpod

### zsh ###
RUN zsh \
    sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" \
    rm -rf .zshrc \
    wget https://secman-team.github.io/docker/.zshrc \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting \
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions \
    sudo gem install colorls \
    source .zshrc

USER gitpod

### secman ###
RUN curl -fsSL https://secman-team.github.io/install.sh | bash
