list:
	@grep '^[^#[:space:]].*:' Makefile
print-hello:
	echo "Hello"
test-symbols:
	go test ./internal/symbols
test-generator:
	go test ./internal/generator
bench-mem:
	go test ./internal/sequence -bench=. -benchmem -memprofile ./internal/sequence/memprofile.out
visual-bench-mem:
	go tool pprof ./internal/sequence/memprofile.out
bench-cpu:
	go test ./internal/sequence -bench=. -cpuprofile ./internal/sequence/cpuprofile.out
visual-bench-cpu:
	go tool pprof ./internal/sequence/cpuprofile.out
run-generator:
	go run ./cmd/generator/main.go
run-sequence:
	cd ./cmd/sequence/ && go run main.go
