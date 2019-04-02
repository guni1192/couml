PROJECT="github.com/guni1192/couml"

all: couml

couml: cmd/*.go
	go build -o bin/couml $(PROJECT)/cmd

clean:
	rm bin/couml

test: **/*.go
	go test -v ./...

fmt: **/*.go
	go fmt ./...

lint: **/*.go
	golint -set_exit_status ./...
