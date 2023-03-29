.PHONY: build build-arm build-debian run

VERSION := test-build
BUILD := $$(git log -1 --pretty=%h)
BUILD_TIME := $$(date -u +"%Y%m%d.%H%M%S")

build:
	@go build \
		-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.BuildTime=${BUILD_TIME}" \
		-o ./bin/bento ./cmd

build-debian:
	@GOOS=linux GOARCH=amd64 go build \
		-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.BuildTime=${BUILD_TIME}" \
		-o ./bin/bento ./cmd

build-arm:
	@GOOS=linux GOARCH=arm GOARM=5 go build \
		-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -X main.BuildTime=${BUILD_TIME}" \
		-o ./bin/bento ./cmd
run:
	@go run ./cmd/main.go
