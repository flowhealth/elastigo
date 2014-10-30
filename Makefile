.PHONY: init deps test

init:
	gvp init
	gvp in gpm local name github.com/flowhealth/elastigo
	gvp in gpm install

build: init
	gvp in go build .

test: init
	gvp in go test ./...

fast-test:
	gvp in go test ./...

