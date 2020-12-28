FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

ARG PACKAGES="bash git tar gzip curl"
ARG smUrl=https://raw.githubusercontent.com/abdfnx/secman/main/secman
ARG smLocLD=/usr/bin

RUN apt-get update
RUN apt-get install -y software-properties-common
RUN apt-get update

RUN apt-get install -y $PACKAGES
RUN apt-get update

RUN apt-get install -y wget
RUN wget -P $smLocLD $smUrl
RUN chmod 755 $smLocLD/secman
