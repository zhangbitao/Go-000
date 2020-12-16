PROJECT_NAME := dialogue

GC=go build

.PHONY: all build pb

all: build

build:
	@$(GC) -o ${PROJECT_NAME} cmd/dialogue/main.go

pb:
	cd api && protoc --gofast_out=plugins=grpc:. dialogue/service/v1/dialogue.proto