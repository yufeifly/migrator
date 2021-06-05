PROJECT_NAME=$(shell basename "$(PWD)")
GORUN=sudo -E go run

help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## run: sudo -E go run main.go
run:
	$(GORUN) main.go

## run2: DST=1 sudo -E go run main.go
rundst:
	DST=1 $(GORUN) main.go

## debug: DEBUG=1 sudo -E go run main.go
debug:
	DEBUG=1 $(GORUN) main.go

## debug2: DEBUG=1 DST=1 sudo -E go run main.go
debugdst:
	DEBUG=1 DST=1 $(GORUN) main.go

## build: Compile the binary.
build:
	@go build -o $(PROJECT_NAME)

## echo: echo the project name
echo:
	@echo $(PROJECT_NAME)

## clean: rm the bin file
clean:
	rm -rf $(PROJECT_NAME)
