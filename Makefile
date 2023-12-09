test:
	go test -v ./...

build:
	go build -o ./tmp/main cmd/main.go

run:
	go run cmd/main.go

run-test-db:
	docker run -it --rm --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres postgres
