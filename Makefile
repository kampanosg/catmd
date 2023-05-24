default: install

.PHONY: install
install: build
	sudo mv catmd /usr/local/bin

.PHONY: build
build:
	go build -o catmd .

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...
