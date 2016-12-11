COMMAND:=theapp

build:
	go build -o $(COMMAND)

test: $(COMMAND)
	npm test

run:
	make build && make test
