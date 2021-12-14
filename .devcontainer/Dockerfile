ARG VARIANT=16-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/javascript-node:0-${VARIANT}

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive

# Install Secman
RUN su node -c "npm install -g secman"

# Initialize `~/.secman`
RUN secman init
