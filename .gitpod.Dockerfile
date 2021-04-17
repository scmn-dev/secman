FROM gitpod/workflow-full

USER gitpod

RUN curl -fsSL https://secman-team.github.io/install.sh | bash
