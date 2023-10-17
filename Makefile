# Variables
BINARY_NAME=main

all: build

build:
	go build -o $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

test:
	go test -v ./...

clean:
	go clean
	rm $(BINARY_NAME)

docker-build:
	docker build -t myapp .

docker-run: docker-build
	docker run -p 8080:8080 myapp
