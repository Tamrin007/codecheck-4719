CMD="theapp"
NPMBIN:=`npm bin`

codecheck:
	$(NPMBIN)/codecheck

build:
	go build -o $(CMD)

run:
	make build && ./$(CMD)
