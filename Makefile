dev:
	./air

start:
	go run main.go

test:
	go test -v ./... -bench=.
