default: all
all: main.go watcher
	go build
install:
	GOBIN=/usr/local/bin go install
