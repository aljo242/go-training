BINARY_NAME = go-training
ARM = arm
MY_ARCH = $(shell go env GOARCH)
OUT_DIR = ./out

.PHONY: all
all: analyze build test

.PHONY: build
build:
	
ifeq ("$(wildcard ${OUT_DIR})", "")
	mkdir ${OUT_DIR}
endif
	go fmt
	go build -o ${OUT_DIR}/${BINARY_NAME}

.PHONY: analyze
analyze:
	golint
	go vet
	go fmt
	gosec ./...

.PHONY: test
test:
ifeq ("$(wildcard ${OUT_DIR})", "")
	mkdir ${OUT_DIR}
endif
# cannot use "-race" flag on ARM systems
ifeq ($(MY_ARCH), $(ARM))
	go test -v ./... -coverprofile=${OUT_DIR}/coverage.out
else
	go test -v ./... -race -coverprofile=${OUT_DIR}/coverage.out
endif
	go tool cover -html ${OUT_DIR}/coverage.out -o ${OUT_DIR}/coverage.html

.PHONY: clean
clean:
	go clean

.PHONY: run
	run: build
	sudo ${OUT_DIR}/${BINARY_NAME}
