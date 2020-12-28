FROM ubuntu:latest

ARG PACKAGES="bash git tar gzip curl"
ARG smUrl=https://raw.githubusercontent.com/abdfnx/secman/HEAD/docker/secman
ARG smLoc=/usr/bin

RUN apt-get update
RUN apt-get install -y software-properties-common
RUN apt-get update

RUN apt-get install -y $PACKAGES
RUN apt-get update

RUN apt-get install -y wget
RUN apt-get update

RUN wget -P $smLoc $smUrl
RUN chmod 755 $smLoc/secman
RUN apt-get update
RUN secman init