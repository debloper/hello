
.PHONY: build
build:
	go build server.go

.PHONY: fmt
fmt:
	@gofmt -l -s -w .

.PHONY: fmtcheck
fmtcheck:
	@gofmt -l -s . | read; if [ $$? == 0 ]; then echo "gofmt check failed for:"; gofmt -l -s .; exit 1; fi

.PHONY: clean
clean:
	rm -f server
