all:
	@go build -o esearch cmd/*.go
	@echo done.


clean:
	@go clean cmd/*
	@rm -f ./esearch
