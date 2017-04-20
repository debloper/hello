
.PHONY: build
build:
	go build server.go

.PHONY: clean
clean:
	rm -f server
