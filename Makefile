LOCAL_BIN=$(CURDIR)/bin

.PHONY: .bindeps
.bindeps:
	mkdir -p bin
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.2

.PHONY: build
build: .bindeps
	protoc --proto_path=broker \
		--go_out=broker --go_opt=paths=source_relative \
		broker/broker.proto
	go build -o $(LOCAL_BIN)/protoc-gen-broker ./cmd/protoc-gen-broker

.PHONY: lint
lint: .bindeps
	$(LOCAL_BIN)/golangci-lint run --fix

.PHONY: clean
clean:
	rm -rf $(LOCAL_BIN)
