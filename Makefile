run:
	go run main.go

install:
	go build -o bin/tacoma main.go
	cp bin/tacoma /usr/local/bin