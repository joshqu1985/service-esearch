all:
	@go build -o esearch cmd/*.go
	@echo done.


clean:
	@go clean
	@rm -f ./esearch
