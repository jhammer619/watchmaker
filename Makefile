default: all
all: main.go
	go build
install:
	GOBIN=/usr/local/bin go install
clean:
	rm -f watchmaker
