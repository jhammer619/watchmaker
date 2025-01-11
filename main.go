package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gohugoio/hugo/watcher"
)

// newWatcher creates a new watcher to watch filesystem events.
func newWatcher(watched string) (*watcher.Batcher, error) {
	const batchInterval = 50

	// Batch filesystem events every batchInterval milliseconds.
	watcher, err := watcher.New(batchInterval * time.Millisecond)

	if err != nil {
		return nil, err
	}

	watcher.Add(watched)

	go func() {
		for {
			select {
			case evs := <-watcher.Events:
				fmt.Println(evs)
				cmd := exec.Command("make")
				cmd.Dir = filepath.Dir(watched)
				stdout, _ := cmd.StdoutPipe()

				cmd.Start()

				scanner := bufio.NewScanner(stdout)

				for scanner.Scan() {
					m := scanner.Text()
					fmt.Println(m)
				}

				cmd.Wait()
			case err := <-watcher.Errors:
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	return watcher, nil
}

func usage(progname string) {
	fmt.Printf(
		"\n%s: watch a file and when it changes automatically call `make` "+
			"in its directory\n\n"+
			"usage:\n"+
			"	%s path/to/file\n\n",
		progname, progname)
}

func main() {
	args := os.Args

	// Take only 1 argument: the path to the file to watch.
	if len(args) != 2 {
		usage(args[0])
		os.Exit(1)
	}

	watched := args[1]
	w, err := newWatcher(watched)

	if err != nil {
		fmt.Println("newWatcher returned error:", err)
		os.Exit(1)
	}

	defer w.Close()

	done := make(chan bool)
	<-done
}
