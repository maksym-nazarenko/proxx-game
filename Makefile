
help: # print this help text
	@sed -n "/^[a-zA-Z0-9_-]*:/ s/:.*#/ -/p" < Makefile | sort

build: # build binary for the host platform
	@go build ./cmd/...

test: # run unit-tests
	@go test -v ./...
