FROM gitpod/workspace-full

RUN brew install gh
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Dev-x-Team/corgit/main/setup)"
RUN npm i -g @abdfnx/manx
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/abdfnx/secman/HEAD/packages/install_linux.sh)"
