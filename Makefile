dev:
	@nodemon --exec go run main.go --signal SIGTERM

run:
	@go run main.go

tidy:
	@go mod tidy

clean:
	@rm -f *.log

build:
	go build -ldflags "-s -w" -o runner

dockerize:
	docker-compose up --build
