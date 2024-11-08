
.PHONY: all changelog clean install build

all: clean install build

changelog:
	conventional-changelog -p angular -o CHANGELOG.md -w -r 0

clean:
	go clean -cache

install: clean
	go install -ldflags="-s -w" -trimpath ./cmd/iocgo

build:
	go build -toolexec iocgo -ldflags="-s -w" -trimpath -o server .
