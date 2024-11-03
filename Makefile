run:
	go mod tidy
	go build -o main main.go
	./main

clean:
	rm -rf test.db

run-docker:
	sudo docker build -t server-db .
	sudo docker run -p 8080:8080 server-db