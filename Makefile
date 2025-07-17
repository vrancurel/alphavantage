test:
	go test

check-fmt:
	gofmt -l *.go

GOPATH ?= $(HOME)/go

lint:
	if [ ! -f $(GOPATH)/bin/revive ]; then go install github.com/mgechev/revive@latest; fi
	$(GOPATH)/bin/revive -config revive.toml .
