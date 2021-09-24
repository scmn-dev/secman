.PHONY: build

LATEST_VERSION=$(shell git describe --abbrev=0 --tags | tr -d \n)
DATE=$(shell git tag -l --sort=-creatordate --format='%(creatordate:short)' $(LATEST_VERSION) | tr -d \n)

build:
		@cd core && \
		go get -d && \
		go build -o secman -ldflags "-X main.version=$(LATEST_VERSION) -X main.versionDate=($(DATE))"

setup: core/secman
		sudo mv core/secman /usr/local/bin
