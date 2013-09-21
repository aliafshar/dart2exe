
build:
	go build -o bin/dart2exe *.go

fmt:
	gofmt -w *.go 
