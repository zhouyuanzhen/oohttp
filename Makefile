all:
	@echo "Help and Usage information"

build_macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o oohttp_macos

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o oohttp_linux

build:
	CGO_ENABLED=0 GOARCH=amd64 go build -o oowebserver

build_all: build_linux build_macos

clean:
	@rm -f oohttp oohttp_linux oohttp_macos
