
help: # print this help text
	@sed -n "/^[a-zA-Z0-9_-]*:/ s/:.*#/ -/p" < Makefile | sort

test: # run unit-tests
	@go test -v ./...
