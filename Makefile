list:
	@grep '^[^#[:space:]].*:' Makefile
print-hello:
	echo "Hello"
test-symbols:
	go test ./internal/symbols
run-generator:
	go run ./cmd/generator/main.go
run-sequence:
	go run ./cmd/sequence/main.go