APP_NAME := pastebin
BIN_DIR := bin

build:
	go tool templ generate && \
	go generate && \
	go build -o $(BIN_DIR)/$(APP_NAME) .

build-release:
	go tool templ generate && \
	go generate && \
	go build -ldflags="-w -s" -o $(BIN_DIR)/$(APP_NAME)

clean:
	go clean
	rm -f $(BIN_DIR)/*

build-armv7:
	go tool templ generate && \
	go generate && \
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 \
	CC=arm-linux-gnueabihf-gcc \
	go build -ldflags='-linkmode external -extldflags "-static"' -o $(BIN_DIR)/$(APP_NAME)-armv7 .

build-all: build build-armv7
