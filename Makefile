all: build test;$(info $(M)...Begin to test and build all of binary.) @

# Build gitdog binary
build: ; $(info $(M)...Begin to build gitdog binary.)  @
	go build -v ./...
	go build -o bin/gitdog src/gitdog.go

# Test
test: ; $(info $(M)...Begin to test gitdog.)  @
	go test ./...;