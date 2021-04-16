all:
	go build
clean:
	go clean
run:
	go run main.go

test: 
	go test -v ./...