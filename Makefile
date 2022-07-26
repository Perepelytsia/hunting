list:
	@grep '^[^#[:space:]].*:' Makefile
print-hello:
	echo "Hello"
test-symbols:
	go test ./internal/symbols
visual-cpubench-fib:
	go tool pprof ./internal/fib/cpuprofile.out
cpubench-fib:
	go test ./internal/fib -bench=. -count 5 -cpuprofile ./internal/fib/cpuprofile.out
visual-membench-fib:
	go tool pprof ./internal/fib/memprofile.out
membench-fib:
	go test ./internal/fib -bench=. -benchmem -memprofile ./internal/fib/memprofile.out
run-generator:
	go run ./cmd/generator/main.go
run-sequence:
	go run ./cmd/sequence/main.go
trace-heap:
	go test ./internal/heap -run TestCopyIt -trace=./internal/heap/copy_trace.out
visual-trace-heap:
	go tool trace ./internal/heap/copy_trace.out
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
