
.PHONY: all changelog clean install build-binary

all: clean install build-binary

changelog:
	conventional-changelog -p angular -o CHANGELOG.md -w -r 0

clean:
	go clean -cache

install: clean
	go install -ldflags="-s -w" -trimpath ./cmd/iocgo

build-binary:
	go build -toolexec iocgo -ldflags="-s -w" -trimpath .
