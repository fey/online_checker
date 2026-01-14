run:
	go run cmd/checker/main.go \
	-url=$(URL)

build:
	goreleaser release --snapshot --clean

release:
	@./bin/release.sh

ssh:
	make -C ansible ssh

deploy:
	make -C ansible deploy

fetch:
	make -C ansible fetch
