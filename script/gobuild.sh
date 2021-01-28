#!/bin/bash

set -eux

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/secman-team/secman"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}
mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCDIR}
cp core/secman.go ${SRCDIR}
(
	echo ${GOPATH}
	cd ${SRCDIR}
	go get .
	go install .
)
