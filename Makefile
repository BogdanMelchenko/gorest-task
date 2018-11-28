.PHONY: clean

clean:
		docker rm database application

.PHONY: build

build:
		go build -o ./application/gotasker -i ./application/main.go		
		docker build -t database ./database
		docker build -t application ./application
		rm ./application/gotasker		

.PHONY: up

up:
		docker-compose up

.PHONY: down

down:
		docker-compose down