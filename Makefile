build:
	go tool templ generate && \
	go generate && \
	go build .

build-release:
	go tool templ generate && \
	go generate && \
	go build -ldflags="-w -s"

clean:
	go clean
