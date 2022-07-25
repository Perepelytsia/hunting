print-hello:
	echo "Hello"
run-generator:
	go run ./cmd/generator/main.go
run-sequence:
	go run ./cmd/sequence/main.go
compile-generator:
	echo "Compiling the generator for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/generator/main-freebsd-386 cmd/generator/main.go
	GOOS=linux GOARCH=386 go build -o bin/generator/main-linux-386 cmd/generator/main.go
	GOOS=windows GOARCH=386 go build -o bin/generator/main-windows-386 cmd/generator/main.go
compile-sequence:
	echo "Compiling the generator for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/sequence/main-freebsd-386 cmd/sequence/main.go
	GOOS=linux GOARCH=386 go build -o bin/sequence/main-linux-386 cmd/sequence/main.go
	GOOS=windows GOARCH=386 go build -o bin/sequence/main-windows-386 cmd/sequence/main.go
