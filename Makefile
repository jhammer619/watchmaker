default: all
all: main.go
	go build
install:
	GOBIN=~/.local/bin go install
clean:
	rm -f watchmaker
