.PHONY: build

LATEST_VERSION=$(shell git describe --abbrev=0 --tags | tr -d '\n')

build:
		@cd core && \
		go get -d && \
		go build -o secman -ldflags "-X main.version=$(LATEST_VERSION)"

setup: core/secman
		@sudo gem install colorize && \
		git clone https://github.com/secman-team/sm ~/sm && \
		sudo cp core/secman /usr/local/bin
