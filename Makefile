.PHONY: run tidy

run:
	go run cmd/thc-cli-tool/main.go

build:
	go build cmd/thc-cli-tool/main.go


tidy:
	go mod tidy


simulate-cli:
	go build cmd/thc-cli-tool/main.go && mkdir simulate && mv main simulate && cd simulate && ./main

delete-simulate:
	rm -rf simulate