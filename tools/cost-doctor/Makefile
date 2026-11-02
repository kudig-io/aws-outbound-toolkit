BINARY  := aws-cost-doctor
VERSION ?= dev
LDFLAGS := -s -w -X github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/cli.Version=$(VERSION)

.PHONY: build test vet fmt fmt-check run clean

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/aws-cost-doctor

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "unformatted files:"; gofmt -l .; exit 1)

run:
	go run ./cmd/aws-cost-doctor scan --input examples/inventory.json

clean:
	rm -rf bin