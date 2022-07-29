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
trace-heap:
	go test ./internal/heap -run TestCopyIt -trace=./internal/heap/copy_trace.out
visual-trace-heap:
	go tool trace ./internal/heap/copy_trace.out
visual-cpubench-fib:
	go tool pprof ./internal/fib/cpuprofile.out
cpubench-fib:
	go test ./internal/fib -bench=. -count 5 -cpuprofile ./internal/fib/cpuprofile.out
run-generator:
	go run ./cmd/generator/main.go
run-sequence:
	cd ./cmd/sequence/ && go run main.go
compile-generator:
	echo "Compiling the generator for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/generator/main-freebsd-386 cmd/generator/main.go
	GOOS=linux GOARCH=386 go build -o bin/generator/main-linux-386 cmd/generator/main.go
	GOOS=windows GOARCH=386 go build -o bin/generator/main-windows-386 cmd/generator/main.go
compile-sequence:
	echo "Compiling the generator for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/sequence/main-freebsd-386 cmd/sequence/main.go
	GOOS=linux GOARCH=386 go build -gcflags '-m -l' -o bin/sequence/main-linux-386 cmd/sequence/main.go
	GOOS=windows GOARCH=386 go build -o bin/sequence/main-windows-386 cmd/sequence/main.go
