test:
	go test

format:
	goimports -w noiferr.go

build:
	cd ./cmd/noiferr && go build -o noiferr

.PHONY: test format build
