package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/creack/pty"
)

func main() {
	start := time.Now()
	run()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed Time: %v\n", elapsed)
}

func run() {
	var wg sync.WaitGroup
	var rspecErr, rubocopErr error
	var rspecBuf, rubocopBuf bytes.Buffer

	fmt.Println("Bulk Test Start...")

	wg.Add(2)

	// rspec
	go func() {
		defer wg.Done()
		rspecErr = runCmd(&rspecBuf, "rspec")
		if rspecErr != nil {
			fmt.Println("RSpec Failed.")
		} else {
			fmt.Println("RSpec Success!")
		}
	}()

	// rubocop
	go func() {
		defer wg.Done()
		rubocopErr = runCmd(&rubocopBuf, "rubocop")
		if rubocopErr != nil {
			fmt.Println("RuboCop Failed.")
		} else {
			fmt.Println("RuboCop Success!")
		}
	}()

	wg.Wait()

	if rspecErr != nil || rubocopErr != nil {
		fmt.Println()
		fmt.Println("Error Result.")
		fmt.Println("---------------")
		if rspecErr != nil {
			fmt.Printf("RSpec Errors:\n %v", rspecBuf.String())
		}
		if rspecErr != nil {
			fmt.Printf("RuboCop Errors:\n %v", rubocopBuf.String())
		}
		fmt.Println("---------------")
	} else {
		fmt.Println("All Success!!")
	}
}

func runCmd(w io.Writer, command string) error {
	cmd := exec.Command(command)

	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer func() { _ = ptmx.Close() }()

	io.Copy(w, ptmx)

	return cmd.Wait()
}
