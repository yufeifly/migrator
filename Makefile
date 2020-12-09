PROJECT_NAME=$(shell basename "$(PWD)")
GORUN=sudo go run

help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'


## run: sudo go run main.go
run:
	$(GORUN) main.go

## build: Compile the binary.
build:
	@go build -o $(PROJECT_NAME)

## echo: echo the project name
echo:
	@echo $(PROJECT_NAME)

## clean: rm the bin file
clean:
	rm -rf $(PROJECT_NAME)
