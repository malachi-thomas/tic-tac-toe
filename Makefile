MAIN_FILE = ./src/main.go

make:
		@go build $(MAIN_FILE) -o ./bin/
run:
	@go run $(MAIN_FILE)
