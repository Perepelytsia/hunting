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
profile-heap-sequence:
	(go run ./cmd/sequence/main.go &)
	sleep 21
	curl -o ./internal/sequence/memprofile.out http://localhost:6060/debug/pprof/heap
visual-heap-sequence:
	go tool pprof -top ./internal/sequence/memprofile.out