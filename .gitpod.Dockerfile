FROM gitpod/worklow-full

USER gitpod

RUN curl -fsSL https://secman-team.github.io/install.sh | bash
