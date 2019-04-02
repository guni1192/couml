PROJECT="github.com/guni1192/couml"
export GO111MODULE=on

all: couml

couml: cmd/*.go
	go build -o bin/couml $(PROJECT)/cmd

clean:
	rm bin/couml

test: **/*.go
	go test -v ./...

fmt: **/*.go
	go fmt ./...

vet: **/*.go
	go vet ./...

lint: **/*.go
	golint -set_exit_status ./...
