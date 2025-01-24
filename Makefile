config ?= config/local.yaml

build:
	@go build -o bin/guided-weapon cmd/main.go
run: build
	@./bin/guided-weapon -config=$(config)