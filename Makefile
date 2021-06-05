.PHONY: test
test:
	go test
	go test ./pkg/text

.PHONY: build
build: deps
	go build ./cmd/greeting
