COMMAND:=theapp
NPMBIN:=`npm bin`

build:
	go build -o $(COMMAND)

test: $(COMMAND)
	npm test

run:
	make build && make test

codecheck:
	$(NPMBIN)/codecheck
