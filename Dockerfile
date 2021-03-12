FROM debian:latest

# vars
ARG GITHUB_URL="https://raw.githubusercontent.com"

RUN curl -fsSL $GITHUB_URL/secman-team/install/HEAD/install_linux.sh | bash

ENTRYPOINT [ "secman" ]
