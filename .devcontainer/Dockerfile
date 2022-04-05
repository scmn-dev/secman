ARG VARIANT=1.18
ARG NODE_VERSION="16"

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive && apt-get install software-properties-common curl

# Install Node.js
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Install Secman
RUN su userx -c "curl -sL https://u.secman.dev | bash"

# Initialize `~/.secman`
RUN secman init
