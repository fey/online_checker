run:
	go run cmd/checker/main.go \
	-url=$(URL)

build:
	goreleaser release --snapshot --clean

release:
	@./bin/release.sh
