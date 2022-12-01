build:
	go build -o jtracer bin/main.go

test:
	go test -v ./...