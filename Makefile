.PHONY:
all: build test examples

.PHONY:
build:
	docker build -t dynamicmultiwriter .

.PHONY:
test: build
	docker run dynamicmultiwriter go test

.PHONY:
examples: build
	docker run -it dynamicmultiwriter go run examples/add-remove/main.go
