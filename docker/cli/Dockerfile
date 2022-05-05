### image name: smcr/secman-cli ###

FROM debian:11

### begin ###

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y npm nodejs
RUN apt-get install -y curl sudo wget unzip

### nodejs & npm ###
RUN curl -sL https://deb.nodesource.com/setup_17.x -o nodesource_setup.sh && \
    bash nodesource_setup.sh && \
    apt-get nodejs build-essential -y && \
    rm -rf nodesource_setup.sh && \
    apt-get update
RUN npm i -g npm@latest

RUN npm i -g secman
RUN npm i -g @secman/scc

RUN secman init

ENTRYPOINT ["secman"]

### end ###
