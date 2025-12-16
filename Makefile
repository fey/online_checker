run:
	go run cmd/checker/main.go \
	-url=$(URL)

build:
	go build -o build/checker cmd/checker/main.go
