VERSION := $(shell git describe --tags --abbrev=0)
APP := pastebin

build: tailwind.css
	go build .

tailwind.css:
	curl https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4 > tailwind.css

release: tailwind.css
	GOOS=linux GOARCH=amd64 \
		go build -tags release \
		-ldflags '-X github.com/joseph0x45/pastebin/internal/buildinfo.Version=$(VERSION)' \
		-o $(APP)
