test:
	go test -v ./...

build:
	go build -o cmd/taskr/taskr ./cmd/taskr
