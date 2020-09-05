test: format
	go test

format:
	goimports -w noiferr.go

.PHONY: test format
