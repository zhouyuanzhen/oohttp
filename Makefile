all:
	@echo "Help and Usage information"

build_macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o oo_http_macos

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o oo_http_linux

build:
	CGO_ENABLED=0 GOARCH=amd64 go build -o oowebserver

build_all: build_linux build_macos

clean:
	@rm -f oohttp oo_http_linux oo_http_macos
